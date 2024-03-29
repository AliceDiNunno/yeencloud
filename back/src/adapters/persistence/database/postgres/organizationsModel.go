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
	ID   string `gorm:"type:uuid;primary_key"`
	Name string
}

func (o Organization) toDomain() domain.Organization {
	return domain.Organization{
		CloudObject: domain.CloudObject{
			ID:   o.ID,
			Name: o.Name,
		},
	}
}

func (o organizationsRepo) ListOrganizationsByUserID(userID uuid.UUID) (domain.Organization, error) {
	//TODO implement me
	panic("implement me")
}

func OrganizationFromDomain(org domain.Organization) Organization {
	return Organization{
		ID:   org.ID,
		Name: org.Name,
	}
}

func NewOrganizationRepo(db *gorm.DB) usecases.OrganizationRepository {
	return &organizationsRepo{
		db: db,
	}
}
