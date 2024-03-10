package domain

import (
	"back/src/core/config/env"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type ApplicationBundle struct {
	Config     *env.Config
	Translator *i18n.Bundle
}
