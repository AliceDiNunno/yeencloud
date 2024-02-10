package usecases

import (
	"back/src/core/domain"
	"back/src/core/domain/requests"
)

type LanguageUsecases interface {
	GetAvailableLanguages() []domain.Language
}

type UserUsecases interface {
	CreateUser(user requests.NewUser) (domain.User, *domain.ErrorDescription)
}

type Usecases interface {
	UserUsecases
	LanguageUsecases
}
