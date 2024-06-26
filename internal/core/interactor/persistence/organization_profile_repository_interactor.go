package persistence

import "github.com/AliceDiNunno/yeencloud/internal/core/domain"

type OrganizationProfileRepository interface {
	LinkProfileToOrganization(profileID domain.ProfileID, organizationID domain.OrganizationID, role domain.Role) error

	GetOrganizationMemberRole(profileID domain.ProfileID, organizationID domain.OrganizationID) (string, error)
	GetOrganizationByIDAndProfileID(profileID domain.ProfileID, organizationID domain.OrganizationID) (domain.Organization, error)
	ListProfileOrganizationsByProfileID(profileID domain.ProfileID) ([]domain.OrganizationMember, error)
	ListOrganizationMembers(organizationID domain.OrganizationID) ([]domain.OrganizationMember, error)

	RemoveAllMembersFromOrganization(organizationID domain.OrganizationID) error
}
