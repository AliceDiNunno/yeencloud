package usecases

import (
	"github.com/AliceDiNunno/yeencloud/src/core/interactor"
	persistenceInteractor "github.com/AliceDiNunno/yeencloud/src/core/interactor/persistence"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type UCs struct {
	i *interactor.Interactor
}

func NewPersistence(user persistenceInteractor.UserRepository, service persistenceInteractor.ServiceRepository, session persistenceInteractor.SessionRepository, profile persistenceInteractor.ProfileRepository, organization persistenceInteractor.OrganizationRepository, organizationProfile persistenceInteractor.OrganizationProfileRepository) persistenceInteractor.Persistence {
	return persistenceInteractor.Persistence{
		User:         user,
		Service:      service,
		Session:      session,
		Profile:      profile,
		Organization: organization,

		OrganizationProfile: organizationProfile,
	}
}

func NewUsecases(c interactor.ClusterAdapter, i18n *i18n.Bundle, validator interactor.Validator, audit interactor.Audit, per persistenceInteractor.Persistence) *UCs {
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
