package postgres

import (
	"back/src/core/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type profileRepo struct {
	db *gorm.DB
}

type Profile struct {
	gorm.Model

	UserID string `gorm:"type:uuid;foreignkey:User;not null;unique"`
	User   User

	Name     string
	Language string
}

func (db *Database) CreateProfile(profile domain.Profile) (domain.Profile, error) {
	profileToCreate := domainToProfile(profile)

	result := db.engine.Create(&profileToCreate)

	if result.Error != nil {
		return domain.Profile{}, result.Error
	}

	return profileToDomain(profileToCreate), nil
}

func (db *Database) FindProfileByUserID(userID uuid.UUID) (domain.Profile, error) {
	var profile Profile

	result := db.engine.Where("user_id = ?", userID).First(&profile)

	if result.Error != nil {
		return domain.Profile{}, result.Error
	}

	return profileToDomain(profile), nil
}

func domainToProfile(profile domain.Profile) Profile {
	return Profile{
		UserID:   profile.UserID,
		Name:     profile.Name,
		Language: profile.Language,
	}
}

func profileToDomain(profile Profile) domain.Profile {
	return domain.Profile{
		UserID:   profile.UserID,
		Name:     profile.Name,
		Language: profile.Language,
	}
}
