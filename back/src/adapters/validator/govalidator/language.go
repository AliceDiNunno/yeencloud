package govalidator

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/fr"
	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	fr_translations "github.com/go-playground/validator/v10/translations/fr"
	"github.com/rs/zerolog/log"
)

func (v *Validator) LoadLanguages() {
	enLg := en.New()
	frLg := fr.New()

	// Create a new universal translator using the English locale
	v.validateTranslator = ut.New(enLg, frLg)

	transEn, _ := v.validateTranslator.GetTranslator("en")
	err := en_translations.RegisterDefaultTranslations(v.validate, transEn)

	if err != nil {
		log.Warn().Err(err).Msg("Error registering default translations for English")
	}

	transFr, _ := v.validateTranslator.GetTranslator("fr")
	err = fr_translations.RegisterDefaultTranslations(v.validate, transFr)

	if err != nil {
		log.Warn().Err(err).Msg("Error registering default translations for French")
	}
}
