package usecases

import (
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/google/uuid"
)

type ProfileUsecases interface {
	GetProfileByUserID(auditID domain.AuditTraceID, userID domain.UserID) (domain.Profile, *domain.ErrorDescription)

	createProfile(auditID domain.AuditTraceID, userID domain.UserID, name string, language string) (domain.Profile, *domain.ErrorDescription)
}

func (self UCs) newProfileID() domain.ProfileID {
	return domain.ProfileID(uuid.New().String())
}

func (self UCs) createProfile(auditID domain.AuditTraceID, userID domain.UserID, name string, language string) (domain.Profile, *domain.ErrorDescription) {
	auditStepID := self.i.Trace.AddStep(auditID, userID, name, language)

	profileToCreate := domain.Profile{
		ID:       self.newProfileID(),
		UserID:   userID,
		Name:     name,
		Language: language,
		// TODO: Should be role user limited however, if we do that, the user will not be able to create organizations and since we create one on user creation it will fail
		// we should send a mail to the user to confirm his email and then change his role to standard then create the organization
		Role: domain.RoleUserStandard.String(),
	}

	profile, err := self.i.Persistence.CreateProfile(profileToCreate)

	if err != nil {
		self.i.Trace.Log(auditID, auditStepID).WithField(domain.LogFieldProfileName, name).Msg("Error creating profile for user")
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.Profile{}, &domain.ErrorUnableToCreateProfile
	}

	self.i.Trace.EndStep(auditID, auditStepID)
	return profile, nil
}

func (self UCs) GetProfileByUserID(auditID domain.AuditTraceID, userID domain.UserID) (domain.Profile, *domain.ErrorDescription) {
	auditStepID := self.i.Trace.AddStep(auditID)

	profile, err := self.i.Persistence.FindProfileByUserID(userID)

	if err != nil {
		self.i.Trace.Log(auditID, auditStepID).WithLevel(domain.LogLevelError).WithField(domain.LogFieldError, err).Msg("Error fetching profile by user ID")
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.Profile{}, &domain.ErrorProfileNotFound
	}

	self.i.Trace.EndStep(auditID, auditStepID)
	return profile, nil
}
