package gin

import (
	"back/src/core/domain"
	"back/src/core/domain/requests"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
)

func (server *ServiceHTTPServer) getUserMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (server *ServiceHTTPServer) getLangMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		_, exists := context.Get("user")
		if !exists {
			acceptLanguage := context.GetHeader("Accept-Language")
			if acceptLanguage != "" {
				context.Set("lang", acceptLanguage)
				return
			}
			context.Set("lang", "enUS")
		}
	}
}

func (server *ServiceHTTPServer) createUserHandler(context *gin.Context) {
	var createUserRequest requests.NewUser
	if err := context.ShouldBindJSON(&createUserRequest); err != nil {
		server.abortWithError(context, domain.ErrorBadRequest)
		return
	}

	if !server.validate(createUserRequest, context) {
		return
	}

	language := context.GetString("lang")
	spew.Dump(language)
	user, err := server.ucs.CreateUser(createUserRequest, language)

	if err != nil {
		server.abortWithError(context, *err)
		return
	}

	context.JSON(201, user)
}
