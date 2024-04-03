package usecases

import (
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/google/uuid"
)

type ProfileUsecases interface {
	GetProfileByUserID(auditID domain.AuditID, userID domain.UserID) (domain.Profile, *domain.ErrorDescription)

	createProfile(auditID domain.AuditID, userID domain.UserID, name string, language string) (domain.Profile, *domain.ErrorDescription)
}

func (self UCs) newProfileID() domain.ProfileID {
	return domain.ProfileID(uuid.New().String())
}

func (self UCs) createProfile(auditID domain.AuditID, userID domain.UserID, name string, language string) (domain.Profile, *domain.ErrorDescription) {
	auditStepID := self.i.Trace.AddStep(auditID, userID, name, language)

	profileToCreate := domain.Profile{
		ID:       self.newProfileID(),
		UserID:   userID,
		Name:     name,
		Language: language,
	}

	profile, err := self.i.Persistence.Profile.CreateProfile(profileToCreate)

	if err != nil {
		self.i.Trace.Log(auditID, auditStepID).WithField(domain.LogFieldProfileName, name).Msg("Error creating profile for user")
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.Profile{}, &domain.ErrorUnableToCreateProfile
	}

	self.i.Trace.EndStep(auditID, auditStepID)
	return profile, nil
}

func (self UCs) GetProfileByUserID(auditID domain.AuditID, userID domain.UserID) (domain.Profile, *domain.ErrorDescription) {
	auditStepID := self.i.Trace.AddStep(auditID)

	profile, err := self.i.Persistence.Profile.FindProfileByUserID(userID)

	if err != nil {
		self.i.Trace.Log(auditID, auditStepID).WithLevel(domain.LogLevelError).WithField(domain.LogFieldError, err).Msg("Error fetching profile by user ID")
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.Profile{}, &domain.ErrorProfileNotFound
	}

	self.i.Trace.EndStep(auditID, auditStepID)
	return profile, nil
}
