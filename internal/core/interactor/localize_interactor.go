package interactor

import "github.com/AliceDiNunno/yeencloud/internal/core/domain"

type Localize interface {
	GetAvailableLanguages() []domain.Language
	DefaultLanguageName() string

	GetLocalizedText(language string, key domain.Translatable, params ...domain.TranslatableArgumentMap) string
}
