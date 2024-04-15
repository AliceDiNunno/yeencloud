package usecases

import (
	"github.com/AliceDiNunno/yeencloud/internal/adapters/audit"
	"github.com/AliceDiNunno/yeencloud/internal/core/domain"
	"github.com/google/uuid"
)

type OrganizationUsecases interface {
	CreateOrganization(auditID domain.AuditTraceID, profileID domain.ProfileID, organization domain.NewOrganization) (domain.Organization, error)

	GetOrganizationByID(auditID domain.AuditTraceID, profileID domain.ProfileID, organizationID domain.OrganizationID) (domain.Organization, error)

	UpdateOrganization(auditID domain.AuditTraceID, profileID domain.ProfileID, organizationID domain.OrganizationID, update domain.UpdateOrganization) (domain.Organization, error)

	DeleteOrganization(auditID domain.AuditTraceID, profileID domain.ProfileID, organizationID domain.OrganizationID) error
}

func (self UCs) CreateOrganization(auditID domain.AuditTraceID, profileID domain.ProfileID, newOrganization domain.NewOrganization) (domain.Organization, error) {
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
		return domain.Organization{}, err
	}

	err = self.i.Persistence.LinkProfileToOrganization(profileID, organization.ID, domain.RoleOrganizationOwner)

	if err != nil {
		self.i.Trace.Log(auditID, auditStepID).WithFields(domain.LogFields{
			domain.LogFieldError:  err,
			domain.LogFieldUserID: profileID.String()}).
			Msg("Error linking user to organization")
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.Organization{}, err
	}

	self.i.Trace.EndStep(auditID, auditStepID)
	return organization, nil
}

func (self UCs) GetOrganizationByID(auditID domain.AuditTraceID, profileID domain.ProfileID, organizationID domain.OrganizationID) (domain.Organization, error) {
	auditStepID := self.i.Trace.AddStep(auditID, audit.DefaultSkip, profileID, organizationID)

	role, err := self.i.Persistence.GetOrganizationMemberRole(profileID, organizationID)

	if err != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.Organization{}, err
	}

	if err := self.checkPermissions(auditID, profileID, &role, domain.PermissionOrganizationRead); err != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.Organization{}, err
	}

	organization, err := self.i.Persistence.GetOrganizationByIDAndProfileID(profileID, organizationID)

	if err != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.Organization{}, err
	}

	self.i.Trace.EndStep(auditID, auditStepID)
	return organization, nil
}

func (self UCs) UpdateOrganization(auditID domain.AuditTraceID, profileID domain.ProfileID, organizationID domain.OrganizationID, update domain.UpdateOrganization) (domain.Organization, error) {
	auditStepID := self.i.Trace.AddStep(auditID, audit.DefaultSkip, profileID, organizationID, update)

	role, err := self.i.Persistence.GetOrganizationMemberRole(profileID, organizationID)

	if err != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.Organization{}, err
	}

	if err := self.checkPermissions(auditID, profileID, &role, domain.PermissionOrganizationMetaUpdate); err != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return domain.Organization{}, err
	}

	organization, err := self.i.Persistence.GetOrganizationByIDAndProfileID(profileID, organizationID)

	if err != nil {
		return domain.Organization{}, err
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
		return domain.Organization{}, err
	}

	self.i.Trace.EndStep(auditID, auditStepID)
	return organization, nil
}

func (self UCs) DeleteOrganization(auditID domain.AuditTraceID, profileID domain.ProfileID, organizationID domain.OrganizationID) error {
	auditStepID := self.i.Trace.AddStep(auditID, audit.DefaultSkip, profileID, organizationID)

	role, err := self.i.Persistence.GetOrganizationMemberRole(profileID, organizationID)

	if err != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return err
	}

	if err = self.checkPermissions(auditID, profileID, &role, domain.PermissionOrganizationMetaUpdate); err != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return err
	}

	err = self.removeAllMembersFromOrganization(auditID, organizationID)
	if err != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return err
	}

	err = self.i.Persistence.DeleteOrganizationByID(organizationID)
	if err != nil {
		self.i.Trace.EndStep(auditID, auditStepID)
		return err
	}

	self.i.Trace.EndStep(auditID, auditStepID)
	return nil
}
