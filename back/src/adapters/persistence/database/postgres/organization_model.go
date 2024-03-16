package postgres

import (
	"back/src/core/domain"
)

type Organization struct {
	ID          string `gorm:"primary_key;unique;not null;default:null;<-:create"`
	Slug        string `gorm:"unique;not null;default:null;<-:create"`
	Name        string
	Description string

	Users []OrganizationProfile
}

func (db *Database) CreateOrganization(organization domain.Organization) (domain.Organization, error) {
	organizationToCreate := domainToOrganization(organization)

	result := db.engine.Create(&organizationToCreate)

	if result.Error != nil {
		return domain.Organization{}, result.Error
	}

	return organizationToDomain(organizationToCreate), nil
}

func (db *Database) FindOrganizationByID(id domain.OrganizationID) (domain.Organization, error) {
	panic("implement me")
}

func (db *Database) FindOrganizationBySlug(slug string) (domain.Organization, error) {
	panic("implement me")
}

func (db *Database) DeleteOrganizationByID(id domain.OrganizationID) error {
	// #YC-10 TODO implement me
	panic("implement me")
}

func domainToOrganization(org domain.Organization) Organization {
	return Organization{
		ID:          org.ID.String(),
		Slug:        org.Slug,
		Name:        org.Name,
		Description: org.Description,
		Users:       nil,
	}
}

func organizationToDomain(org Organization) domain.Organization {
	return domain.Organization{
		ID:          domain.OrganizationID(org.ID),
		Slug:        org.Slug,
		Name:        org.Name,
		Description: org.Description,
	}
}
