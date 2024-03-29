package persistence

import "github.com/AliceDiNunno/yeencloud/src/core/domain"

type ProfileRepository interface {
	CreateProfile(profile domain.Profile) (domain.Profile, error)

	FindProfileByUserID(userID domain.UserID) (domain.Profile, error)
}
