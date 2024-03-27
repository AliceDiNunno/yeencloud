package interactor

import "github.com/nicksnyder/go-i18n/v2/i18n"

type Interactor struct {
	Log Logger

	Cluster    ClusterAdapter
	Validator  Validator
	Translator *i18n.Bundle
	Auditer    Audit

	Persistence Persistence
}
