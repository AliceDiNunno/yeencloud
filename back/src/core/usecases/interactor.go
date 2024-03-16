package usecases

import (
	"back/src/core/domain"
	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Audit interface {
	NewTrace(trigger string, data ...interface{}) domain.AuditID
	AddStep(id domain.AuditID, details ...interface{})
	EndTrace(id domain.AuditID, result ...interface{})
}

type UserRepository interface {
	CreateUser(user domain.User) (domain.User, error)

	CountUsers() int64
	FindUserByID(userID domain.UserID) (domain.User, error)
	FindUserByEmail(email string) (domain.User, error)
}

type ProfileRepository interface {
	CreateProfile(profile domain.Profile) (domain.Profile, error)

	FindProfileByUserID(userID domain.UserID) (domain.Profile, error)
}

type OrganizationRepository interface {
	CreateOrganization(organization domain.Organization) (domain.Organization, error)
	DeleteOrganizationByID(id domain.OrganizationID) error
}

type OrganizationProfileRepository interface {
	LinkProfileToOrganization(profileID domain.ProfileID, organizationID domain.OrganizationID, role domain.OrganizationRole) error

	GetProfileOrganizationsByProfileID(profileID domain.ProfileID) ([]domain.OrganizationMember, error)
	GetOrganizationMembers(organizationID domain.OrganizationID) ([]domain.OrganizationMember, error)
}

type ServiceRepository interface {
}

type SessionRepository interface {
	CreateSession(session domain.Session) (domain.Session, error)

	FindSessionByToken(token string) (domain.Session, error)
}

type ClusterAdapter interface {
	IsRunningInsideCluster() bool
	IsConfigurationValid(ClusterConfiguration []byte) bool
}

type Validator interface {
	ValidateStruct(s interface{}) (bool, domain.ValidationErrors)
	ValidateStructWithLang(s interface{}, lang string) (bool, domain.ValidationErrors)
	AddCustomValidation(tag string, fn validator.Func)
}

type Persistence struct {
	// Main models
	user         UserRepository
	service      ServiceRepository
	session      SessionRepository
	profile      ProfileRepository
	organization OrganizationRepository

	// Linking models
	organizationProfile OrganizationProfileRepository
}

type Interactor struct {
	cluster    ClusterAdapter
	validator  Validator
	translator *i18n.Bundle
	auditer    Audit

	persistence Persistence
}

func NewInteractor(c ClusterAdapter, i18n *i18n.Bundle, validator Validator, audit Audit, per Persistence) *Interactor {
	inter := &Interactor{
		cluster:    c,
		translator: i18n,
		validator:  validator,
		auditer:    audit,

		persistence: per,
	}

	// Custom validations.
	// #YC-16 TODO: add better validation system that allows for custom error messages
	inter.validator.AddCustomValidation("password", inter.PasswordValidator())
	inter.validator.AddCustomValidation("unique_email", inter.UniqueMailValidator())

	return inter
}

func NewPersistence(user UserRepository, service ServiceRepository, session SessionRepository, profile ProfileRepository, organization OrganizationRepository, organizationProfile OrganizationProfileRepository) Persistence {
	return Persistence{
		user:         user,
		service:      service,
		session:      session,
		profile:      profile,
		organization: organization,

		organizationProfile: organizationProfile,
	}
}
