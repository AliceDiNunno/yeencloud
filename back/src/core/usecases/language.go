package usecases

import (
	"back/src/core/domain"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func (i interactor) GetAvailableLanguages() []domain.Language {
	tags := i.translator.LanguageTags()

	languages := []domain.Language{}
	for _, tag := range tags {
		localizer := i18n.NewLocalizer(i.translator, tag.String())
		displayName := localizer.MustLocalize(&i18n.LocalizeConfig{DefaultMessage: &i18n.Message{ID: "LanguageDisplayName"}})
		flag := localizer.MustLocalize(&i18n.LocalizeConfig{DefaultMessage: &i18n.Message{ID: "LanguageFlag"}})

		languages = append(languages, domain.Language{
			Identifier:  tag.String(),
			Flag:        flag,
			DisplayName: displayName,
		})
	}

	return languages
}
