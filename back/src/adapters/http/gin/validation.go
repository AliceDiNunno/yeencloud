package gin

import (
	"back/src/core/domain"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func (server *ServiceHTTPServer) validate(i interface{}, c *gin.Context) bool {
	server.auditer.AddStep(server.getTrace(c))

	lang := c.GetString("lang")
	msg := i18n.NewLocalizer(server.translator, lang)
	shortLanguage, _, _ := msg.LocalizeWithTag(&i18n.LocalizeConfig{
		MessageID: "ShortLanguageCode",
	})

	succeeded, validationErrors := server.validator.ValidateStructWithLang(i, shortLanguage)

	if succeeded {
		return true
	}

	server.abortWithError(c, domain.ErrorBadRequest, validationErrors)

	return false
}
