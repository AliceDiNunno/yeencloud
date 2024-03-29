package usecases

import (
	"back/src/core/domain"
	"back/src/core/domain/requests"
)

type LanguageUsecases interface {
	GetAvailableLanguages() []domain.Language
}

type UserUsecases interface {
	CreateUser(auditID domain.AuditID, user requests.NewUser, language string) (domain.Profile, *domain.ErrorDescription)

	GetUserByID(auditID domain.AuditID, userID domain.UserID) (domain.User, *domain.ErrorDescription)
}

type ProfileUsecases interface {
	GetProfileByUserID(auditID domain.AuditID, userID domain.UserID) (domain.Profile, *domain.ErrorDescription)
}

type SessionUsecases interface {
	CreateSession(auditID domain.AuditID, user requests.NewSession) (domain.Session, *domain.ErrorDescription)

	GetSessionByToken(auditID domain.AuditID, token string) (domain.Session, *domain.ErrorDescription)
}

type OrganizationUsecases interface {
	CreateOrganization(auditID domain.AuditID, userID domain.UserID, organization requests.NewOrganization) (domain.Organization, *domain.ErrorDescription)

	GetOrganizationsByUserID(auditID domain.AuditID, userID domain.UserID) ([]domain.OrganizationMember, *domain.ErrorDescription)
}

type Usecases interface {
	UserUsecases
	ProfileUsecases
	SessionUsecases
	LanguageUsecases
	OrganizationUsecases
}
