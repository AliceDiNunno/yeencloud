package usecases

import (
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/google/uuid"
)

func (self UCs) CreateOrganization(auditID domain.AuditID, profileID domain.ProfileID, newOrganization domain.NewOrganization) (domain.Organization, *domain.ErrorDescription) {
	auditStepID := self.i.Trace.AddStep(auditID, newOrganization)

	organizationToCreate := domain.Organization{
		ID:          domain.OrganizationID(uuid.New().String()),
		Slug:        newOrganization.Name, // #YC-17 TODO: generate unique slug
		Name:        newOrganization.Name,
		Description: newOrganization.Description,
	}

	organization, err := self.i.Persistence.Organization.CreateOrganization(organizationToCreate)

	if err != nil {
		self.i.Trace.Log(auditID, auditStepID).WithFields(domain.LogFields{
			domain.LogFieldError:  err,
			domain.LogFieldUserID: profileID.String()}).
			Msg("Error creating organization for user")
	}

	err = self.i.Persistence.OrganizationProfile.LinkProfileToOrganization(profileID, organization.ID, "admin")

	if err != nil {
		self.i.Trace.Log(auditID, auditStepID).WithFields(domain.LogFields{
			domain.LogFieldError:  err,
			domain.LogFieldUserID: profileID.String()}).
			Msg("Error linking user to organization")
	}

	self.i.Trace.EndStep(auditID, auditStepID)

	return organization, nil
}

func (self UCs) GetOrganizationsByProfileID(auditID domain.AuditID, profileID domain.ProfileID) ([]domain.OrganizationMember, *domain.ErrorDescription) {
	auditStepID := self.i.Trace.AddStep(auditID, profileID)

	organizations, err := self.i.Persistence.OrganizationProfile.GetProfileOrganizationsByProfileID(profileID)

	if err != nil {
		self.i.Trace.Log(auditID, auditStepID).WithLevel(domain.LogLevelError).WithFields(domain.LogFields{
			domain.LogFieldError:     err,
			domain.LogFieldProfileID: profileID.String()}).
			Msg("Error getting organizations for user")

		return nil, &domain.ErrorUnableToGetUserOrganizations
	}

	return organizations, nil
}
