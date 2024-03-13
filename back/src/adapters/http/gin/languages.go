package gin

import (
	"back/src/core/domain"
	"github.com/gin-gonic/gin"
)

// #YC-8 TODO: move to status route
func (server *ServiceHTTPServer) getLanguagesHandler(ctx *gin.Context) {
	languages := server.ucs.GetAvailableLanguages()

	server.success(ctx, languages)
}

func (server *ServiceHTTPServer) getLangMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lang := ""

		// priority to accept-language header over user profile so an api call can override the set language if needed
		acceptLanguage := ctx.GetHeader("Accept-Language")
		if acceptLanguage != "" {
			lang = acceptLanguage
		} else {
			userID, exists := ctx.Get("user")
			if exists {
				user, err := server.ucs.GetProfileByUserID(server.getTrace(ctx), userID.(domain.UserID))
				if err == nil {
					lang = user.Language
				}
			}
		}

		if lang == "" {
			lang = "enUS"
		}

		ctx.Set("lang", lang)
	}
}
