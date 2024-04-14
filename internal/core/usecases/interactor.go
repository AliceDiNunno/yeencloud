package usecases

import (
	"time"

	"github.com/AliceDiNunno/yeencloud/internal/core/interactor"
	persistenceInteractor "github.com/AliceDiNunno/yeencloud/internal/core/interactor/persistence"
)

type UCs struct {
	i *interactor.Interactor

	requestTimer *time.Timer
}

func NewUsecases(c interactor.ClusterAdapter, mailer interactor.Mailer, i18n interactor.Localize, validator interactor.Validator, audit interactor.Audit, per persistenceInteractor.Persistence) *UCs {
	ucs := &UCs{
		i: &interactor.Interactor{Cluster: c,
			Localize:  i18n,
			Validator: validator,
			Trace:     audit,
			Mailer:    mailer,

			Persistence: per,
		},
	}

	// Custom validations.
	// #YC-16 TODO: add better validation system that allows for custom error messages
	ucs.i.Validator.RegisterValidation("unique_email", ucs.UniqueMailValidator)

	return ucs
}
