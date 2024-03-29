package usecases

import (
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/google/uuid"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/crypto/bcrypt"
)

func (self UCs) newUserID() domain.UserID {
	return domain.UserID(uuid.New().String())
}
func (self UCs) newProfileID() domain.ProfileID {
	return domain.ProfileID(uuid.New().String())
}

func (self UCs) CreateUser(auditID domain.AuditID, newUser domain.NewUser, profileLanguage string) (domain.Profile, *domain.ErrorDescription) {
	auditStepID := self.i.Auditer.AddStep(auditID, newUser.Secure())

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)

	if err != nil {
		self.i.Auditer.Log(auditID, auditStepID).WithField(domain.LogFieldMail, newUser.Email).Msg("Error hashing password")
		self.i.Auditer.EndStep(auditID, auditStepID)
		return domain.Profile{}, &domain.ErrorUnableToHashPassword
	}

	userToCreate := domain.User{
		ID:       self.newUserID(),
		Email:    newUser.Email,
		Password: string(hashedPassword),
	}

	user, err := self.i.Persistence.User.CreateUser(userToCreate)

	if err != nil {
		self.i.Auditer.Log(auditID, auditStepID).WithField(domain.LogFieldMail, newUser.Email).Msg("Error creating user")
		self.i.Auditer.EndStep(auditID, auditStepID)
		return domain.Profile{}, &domain.ErrorUserAlreadyExists //TODO: wrong error ?
	}

	profileToCreate := domain.Profile{
		ID:       self.newProfileID(),
		UserID:   user.ID,
		Name:     newUser.Name,
		Language: profileLanguage,
	}

	profile, err := self.i.Persistence.Profile.CreateProfile(profileToCreate)

	if err != nil {
		self.i.Auditer.Log(auditID, auditStepID).WithField(domain.LogFieldMail, newUser.Email).Msg("Error creating profile for user")
		self.i.Auditer.EndStep(auditID, auditStepID)
		return domain.Profile{}, &domain.ErrorUserAlreadyExists //TODO: wrong error ?
	}

	msg := i18n.NewLocalizer(self.i.Translator, profileLanguage)

	localizedDescription, _, _ := msg.LocalizeWithTag(&i18n.LocalizeConfig{
		MessageID: domain.DefaultOrganizationDescription,
		TemplateData: map[string]interface{}{
			domain.DefaultOrganizationDescriptionKey: newUser.Name,
		},
	})

	organizationToCreate := domain.NewOrganization{
		Name:        newUser.Name,
		Description: localizedDescription,
	}

	_, derr := self.CreateOrganization(auditID, profile.ID, organizationToCreate)

	if derr != nil {
		self.i.Auditer.EndStep(auditID, auditStepID)
		return domain.Profile{}, derr
	}

	self.i.Auditer.Log(auditID, auditStepID).WithField(domain.LogFieldMail, newUser.Email).Msg("Profile created")
	self.i.Auditer.EndStep(auditID, auditStepID)
	return profileToCreate, nil
}

func (self UCs) GetUserByID(auditID domain.AuditID, id domain.UserID) (domain.User, *domain.ErrorDescription) {
	auditStepID := self.i.Auditer.AddStep(auditID)
	user, err := self.i.Persistence.User.FindUserByID(id)

	if err != nil {
		self.i.Auditer.Log(auditID, auditStepID).WithLevel(domain.LogLevelError).WithField("error", err).WithField("id", id.String()).Msg("Error finding user")
		self.i.Auditer.EndStep(auditID, auditStepID)
		return domain.User{}, &domain.ErrorUserNotFound
	}
	self.i.Auditer.EndStep(auditID, auditStepID)
	return user, nil
}

func (self UCs) GetProfileByUserID(auditID domain.AuditID, userID domain.UserID) (domain.Profile, *domain.ErrorDescription) {
	auditStepID := self.i.Auditer.AddStep(auditID)

	profile, err := self.i.Persistence.Profile.FindProfileByUserID(userID)

	// #YC-22 TODO: this should never happen, a profile should be created if it ever is missing (while also reporting the error so it can be investigated)
	if err != nil {
		self.i.Auditer.Log(auditID, auditStepID).WithLevel(domain.LogLevelError).WithField("error", err).Msg("Error finding user")
		self.i.Auditer.EndStep(auditID, auditStepID)
		return domain.Profile{}, &domain.ErrorProfileNotFound
	}

	self.i.Auditer.EndStep(auditID, auditStepID)
	return profile, nil
}
