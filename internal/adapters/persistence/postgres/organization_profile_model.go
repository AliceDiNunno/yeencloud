package postgres

import (
	"github.com/AliceDiNunno/yeencloud/internal/core/domain"
	"gorm.io/gorm"
)

type OrganizationProfile struct {
	gorm.Model

	OrganizationID string
	Organization   Organization

	ProfileID string
	Profile   Profile

	UserRole string
}

func (db *Database) LinkProfileToOrganization(profileID domain.ProfileID, organizationID domain.OrganizationID, role domain.Role) error {
	NewLink := OrganizationProfile{
		OrganizationID: organizationID.String(),
		ProfileID:      profileID.String(),
		UserRole:       role.String(),
	}

	result := db.engine.Create(&NewLink)

	return sqlerr(result.Error)
}

func (db *Database) ListProfileOrganizationsByProfileID(profileID domain.ProfileID) ([]domain.OrganizationMember, error) {
	var orgs []OrganizationProfile

	result := db.engine.Preload("Profile").Preload("Organization").Where("profile_id = ?", profileID).Find(&orgs)

	if result.Error != nil {
		return nil, sqlerr(result.Error)
	}

	return organizationMembersToDomain(orgs), nil
}

func (db *Database) ListOrganizationMembers(orgID domain.OrganizationID) ([]domain.OrganizationMember, error) {
	var users []OrganizationProfile

	result := db.engine.Where("organization_id = ?", orgID).Find(&users)

	if result.Error != nil {
		return nil, sqlerr(result.Error)
	}

	return organizationMembersToDomain(users), nil
}

func (db *Database) GetOrganizationByIDAndProfileID(profileID domain.ProfileID, organizationID domain.OrganizationID) (domain.Organization, error) {
	var org OrganizationProfile

	result := db.engine.Preload("Organization").Where("profile_id = ? AND organization_id = ?", profileID, organizationID).First(&org)

	if result.Error != nil {
		return domain.Organization{}, sqlerr(result.Error)
	}

	return organizationToDomain(org.Organization), nil
}

func (db *Database) GetOrganizationMemberRole(profileID domain.ProfileID, organizationID domain.OrganizationID) (string, error) {
	organizationProfile := OrganizationProfile{}

	result := db.engine.Where("profile_id = ? AND organization_id = ?", profileID.String(), organizationID.String()).First(&organizationProfile)

	if result.Error != nil {
		return "", sqlerr(result.Error)
	}

	return organizationProfile.UserRole, nil
}

func (db *Database) RemoveAllMembersFromOrganization(organizationID domain.OrganizationID) error {
	result := db.engine.Where("organization_id = ?", organizationID).Delete(&OrganizationProfile{})

	return sqlerr(result.Error)
}

func organizationMembersToDomain(organizationProfiles []OrganizationProfile) []domain.OrganizationMember {
	var result []domain.OrganizationMember
	for _, profile := range organizationProfiles {
		result = append(result, organizationMemberToDomain(profile))
	}
	return result
}

func organizationMemberToDomain(organizationProfile OrganizationProfile) domain.OrganizationMember {
	return domain.OrganizationMember{
		Profile:      profileToDomain(organizationProfile.Profile),
		Organization: organizationToDomain(organizationProfile.Organization),
		Role:         organizationProfile.UserRole,
	}
}
