package usecases

import (
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
)

type LanguageUsecases interface {
	GetAvailableLanguages() []domain.Language
}

type UserUsecases interface {
	CreateUser(auditID domain.AuditID, user domain.NewUser, language string) (domain.Profile, *domain.ErrorDescription)

	GetUserByID(auditID domain.AuditID, userID domain.UserID) (domain.User, *domain.ErrorDescription)
}

type ProfileUsecases interface {
	GetProfileByUserID(auditID domain.AuditID, userID domain.UserID) (domain.Profile, *domain.ErrorDescription)
}

type SessionUsecases interface {
	CreateSession(auditID domain.AuditID, user domain.NewSession) (domain.Session, *domain.ErrorDescription)

	GetSessionByToken(auditID domain.AuditID, token string) (domain.Session, *domain.ErrorDescription)
}

type OrganizationUsecases interface {
	CreateOrganization(auditID domain.AuditID, profileID domain.ProfileID, organization domain.NewOrganization) (domain.Organization, *domain.ErrorDescription)

	GetOrganizationsByProfileID(auditID domain.AuditID, profileID domain.ProfileID) ([]domain.OrganizationMember, *domain.ErrorDescription)
}

type Usecases interface {
	UserUsecases
	ProfileUsecases
	SessionUsecases
	LanguageUsecases
	OrganizationUsecases
}
