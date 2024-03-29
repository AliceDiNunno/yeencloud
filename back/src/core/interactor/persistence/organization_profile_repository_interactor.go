package persistence

import "github.com/AliceDiNunno/yeencloud/src/core/domain"

type OrganizationProfileRepository interface {
	LinkProfileToOrganization(profileID domain.ProfileID, organizationID domain.OrganizationID, role domain.OrganizationRole) error

	GetProfileOrganizationsByProfileID(profileID domain.ProfileID) ([]domain.OrganizationMember, error)
	GetOrganizationMembers(organizationID domain.OrganizationID) ([]domain.OrganizationMember, error)
}
