package usecases

import (
	"back/src/core/domain"
	"back/src/core/domain/requests"
)

type LanguageUsecases interface {
	GetAvailableLanguages() []domain.Language
}

type UserUsecases interface {
	CreateUser(user requests.NewUser, language string) (domain.Profile, *domain.ErrorDescription)
	GetUserByID(id string) (domain.User, *domain.ErrorDescription)
}

type ProfileUsecases interface {
	GetProfileByUserID(id string) (domain.Profile, *domain.ErrorDescription)
}

type SessionUsecases interface {
	CreateSession(user requests.NewSession) (domain.Session, *domain.ErrorDescription)
	GetSessionByToken(token string) (domain.Session, *domain.ErrorDescription)
}

type Usecases interface {
	UserUsecases
	ProfileUsecases
	SessionUsecases
	LanguageUsecases
}
