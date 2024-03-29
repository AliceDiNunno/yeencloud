package persistence

import "github.com/AliceDiNunno/yeencloud/src/core/domain"

type UserRepository interface {
	CreateUser(user domain.User) (domain.User, error)

	CountUsers() int64
	FindUserByID(userID domain.UserID) (domain.User, error)
	FindUserByEmail(email string) (domain.User, error)
}
