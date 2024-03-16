package postgres

import (
	"back/src/core/domain"
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model

	ID string `gorm:"type:uuid;primary_key"`

	UserID string `gorm:"foreignkey:Profile;not null;unique;default:null;<-:create"`
	User   User

	Name     string
	Language string

	Organizations []OrganizationProfile
}

func (db *Database) CreateProfile(profile domain.Profile) (domain.Profile, error) {
	profileToCreate := domainToProfile(profile)

	result := db.engine.Create(&profileToCreate)

	if result.Error != nil {
		return domain.Profile{}, result.Error
	}

	return profileToDomain(profileToCreate), nil
}

func (db *Database) FindProfileByUserID(userID domain.UserID) (domain.Profile, error) {
	var profile Profile

	result := db.engine.Where("user_id = ?", userID).First(&profile)

	if result.Error != nil {
		return domain.Profile{}, result.Error
	}

	return profileToDomain(profile), nil
}

func domainToProfile(profile domain.Profile) Profile {
	return Profile{
		ID:       profile.ID.String(),
		UserID:   profile.UserID.String(),
		Name:     profile.Name,
		Language: profile.Language,
	}
}

func profileToDomain(profile Profile) domain.Profile {
	return domain.Profile{
		ID:       domain.ProfileID(profile.ID),
		UserID:   domain.UserID(profile.UserID),
		Name:     profile.Name,
		Language: profile.Language,
	}
}
