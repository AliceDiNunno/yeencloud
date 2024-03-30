package localization

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/AliceDiNunno/yeencloud/src/core/interactor"
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var LogFieldLanguage = domain.LogField{Name: "language"}
var LogFieldLanguageTag = domain.LogField{Parent: &LogFieldLanguage, Name: "tag"}
var LogFieldLanguageName = domain.LogField{Parent: &LogFieldLanguage, Name: "name"}
var LogFieldLanguageMessage = domain.LogField{Parent: &LogFieldLanguage, Name: "message"}

type localize struct {
	bundle *i18n.Bundle

	defaultLanguage language.Tag
	logger          interactor.Logger
}

func (l localize) GetAvailableLanguages() []domain.Language {
	tags := l.bundle.LanguageTags()

	languages := []domain.Language{}
	for _, tag := range tags {
		localizer := i18n.NewLocalizer(l.bundle, tag.String())

		displayName, err := localizer.Localize(&i18n.LocalizeConfig{DefaultMessage: &i18n.Message{ID: "LanguageDisplayName"}})
		if err != nil {
			l.logger.Log(domain.LogLevelError).WithField(LogFieldLanguageTag, tag).WithField(domain.LogFieldError, err).Msg("Error getting language display name")
			continue
		}

		flag, err := localizer.Localize(&i18n.LocalizeConfig{DefaultMessage: &i18n.Message{ID: "LanguageFlag"}})
		if err != nil {
			l.logger.Log(domain.LogLevelError).WithField(LogFieldLanguageTag, tag).WithField(LogFieldLanguageName, displayName).WithField(domain.LogFieldError, err).Msg("Error getting language flag")
			continue
		}

		languages = append(languages, domain.Language{
			Tag:         tag.String(),
			Flag:        flag,
			DisplayName: displayName,
		})
	}

	return languages
}

func (l localize) translatableArgumentsToMap(params []domain.TranslatableArgumentMap) map[string]interface{} {
	result := map[string]interface{}{}
	for _, arg := range params {
		for key, value := range arg {
			result[key.Key] = value
		}
	}

	return result
}

func (l localize) GetLocalizedText(language string, translatable domain.Translatable, params ...domain.TranslatableArgumentMap) string {
	msg := i18n.NewLocalizer(l.bundle, language)

	if language == "" {
		language = l.defaultLanguage.String()
	}

	var isDefault bool = false
	languages := l.GetAvailableLanguages()
	for _, lang := range languages {
		if lang.Tag == language {
			isDefault = lang.Tag == l.defaultLanguage.String()
			break
		}
	}

	key := translatable.Key
	localizedMessage, _, err := msg.LocalizeWithTag(&i18n.LocalizeConfig{
		MessageID:    key,
		TemplateData: l.translatableArgumentsToMap(params),
	})

	if err != nil {
		l.logger.Log(domain.LogLevelError).
			WithField(LogFieldLanguageTag, language).
			WithField(LogFieldLanguageMessage, key).
			WithField(domain.LogFieldError, err).
			Msg("Error getting localized text")

		if !isDefault {
			return l.GetLocalizedText(l.defaultLanguage.String(), translatable, params...)
		}

		return key
	}

	return localizedMessage
}

func NewLocalize(logger interactor.Logger, localizationPath string, defaultLanguage language.Tag) interactor.Localize {
	bundle := i18n.NewBundle(defaultLanguage)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	files, err := os.ReadDir(localizationPath)
	if err != nil {
		log.Fatalln("Error reading translation directory " + err.Error())
	}

	for _, file := range files {
		path := fmt.Sprintf("%s/%s", localizationPath, file.Name())

		if strings.HasSuffix(path, ".toml") {
			bundle.MustLoadMessageFile(path)
		}
	}

	return &localize{
		bundle:          bundle,
		logger:          logger,
		defaultLanguage: defaultLanguage,
	}
}
