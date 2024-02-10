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

	FindUserByID(id uuid.UUID) (domain.User, error)
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
	validator         Validator

	translator *i18n.Bundle
}

func NewInteractor(c ClusterAdapter, sR SettingsRepository, uR UserRepository, servR ServiceRepository,
	i18n *i18n.Bundle, validator Validator) *interactor {
	inter := &interactor{
		cluster: c,

		settingsRepo: sR,
		userRepo:     uR,
		serviceRepo:  servR,

		translator: i18n,
		validator:  validator,
	}

	inter.validator.AddCustomValidation("password", inter.PasswordValidator())

	return inter
}
