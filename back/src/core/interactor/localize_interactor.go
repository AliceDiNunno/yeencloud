package interactor

import "github.com/AliceDiNunno/yeencloud/src/core/domain"

type Localize interface {
	GetAvailableLanguages() []domain.Language

	GetLocalizedText(language string, key string, params ...map[string]interface{}) string
}
