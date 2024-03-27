package postgres

import "github.com/AliceDiNunno/yeencloud/src/core/domain"

type OrganizationProfile struct {
	OrganizationID string
	Organization   Organization
	ProfileID      string
	Profile        Profile
	UserRole       string
}

func (db *Database) LinkProfileToOrganization(profileID domain.ProfileID, organizationID domain.OrganizationID, role domain.OrganizationRole) error {
	NewLink := OrganizationProfile{
		OrganizationID: organizationID.String(),
		ProfileID:      profileID.String(),
		UserRole:       role.String(),
	}

	result := db.engine.Create(&NewLink)

	return result.Error
}

func (db *Database) GetProfileOrganizationsByProfileID(profileID domain.ProfileID) ([]domain.OrganizationMember, error) {
	var orgs []OrganizationProfile

	result := db.engine.Preload("Profile").Preload("Organization").Where("profile_id = ?", profileID).Find(&orgs)

	if result.Error != nil {
		return nil, result.Error
	}

	return organizationMembersToDomain(orgs), nil
}

func (db *Database) GetOrganizationMembers(orgID domain.OrganizationID) ([]domain.OrganizationMember, error) {
	var users []OrganizationProfile

	result := db.engine.Where("organization_id = ?", orgID).Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return organizationMembersToDomain(users), nil
}

func organizationMembersToDomain(profiles []OrganizationProfile) []domain.OrganizationMember {
	var result []domain.OrganizationMember
	for _, profile := range profiles {
		result = append(result, organizationMemberToDomain(profile))
	}
	return result
}

func organizationMemberToDomain(user OrganizationProfile) domain.OrganizationMember {
	return domain.OrganizationMember{
		Profile:      profileToDomain(user.Profile),
		Organization: organizationToDomain(user.Organization),
		Role:         domain.OrganizationRole(user.UserRole),
	}
}
