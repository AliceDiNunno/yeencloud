package usecases

import (
	"github.com/AliceDiNunno/yeencloud/internal/adapters/audit"
	"github.com/AliceDiNunno/yeencloud/internal/core/domain"
)

func (self UCs) roleHasPermission(role domain.Role, permission domain.Permission) bool {
	for _, p := range role.Permissions {
		if p == permission {
			return true
		}
	}

	return false
}

func (self UCs) userHasPermission(auditID domain.AuditTraceID, profileID domain.ProfileID, permission domain.Permission) bool {
	auditStepID := self.i.Trace.AddStep(auditID, audit.DefaultSkip, profileID, permission)

	profile, err := self.i.Persistence.FindProfileByID(profileID)
	if err != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return false
	}

	role := domain.RoleByName(profile.Role)

	if role.HasPermission(permission) {
		self.i.Trace.EndStep(auditID, auditStepID)
		return true

	}

	self.i.Trace.EndStep(auditID, auditStepID)
	return false
}

func (self UCs) checkPermissions(auditID domain.AuditTraceID, profileID domain.ProfileID, profileToObjectRelationRole *string, permission domain.Permission) error {
	auditStepID := self.i.Trace.AddStep(auditID, audit.DefaultSkip, profileID, profileToObjectRelationRole, permission)

	if self.userHasPermission(auditID, profileID, permission) {
		self.i.Trace.EndStep(auditID, auditStepID)
		return nil
	}

	if profileToObjectRelationRole != nil {
		role := domain.RoleByName(*profileToObjectRelationRole)

		if role.HasPermission(permission) {
			return nil
		}
	}

	self.i.Trace.EndStep(auditID, auditStepID)
	return domain.ErrorPermissionRequired(permission)
}
