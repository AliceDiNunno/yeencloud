package usecases

import (
	"github.com/AliceDiNunno/yeencloud/src/adapters/audit"
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/google/uuid"
)

type OrganizationUsecases interface {
	CreateOrganization(auditID domain.AuditTraceID, profileID domain.ProfileID, organization domain.NewOrganization) (domain.Organization, *domain.ErrorDescription)

	GetOrganizationByID(auditID domain.AuditTraceID, profileID domain.ProfileID, organizationID domain.OrganizationID) (domain.Organization, *domain.ErrorDescription)

	UpdateOrganization(auditID domain.AuditTraceID, profileID domain.ProfileID, organizationID domain.OrganizationID, update domain.UpdateOrganization) (domain.Organization, *domain.ErrorDescription)

	DeleteOrganization(auditID domain.AuditTraceID, profileID domain.ProfileID, organizationID domain.OrganizationID) *domain.ErrorDescription
}

func (self UCs) CreateOrganization(auditID domain.AuditTraceID, profileID domain.ProfileID, newOrganization domain.NewOrganization) (domain.Organization, *domain.ErrorDescription) {
	auditStepID := self.i.Trace.AddStep(auditID, audit.DefaultSkip, newOrganization)

	if err := self.checkPermissions(auditID, profileID, nil, domain.PermissionGlobalOrganizationCreation); err != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.Organization{}, err
	}

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

	err = self.i.Persistence.LinkProfileToOrganization(profileID, organization.ID, domain.RoleOrganizationOwner)

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

func (self UCs) GetOrganizationByID(auditID domain.AuditTraceID, profileID domain.ProfileID, organizationID domain.OrganizationID) (domain.Organization, *domain.ErrorDescription) {
	auditStepID := self.i.Trace.AddStep(auditID, audit.DefaultSkip, profileID, organizationID)

	role, err := self.i.Persistence.GetOrganizationMemberRole(profileID, organizationID)

	if err != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.Organization{}, &domain.ErrorOrganizationNotFound
	}

	if err := self.checkPermissions(auditID, profileID, &role, domain.PermissionOrganizationRead); err != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.Organization{}, err
	}

	organization, err := self.i.Persistence.GetOrganizationByIDAndProfileID(profileID, organizationID)

	if err != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.Organization{}, &domain.ErrorOrganizationNotFound
	}

	self.i.Trace.EndStep(auditID, auditStepID)
	return organization, nil
}

func (self UCs) UpdateOrganization(auditID domain.AuditTraceID, profileID domain.ProfileID, organizationID domain.OrganizationID, update domain.UpdateOrganization) (domain.Organization, *domain.ErrorDescription) {
	auditStepID := self.i.Trace.AddStep(auditID, audit.DefaultSkip, profileID, organizationID, update)

	role, err := self.i.Persistence.GetOrganizationMemberRole(profileID, organizationID)

	if err != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.Organization{}, &domain.ErrorOrganizationNotFound
	}

	if err := self.checkPermissions(auditID, profileID, &role, domain.PermissionOrganizationMetaUpdate); err != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.Organization{}, err
	}

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

func (self UCs) DeleteOrganization(auditID domain.AuditTraceID, profileID domain.ProfileID, organizationID domain.OrganizationID) *domain.ErrorDescription {
	auditStepID := self.i.Trace.AddStep(auditID, audit.DefaultSkip, profileID, organizationID)

	role, err := self.i.Persistence.GetOrganizationMemberRole(profileID, organizationID)

	if err != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return &domain.ErrorOrganizationNotFound
	}

	if err := self.checkPermissions(auditID, profileID, &role, domain.PermissionOrganizationMetaUpdate); err != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return err
	}

	derr := self.removeAllMembersFromOrganization(auditID, organizationID)
	if derr != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return derr
	}

	err = self.i.Persistence.DeleteOrganizationByID(organizationID)
	if err != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return &domain.ErrorUnableToDeleteOrganization
	}

	self.i.Trace.EndStep(auditID, auditStepID)
	return nil
}
