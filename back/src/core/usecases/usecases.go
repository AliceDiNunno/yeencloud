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

	GetUserByID(id domain.UserID) (domain.User, *domain.ErrorDescription)
}

type ProfileUsecases interface {
	GetProfileByUserID(id domain.UserID) (domain.Profile, *domain.ErrorDescription)
}

type SessionUsecases interface {
	CreateSession(user requests.NewSession) (domain.Session, *domain.ErrorDescription)

	GetSessionByToken(token string) (domain.Session, *domain.ErrorDescription)
}

type OrganizationUsecases interface {
	CreateOrganization(user domain.UserID, organization requests.NewOrganization) (domain.Organization, *domain.ErrorDescription)

	GetOrganizationsByUserID(userID domain.UserID) ([]domain.OrganizationMember, *domain.ErrorDescription)
}

type Usecases interface {
	UserUsecases
	ProfileUsecases
	SessionUsecases
	LanguageUsecases
	OrganizationUsecases
}
