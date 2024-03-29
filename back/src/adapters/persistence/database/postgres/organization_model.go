package postgres

import (
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
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
	org := Organization{}

	result := db.engine.Where("id = ?", id.String()).First(&org)

	if result.Error != nil {
		return domain.Organization{}, result.Error
	}

	return organizationToDomain(org), nil
}

func (db *Database) FindOrganizationBySlug(slug string) (domain.Organization, error) {
	org := Organization{}

	result := db.engine.Where("slug = ?", slug).First(&org)

	if result.Error != nil {
		return domain.Organization{}, result.Error
	}

	return organizationToDomain(org), nil
}

func (db *Database) UpdateOrganization(organization domain.OrganizationID, updateOrganization domain.UpdateOrganization) (domain.Organization, error) {
	err := db.engine.Model(&Organization{}).Where("id = ?", organization.String()).Updates(updateOrganization).Error

	if err != nil {
		return domain.Organization{}, err
	}

	return db.FindOrganizationByID(organization)
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
