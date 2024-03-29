package domain

import (
	"back/src/core/config"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type ApplicationBundle struct {
	Config     *config.Config
	Translator *i18n.Bundle
}
