package usecases

import (
	"back/src/core/domain"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type SettingsRepository interface {
	GetSettingsValue(key string) string
	SetSettingsValue(key string, value string)
}

type UserRepository interface {
	CountUsers() int64

	CreateUser(user domain.User) (domain.User, error)

	FindUserByID(id uuid.UUID) (domain.User, error)
	FindUserByEmail(email string) (domain.User, error)
}

type ProfileRepository interface {
	CreateProfile(profile domain.Profile) (domain.Profile, error)
}

type OrganizationRepository interface {
	ListOrganizationsByUserID(userID uuid.UUID) (domain.Organization, error)
}

type ServiceRepository interface {
}

type SessionRepository interface {
}

type ClusterAdapter interface {
	IsRunningInsideCluster() bool
	IsConfigurationValid([]byte) bool
}

type Validator interface {
	ValidateStruct(s interface{}) (bool, domain.ValidationErrors)
	ValidateStructWithLang(s interface{}, lang string) (bool, domain.ValidationErrors)
	AddCustomValidation(tag string, fn validator.Func)
}

type interactor struct {
	cluster ClusterAdapter

	settingsRepo      SettingsRepository
	userRepo          UserRepository
	serviceRepo       ServiceRepository
	sessionRepository SessionRepository
	profileRepo       ProfileRepository
	validator         Validator

	translator *i18n.Bundle
}

func NewInteractor(c ClusterAdapter, sR SettingsRepository, uR UserRepository, proR ProfileRepository, servR ServiceRepository,
	i18n *i18n.Bundle, validator Validator) *interactor {
	inter := &interactor{
		cluster: c,

		settingsRepo: sR,
		userRepo:     uR,
		profileRepo:  proR,
		serviceRepo:  servR,

		translator: i18n,
		validator:  validator,
	}

	// custom validations
	// TODO: add better validation system that allows for custom error messages
	inter.validator.AddCustomValidation("password", inter.PasswordValidator())
	inter.validator.AddCustomValidation("unique_email", inter.UniqueMailValidator())

	return inter
}
