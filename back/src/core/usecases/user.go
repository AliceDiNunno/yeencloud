package usecases

import (
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (self UCs) newUserID() domain.UserID {
	return domain.UserID(uuid.New().String())
}
func (self UCs) newProfileID() domain.ProfileID {
	return domain.ProfileID(uuid.New().String())
}

func (self UCs) CreateUser(auditID domain.AuditID, newUser domain.NewUser, profileLanguage string) (domain.Profile, *domain.ErrorDescription) {
	auditStepID := self.i.Trace.AddStep(auditID, newUser.Secure())

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)

	if err != nil {
		self.i.Trace.Log(auditID, auditStepID).WithField(domain.LogFieldProfileMail, newUser.Email).Msg("Error hashing password")
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.Profile{}, &domain.ErrorUnableToHashPassword
	}

	userToCreate := domain.User{
		ID:       self.newUserID(),
		Email:    newUser.Email,
		Password: string(hashedPassword),
	}

	user, err := self.i.Persistence.User.CreateUser(userToCreate)

	if err != nil {
		self.i.Trace.Log(auditID, auditStepID).WithField(domain.LogFieldProfileMail, newUser.Email).Msg("Error creating user")
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.Profile{}, &domain.ErrorUserAlreadyExists // TODO: wrong error ?
	}

	profileToCreate := domain.Profile{
		ID:       self.newProfileID(),
		UserID:   user.ID,
		Name:     newUser.Name,
		Language: profileLanguage,
	}

	profile, err := self.i.Persistence.Profile.CreateProfile(profileToCreate)

	if err != nil {
		self.i.Trace.Log(auditID, auditStepID).WithField(domain.LogFieldProfileMail, newUser.Email).Msg("Error creating profile for user")
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.Profile{}, &domain.ErrorUserAlreadyExists // TODO: wrong error ?
	}

	localizedDescription := self.i.Localize.GetLocalizedText(profileLanguage, domain.TranslatableDefaultOrganization, domain.TranslatableArgumentMap{
		domain.TranslatableArgumentUserFullName: newUser.Name,
	})

	organizationToCreate := domain.NewOrganization{
		Name:        newUser.Name,
		Description: localizedDescription,
	}

	_, derr := self.CreateOrganization(auditID, profile.ID, organizationToCreate)

	if derr != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.Profile{}, derr
	}

	self.i.Trace.Log(auditID, auditStepID).WithField(domain.LogFieldProfileMail, newUser.Email).Msg("Profile created")
	self.i.Trace.EndStep(auditID, auditStepID)
	return profileToCreate, nil
}

func (self UCs) GetUserByID(auditID domain.AuditID, id domain.UserID) (domain.User, *domain.ErrorDescription) {
	auditStepID := self.i.Trace.AddStep(auditID)
	user, err := self.i.Persistence.User.FindUserByID(id)

	if err != nil {
		self.i.Trace.Log(auditID, auditStepID).WithLevel(domain.LogLevelError).WithField(domain.LogFieldError, err).WithField(domain.LogFieldUserID, id).Msg("Error finding user")
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.User{}, &domain.ErrorUserNotFound
	}
	self.i.Trace.EndStep(auditID, auditStepID)
	return user, nil
}

func (self UCs) GetProfileByUserID(auditID domain.AuditID, userID domain.UserID) (domain.Profile, *domain.ErrorDescription) {
	auditStepID := self.i.Trace.AddStep(auditID)

	profile, err := self.i.Persistence.Profile.FindProfileByUserID(userID)

	// #YC-22 TODO: this should never happen, a profile should be created if it ever is missing (while also reporting the error so it can be investigated)
	if err != nil {
		self.i.Trace.Log(auditID, auditStepID).WithLevel(domain.LogLevelError).WithField(domain.LogFieldError, err).Msg("Error finding user")
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.Profile{}, &domain.ErrorProfileNotFound
	}

	self.i.Trace.EndStep(auditID, auditStepID)
	return profile, nil
}
