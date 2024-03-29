package postgres

import "back/src/core/domain"

type OrganizationUser struct {
	OrganizationID string
	Organization   Organization
	UserID         string
	User           User
	UserRole       string
}

func (db *Database) LinkUserToOrganization(userID domain.UserID, organizationID domain.OrganizationID, role domain.OrganizationRole) error {
	NewLink := OrganizationUser{
		OrganizationID: organizationID.String(),
		UserID:         userID.String(),
		UserRole:       role.String(),
	}

	result := db.engine.Create(&NewLink)

	return result.Error
}

func (db *Database) GetUserOrganizationsByUserID(userID domain.UserID) ([]domain.OrganizationMember, error) {
	var orgs []OrganizationUser

	result := db.engine.Preload("User").Preload("Organization").Where("user_id = ?", userID).Find(&orgs)

	if result.Error != nil {
		return nil, result.Error
	}

	return organizationMembersToDomain(orgs), nil
}

func (db *Database) GetOrganizationMembers(orgID domain.OrganizationID) ([]domain.OrganizationMember, error) {
	var users []OrganizationUser

	result := db.engine.Where("organization_id = ?", orgID).Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return organizationMembersToDomain(users), nil
}

func organizationMembersToDomain(users []OrganizationUser) []domain.OrganizationMember {
	var result []domain.OrganizationMember
	for _, user := range users {
		result = append(result, organizationMemberToDomain(user))
	}
	return result
}

func organizationMemberToDomain(user OrganizationUser) domain.OrganizationMember {
	return domain.OrganizationMember{
		User:         userToDomain(user.User),
		Organization: organizationToDomain(user.Organization),
		Role:         domain.OrganizationRole(user.UserRole),
	}
}
