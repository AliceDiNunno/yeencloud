package gin

import (
	"back/src/core/domain"
	"github.com/gin-gonic/gin"
)

// #YC-8 TODO: move to status route
func (server *ServiceHTTPServer) getLanguagesHandler(context *gin.Context) {
	languages := server.ucs.GetAvailableLanguages()

	context.JSON(200, languages)
}

func (server *ServiceHTTPServer) getLangMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		lang := ""

		// priority to accept-language header over user profile so an api call can override the set language if needed
		acceptLanguage := context.GetHeader("Accept-Language")
		if acceptLanguage != "" {
			lang = acceptLanguage
		} else {
			userID, exists := context.Get("user")
			if exists {
				user, err := server.ucs.GetProfileByUserID(userID.(domain.UserID))
				if err == nil {
					lang = user.Language
				}
			}
		}

		if lang == "" {
			lang = "enUS"
		}

		context.Set("lang", lang)
	}
}
