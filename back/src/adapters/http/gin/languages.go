package gin

import (
	"back/src/core/domain"
	"github.com/gin-gonic/gin"
)

func (server *ServiceHTTPServer) getLangMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lang := ""

		// priority to accept-language header over user profile so an api call can override the set language if needed
		acceptLanguage := ctx.GetHeader(HeaderAcceptLanguage)
		if acceptLanguage != "" {
			lang = acceptLanguage
		} else {
			profileField, exists := ctx.Get(CtxProfileField)
			if exists {
				profile, ok := profileField.(domain.Profile)
				if ok {
					lang = profile.Language
				}
			}
		}

		if lang == "" {
			lang = "enUS"
		}

		ctx.Set(CtxLanguageField, lang)
	}
}
