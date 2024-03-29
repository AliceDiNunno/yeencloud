package gin

import (
	"back/src/core/domain"
	"back/src/core/domain/requests"
	"github.com/gin-gonic/gin"
)

func (server *ServiceHTTPServer) getUserMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		session, exists := context.Get("session")

		if !exists {
			return
		}

		//This is necessary to make sure the user is still valid
		user, err := server.ucs.GetUserByID(session.(domain.Session).UserID)

		if err != nil {
			server.abortWithError(context, *err)
			return
		}

		context.Set("user", user.ID)
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

	profile, err := server.ucs.CreateUser(createUserRequest, language)

	if err != nil {
		server.abortWithError(context, *err)
		return
	}

	context.JSON(201, profile)
}

func (server *ServiceHTTPServer) getCurrentUserHandler(context *gin.Context) {
	id, found := context.Get("user")

	if !found {
		return
	}

	userID, ok := id.(domain.UserID)

	if !ok {
		return
	}

	user, err := server.ucs.GetUserByID(userID)
	if err != nil {
		server.abortWithError(context, *err)
		return
	}

	profile, err := server.ucs.GetProfileByUserID(userID)
	if err != nil {
		server.abortWithError(context, *err)
		return
	}

	context.JSON(200, map[string]interface{}{
		"user":    user,
		"profile": profile,
	})
}
