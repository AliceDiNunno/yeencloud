package usecases

import (
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/google/uuid"
)

type OrganizationUsecases interface {
	CreateOrganization(auditID domain.AuditID, profileID domain.ProfileID, organization domain.NewOrganization) (domain.Organization, *domain.ErrorDescription)

	ListOrganizationsByProfile(auditID domain.AuditID, profileID domain.ProfileID) ([]domain.OrganizationMember, *domain.ErrorDescription)
	ListOrganizationsMembers(auditID domain.AuditID, organizationID domain.OrganizationID) ([]domain.OrganizationMember, *domain.ErrorDescription)
	GetOrganizationByID(auditID domain.AuditID, profileID domain.ProfileID, organizationID domain.OrganizationID) (domain.Organization, *domain.ErrorDescription)

	UpdateOrganization(auditID domain.AuditID, profileID domain.ProfileID, organizationID domain.OrganizationID, update domain.UpdateOrganization) (domain.Organization, *domain.ErrorDescription)

	DeleteOrganization(auditID domain.AuditID, profileID domain.ProfileID, organizationID domain.OrganizationID) *domain.ErrorDescription
}

func (self UCs) CreateOrganization(auditID domain.AuditID, profileID domain.ProfileID, newOrganization domain.NewOrganization) (domain.Organization, *domain.ErrorDescription) {
	auditStepID := self.i.Trace.AddStep(auditID, newOrganization)

	organizationToCreate := domain.Organization{
		ID:          domain.OrganizationID(uuid.New().String()),
		Slug:        newOrganization.Name, // #YC-17 TODO: generate unique slug
		Name:        newOrganization.Name,
		Description: newOrganization.Description,
	}

	organization, err := self.i.Persistence.CreateOrganization(organizationToCreate)

	if err != nil {
		self.i.Trace.Log(auditID, auditStepID).WithFields(domain.LogFields{
			domain.LogFieldError:  err,
			domain.LogFieldUserID: profileID.String()}).
			Msg("Error creating organization for user")
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.Organization{}, &domain.ErrorUnableToCreateOrganization
	}

	err = self.i.Persistence.LinkProfileToOrganization(profileID, organization.ID, domain.OrganizationRoleOwner)

	if err != nil {
		self.i.Trace.Log(auditID, auditStepID).WithFields(domain.LogFields{
			domain.LogFieldError:  err,
			domain.LogFieldUserID: profileID.String()}).
			Msg("Error linking user to organization")
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.Organization{}, &domain.ErrorUnableToLinkUserOrganization
	}

	self.i.Trace.EndStep(auditID, auditStepID)
	return organization, nil
}

func (self UCs) GetOrganizationsByProfileID(auditID domain.AuditID, profileID domain.ProfileID) ([]domain.OrganizationMember, *domain.ErrorDescription) {
	auditStepID := self.i.Trace.AddStep(auditID, profileID)

	organizations, err := self.i.Persistence.ListProfileOrganizationsByProfileID(profileID)

	if err != nil {
		self.i.Trace.Log(auditID, auditStepID).WithLevel(domain.LogLevelError).WithFields(domain.LogFields{
			domain.LogFieldError:     err,
			domain.LogFieldProfileID: profileID.String()}).
			Msg("Error getting organizations for user")

		self.i.Trace.EndStep(auditID, auditStepID)
		return nil, &domain.ErrorUnableToGetUserOrganizations
	}

	self.i.Trace.EndStep(auditID, auditStepID)
	return organizations, nil
}

func (self UCs) ListOrganizationsByProfile(auditID domain.AuditID, profileID domain.ProfileID) ([]domain.OrganizationMember, *domain.ErrorDescription) {
	auditStepID := self.i.Trace.AddStep(auditID, profileID)

	organizations, err := self.i.Persistence.ListProfileOrganizationsByProfileID(profileID)

	if err != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return nil, &domain.ErrorUnableToGetOrganizationMembers
	}

	self.i.Trace.EndStep(auditID, auditStepID)
	return organizations, nil
}

func (self UCs) ListOrganizationsMembers(auditID domain.AuditID, organizationID domain.OrganizationID) ([]domain.OrganizationMember, *domain.ErrorDescription) {
	auditStepID := self.i.Trace.AddStep(auditID, organizationID)

	organizationMembers, err := self.i.Persistence.ListOrganizationMembers(organizationID)

	if err != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return nil, &domain.ErrorUnableToGetOrganizationMembers
	}

	self.i.Trace.EndStep(auditID, auditStepID)
	return organizationMembers, nil
}

func (self UCs) GetOrganizationByID(auditID domain.AuditID, profileID domain.ProfileID, organizationID domain.OrganizationID) (domain.Organization, *domain.ErrorDescription) {
	auditStepID := self.i.Trace.AddStep(auditID, profileID, organizationID)

	organization, err := self.i.Persistence.GetOrganizationByIDAndProfileID(profileID, organizationID)

	if err != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.Organization{}, &domain.ErrorOrganizationNotFound
	}

	self.i.Trace.EndStep(auditID, auditStepID)
	return organization, nil
}

func (self UCs) UpdateOrganization(auditID domain.AuditID, profileID domain.ProfileID, organizationID domain.OrganizationID, update domain.UpdateOrganization) (domain.Organization, *domain.ErrorDescription) {
	auditStepID := self.i.Trace.AddStep(auditID, profileID, organizationID, update)

	organization, err := self.i.Persistence.GetOrganizationByIDAndProfileID(profileID, organizationID)

	if err != nil {
		return domain.Organization{}, &domain.ErrorOrganizationNotFound
	}

	if update.Name == "" {
		update.Name = organization.Name
	}

	if update.Description == "" {
		update.Description = organization.Description
	}

	organization, err = self.i.Persistence.UpdateOrganization(organization.ID, update)

	if err != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.Organization{}, &domain.ErrorUnableToUpdateOrganization
	}

	self.i.Trace.EndStep(auditID, auditStepID)
	return organization, nil
}

func (self UCs) DeleteOrganization(auditID domain.AuditID, profileID domain.ProfileID, organizationID domain.OrganizationID) *domain.ErrorDescription {
	auditStepID := self.i.Trace.AddStep(auditID, profileID, organizationID)

	err := self.i.Persistence.DeleteOrganizationByID(organizationID)
	if err != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return &domain.ErrorUnableToDeleteOrganization
	}

	self.i.Trace.EndStep(auditID, auditStepID)
	return nil
}
