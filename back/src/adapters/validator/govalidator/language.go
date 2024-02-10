package govalidator

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/fr"
	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	fr_translations "github.com/go-playground/validator/v10/translations/fr"
)

func (v *Validator) LoadLanguages() {
	en_lg := en.New()
	fr_lg := fr.New()

	// Create a new universal translator using the English locale
	v.validateTranslator = ut.New(en_lg, fr_lg)

	transEn, _ := v.validateTranslator.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(v.validate, transEn)

	transFr, _ := v.validateTranslator.GetTranslator("fr")
	fr_translations.RegisterDefaultTranslations(v.validate, transFr)

}
