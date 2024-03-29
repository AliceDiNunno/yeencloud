package postgres

import (
	"back/src/core/domain"
	"back/src/core/usecases"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type organizationsRepo struct {
	db *gorm.DB
}

type Organization struct {
	CloudObject
}

func (o Organization) toDomain() domain.Organization {
	return domain.Organization{
		CloudObject: domain.CloudObject{
			ID:   o.CloudObject.ID,
			Name: o.CloudObject.Name,
		},
	}
}

func (o organizationsRepo) ListOrganizationsByUserID(userID uuid.UUID) (domain.Organization, error) {
	//TODO implement me
	panic("implement me")
}

func OrganizationFromDomain(org domain.Organization) Organization {
	return Organization{
		CloudObject: CloudObject{
			ID:   org.ID,
			Name: org.Name,
		},
	}
}

func NewOrganizationRepo(db *gorm.DB) usecases.OrganizationRepository {
	return &organizationsRepo{
		db: db,
	}
}
