package interactor

import (
	"github.com/AliceDiNunno/yeencloud/src/core/interactor/persistence"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Interactor struct {
	Log Logger

	Cluster    ClusterAdapter
	Validator  Validator
	Translator *i18n.Bundle
	Auditer    Audit

	Persistence persistence.Persistence
}
