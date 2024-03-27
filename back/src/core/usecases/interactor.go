package usecases

import (
	"github.com/AliceDiNunno/yeencloud/src/core/interactor"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type UCs struct {
	i *interactor.Interactor
}

func NewUsecases(c interactor.ClusterAdapter, i18n *i18n.Bundle, validator interactor.Validator, audit interactor.Audit, per interactor.Persistence) *UCs {
	ucs := &UCs{
		i: &interactor.Interactor{Cluster: c,
			Translator: i18n,
			Validator:  validator,
			Auditer:    audit,

			Persistence: per,
		},
	}

	// Custom validations.
	// #YC-16 TODO: add better validation system that allows for custom error messages
	ucs.i.Validator.RegisterValidation("unique_email", ucs.UniqueMailValidator)

	return ucs
}

func NewPersistence(user interactor.UserRepository, service interactor.ServiceRepository, session interactor.SessionRepository, profile interactor.ProfileRepository, organization interactor.OrganizationRepository, organizationProfile interactor.OrganizationProfileRepository) interactor.Persistence {
	return interactor.Persistence{
		User:         user,
		Service:      service,
		Session:      session,
		Profile:      profile,
		Organization: organization,

		OrganizationProfile: organizationProfile,
	}
}
