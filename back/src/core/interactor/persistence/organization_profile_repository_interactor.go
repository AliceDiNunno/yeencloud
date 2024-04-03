package persistence

import "github.com/AliceDiNunno/yeencloud/src/core/domain"

type OrganizationProfileRepository interface {
	LinkProfileToOrganization(profileID domain.ProfileID, organizationID domain.OrganizationID, role domain.OrganizationRole) error

	GetOrganizationByIDAndProfileID(profileID domain.ProfileID, organizationID domain.OrganizationID) (domain.Organization, error)
	ListProfileOrganizationsByProfileID(profileID domain.ProfileID) ([]domain.OrganizationMember, error)
	ListOrganizationMembers(organizationID domain.OrganizationID) ([]domain.OrganizationMember, error)
}
