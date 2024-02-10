package gin

import (
	"back/src/core/domain"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func (server *ServiceHTTPServer) validate(i interface{}, c *gin.Context) bool {
	lang := c.GetString("lang")
	msg := i18n.NewLocalizer(server.translator, lang)
	shortLanguage, _, _ := msg.LocalizeWithTag(&i18n.LocalizeConfig{
		MessageID: "ShortLanguageCode",
	})

	succeeded, validationErrors := server.validator.ValidateStructWithLang(i, shortLanguage)

	if succeeded {
		return true
	}

	err := domain.ErrorBadRequest

	localized, _, _ := msg.LocalizeWithTag(&i18n.LocalizeConfig{
		MessageID: err.Code,
	})

	c.AbortWithStatusJSON(err.HttpCode, Response{
		StatusCode: err.HttpCode,
		Error: &ResponseError{
			Code:    err.Code,
			Message: localized,
		},
		Body: validationErrors,
	})
	return false
}
