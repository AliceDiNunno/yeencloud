package usecases

import (
	"back/src/core/domain"
	"back/src/core/domain/requests"
)

type LanguageUsecases interface {
	GetAvailableLanguages() []domain.Language
}

type UserUsecases interface {
	CreateUser(user requests.NewUser, language string) (domain.User, *domain.ErrorDescription)
}

type ProfileUsecases interface {
}

type Usecases interface {
	UserUsecases
	ProfileUsecases
	LanguageUsecases
}
