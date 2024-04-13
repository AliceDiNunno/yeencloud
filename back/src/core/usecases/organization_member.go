package usecases

import (
	"github.com/AliceDiNunno/yeencloud/src/adapters/audit"
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
)

type OrganizationMemberUsecases interface {
	ListOrganizationsByProfile(rc *domain.RequestContext, profileID domain.ProfileID)
	ListOrganizationsMembers(auditID domain.AuditTraceID, organizationID domain.OrganizationID) ([]domain.OrganizationMember, *domain.ErrorDescription)

	removeAllMembersFromOrganization(auditID domain.AuditTraceID, organizationID domain.OrganizationID) *domain.ErrorDescription
}

func (self UCs) listOrganizationsByProfile(rc *domain.RequestContext, profileID domain.ProfileID) ([]domain.OrganizationMember, *domain.ErrorDescription) {
	organizations, err := self.i.Persistence.ListProfileOrganizationsByProfileID(profileID)

	if err != nil {
		self.log(rc, domain.LogLevelError).WithFields(domain.LogFields{
			domain.LogFieldError:     err,
			domain.LogFieldProfileID: profileID.String()}).
			Msg("Error getting organizations for user")
		return nil, &domain.ErrorUnableToGetOrganizationMembers
	}

	return organizations, nil
}

func (self UCs) ListOrganizationsByProfile(rc *domain.RequestContext, profileID domain.ProfileID) {
	self.traceRequest(rc, func() {
		self.requirePermission(rc, func() {
			rc.Done(self.listOrganizationsByProfile(rc, profileID))
		}, domain.PermissionOrganizationRead)
	})
}

func (self UCs) ListOrganizationsMembers(auditID domain.AuditTraceID, organizationID domain.OrganizationID) ([]domain.OrganizationMember, *domain.ErrorDescription) {
	auditStepID := self.i.Trace.AddStep(auditID, audit.DefaultSkip, organizationID)

	organizationMembers, err := self.i.Persistence.ListOrganizationMembers(organizationID)

	if err != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return nil, &domain.ErrorUnableToGetOrganizationMembers
	}

	self.i.Trace.EndStep(auditID, auditStepID)
	return organizationMembers, nil
}

func (self UCs) removeAllMembersFromOrganization(auditID domain.AuditTraceID, organizationID domain.OrganizationID) *domain.ErrorDescription {
	auditStepID := self.i.Trace.AddStep(auditID, audit.DefaultSkip, organizationID)

	err := self.i.Persistence.RemoveAllMembersFromOrganization(organizationID)

	if err != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return &domain.ErrorUnableToRemoveOrganizationMembers
	}

	self.i.Trace.EndStep(auditID, auditStepID)
	return nil

}
