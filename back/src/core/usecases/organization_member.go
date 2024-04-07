package usecases

import "github.com/AliceDiNunno/yeencloud/src/core/domain"

type OrganizationMemberUsecases interface {
	ListOrganizationsByProfile(auditID domain.AuditTraceID, profileID domain.ProfileID) ([]domain.OrganizationMember, *domain.ErrorDescription)
	ListOrganizationsMembers(auditID domain.AuditTraceID, organizationID domain.OrganizationID) ([]domain.OrganizationMember, *domain.ErrorDescription)

	removeAllMembersFromOrganization(auditID domain.AuditTraceID, organizationID domain.OrganizationID) *domain.ErrorDescription
}

func (self UCs) ListOrganizationsByProfile(auditID domain.AuditTraceID, profileID domain.ProfileID) ([]domain.OrganizationMember, *domain.ErrorDescription) {
	auditStepID := self.i.Trace.AddStep(auditID, profileID)

	organizations, err := self.i.Persistence.ListProfileOrganizationsByProfileID(profileID)

	if err != nil {
		self.i.Trace.Log(auditID, auditStepID).WithLevel(domain.LogLevelError).WithFields(domain.LogFields{
			domain.LogFieldError:     err,
			domain.LogFieldProfileID: profileID.String()}).
			Msg("Error getting organizations for user")

		self.i.Trace.EndStep(auditID, auditStepID)
		return nil, &domain.ErrorUnableToGetOrganizationMembers
	}

	self.i.Trace.EndStep(auditID, auditStepID)
	return organizations, nil
}

func (self UCs) ListOrganizationsMembers(auditID domain.AuditTraceID, organizationID domain.OrganizationID) ([]domain.OrganizationMember, *domain.ErrorDescription) {
	auditStepID := self.i.Trace.AddStep(auditID, organizationID)

	organizationMembers, err := self.i.Persistence.ListOrganizationMembers(organizationID)

	if err != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return nil, &domain.ErrorUnableToGetOrganizationMembers
	}

	self.i.Trace.EndStep(auditID, auditStepID)
	return organizationMembers, nil
}

func (self UCs) removeAllMembersFromOrganization(auditID domain.AuditTraceID, organizationID domain.OrganizationID) *domain.ErrorDescription {
	auditStepID := self.i.Trace.AddStep(auditID, organizationID)

	err := self.i.Persistence.RemoveAllMembersFromOrganization(organizationID)

	if err != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return &domain.ErrorUnableToRemoveOrganizationMembers
	}

	self.i.Trace.EndStep(auditID, auditStepID)
	return nil

}
