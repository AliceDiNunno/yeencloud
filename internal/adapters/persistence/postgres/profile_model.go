package postgres

import (
	"github.com/AliceDiNunno/yeencloud/internal/core/domain"
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model

	ID string `gorm:"type:uuid;primary_key"`

	UserID string `gorm:"foreignkey:Profile;not null;unique;default:null;<-:create"`
	User   User

	Name     string
	Language string

	Role string

	Organizations []OrganizationProfile
}

func (db *Database) CreateProfile(profile domain.Profile) (domain.Profile, error) {
	profileToCreate := domainToProfile(profile)

	result := db.engine.Create(&profileToCreate)

	if result.Error != nil {
		return domain.Profile{}, sqlerr(result.Error)
	}

	return profileToDomain(profileToCreate), nil
}

func (db *Database) FindProfileByID(profileID domain.ProfileID) (domain.Profile, error) {
	var profile Profile

	result := db.engine.Where("id = ?", profileID).First(&profile)

	if result.Error != nil {
		return domain.Profile{}, sqlerr(result.Error)
	}

	return profileToDomain(profile), nil
}

func (db *Database) FindProfileByUserID(userID domain.UserID) (domain.Profile, error) {
	var profile Profile

	result := db.engine.Where("user_id = ?", userID).First(&profile)

	if result.Error != nil {
		return domain.Profile{}, sqlerr(result.Error)
	}

	return profileToDomain(profile), nil
}

func (db *Database) SetProfileRole(profileID domain.ProfileID, role domain.Role) error {
	result := db.engine.Model(&Profile{}).Where("id = ?", profileID).Update("role", role.String())

	if result.Error != nil {
		return sqlerr(result.Error)
	}

	return nil
}

func domainToProfile(profile domain.Profile) Profile {
	return Profile{
		ID:       profile.ID.String(),
		UserID:   profile.UserID.String(),
		Name:     profile.Name,
		Language: profile.Language,
		Role:     profile.Role,
	}
}

func profileToDomain(profile Profile) domain.Profile {
	return domain.Profile{
		ID:       domain.ProfileID(profile.ID),
		UserID:   domain.UserID(profile.UserID),
		Name:     profile.Name,
		Language: profile.Language,
		Role:     profile.Role,
	}
}
