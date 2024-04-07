package usecases

import (
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecases interface {
	CreateUser(auditID domain.AuditTraceID, user domain.NewUser, language string) (domain.Profile, *domain.ErrorDescription)

	GetUserByID(auditID domain.AuditTraceID, userID domain.UserID) (domain.User, *domain.ErrorDescription)
}

func (self UCs) newUserID() domain.UserID {
	return domain.UserID(uuid.New().String())
}

func (self UCs) CreateUser(auditID domain.AuditTraceID, newUser domain.NewUser, profileLanguage string) (domain.Profile, *domain.ErrorDescription) {
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

	user, err := self.i.Persistence.CreateUser(userToCreate)

	if err != nil {
		self.i.Trace.Log(auditID, auditStepID).WithField(domain.LogFieldProfileMail, newUser.Email).WithField(domain.LogFieldError, err).Msg("Error creating user")
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.Profile{}, &domain.ErrorUnableToCreateUser
	}

	profile, derr := self.createProfile(auditID, user.ID, newUser.Name, profileLanguage)

	if derr != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.Profile{}, derr
	}

	localizedDescription := self.i.Localize.GetLocalizedText(profileLanguage, domain.TranslatableDefaultOrganization, domain.TranslatableArgumentMap{
		domain.TranslatableArgumentUserFullName: newUser.Name,
	})

	organizationToCreate := domain.NewOrganization{
		Name:        newUser.Name,
		Description: localizedDescription,
	}

	_, derr = self.CreateOrganization(auditID, profile.ID, organizationToCreate)

	if derr != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.Profile{}, derr
	}

	self.i.Trace.Log(auditID, auditStepID).WithField(domain.LogFieldProfileMail, newUser.Email).Msg("Profile created")
	self.i.Trace.EndStep(auditID, auditStepID)
	return profile, nil
}

func (self UCs) GetUserByID(auditID domain.AuditTraceID, id domain.UserID) (domain.User, *domain.ErrorDescription) {
	auditStepID := self.i.Trace.AddStep(auditID)
	user, err := self.i.Persistence.FindUserByID(id)

	if err != nil {
		self.i.Trace.Log(auditID, auditStepID).WithLevel(domain.LogLevelError).WithField(domain.LogFieldError, err).WithField(domain.LogFieldUserID, id).Msg("Error finding user")
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.User{}, &domain.ErrorUserNotFound
	}
	self.i.Trace.EndStep(auditID, auditStepID)
	return user, nil
}
