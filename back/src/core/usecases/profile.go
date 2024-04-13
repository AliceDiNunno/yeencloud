package usecases

import (
	"github.com/AliceDiNunno/yeencloud/src/adapters/audit"
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
	auditStepID := self.i.Trace.AddStep(auditID, audit.DefaultSkip, userID, name, language)

	profileToCreate := domain.Profile{
		ID:       self.newProfileID(),
		UserID:   userID,
		Name:     name,
		Language: language,
		Role:     domain.RoleProfileUnvalidated.String(),
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

func (self UCs) SetProfileRole(auditID domain.AuditTraceID, profileID domain.ProfileID, role domain.Role) *domain.ErrorDescription {
	auditStepID := self.i.Trace.AddStep(auditID, audit.DefaultSkip, profileID, role)

	err := self.i.Persistence.SetProfileRole(profileID, role)

	if err != nil {
		self.i.Trace.Log(auditID, auditStepID).WithField(domain.LogFieldProfileID, profileID).WithField(domain.LogFieldProfileRole, role).Msg("Error setting profile role")
		self.i.Trace.EndStep(auditID, auditStepID)
		return &domain.ErrUnableToSetProfileRole
	}

	self.i.Trace.EndStep(auditID, auditStepID)
	return nil
}

func (self UCs) GetProfileByUserID(auditID domain.AuditTraceID, userID domain.UserID) (domain.Profile, *domain.ErrorDescription) {
	auditStepID := self.i.Trace.AddStep(auditID, audit.DefaultSkip)

	profile, err := self.i.Persistence.FindProfileByUserID(userID)

	if err != nil {
		self.i.Trace.Log(auditID, auditStepID).WithLevel(domain.LogLevelError).WithField(domain.LogFieldError, err).Msg("Error fetching profile by user ID")
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.Profile{}, &domain.ErrorProfileNotFound
	}

	self.i.Trace.EndStep(auditID, auditStepID)
	return profile, nil
}
