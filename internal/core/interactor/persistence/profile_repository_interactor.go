package persistence

import "github.com/AliceDiNunno/yeencloud/internal/core/domain"

type ProfileRepository interface {
	CreateProfile(profile domain.Profile) (domain.Profile, error)

	FindProfileByID(profileID domain.ProfileID) (domain.Profile, error)
	FindProfileByUserID(userID domain.UserID) (domain.Profile, error)

	SetProfileRole(profileID domain.ProfileID, role domain.Role) error
}
