package persistence

import "github.com/AliceDiNunno/yeencloud/src/core/domain"

type OrganizationRepository interface {
	CreateOrganization(organization domain.Organization) (domain.Organization, error)

	UpdateOrganization(organization domain.OrganizationID, update domain.UpdateOrganization) (domain.Organization, error)

	DeleteOrganizationByID(id domain.OrganizationID) error
}
