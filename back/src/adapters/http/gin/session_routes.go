package gin

import (
	"back/src/core/domain"
	"back/src/core/domain/requests"
	"github.com/gin-gonic/gin"
)

func (server *ServiceHTTPServer) getSessionMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		if token == "" {
			return
		}

		session, err := server.ucs.GetSessionByToken(token)
		if err != nil {
			server.abortWithError(context, *err)
			return
		}

		context.Set("session", session)
	}
}

func (server *ServiceHTTPServer) createSessionHandler(context *gin.Context) {
	var createSessionRequest requests.NewSession
	if err := context.ShouldBindJSON(&createSessionRequest); err != nil {
		server.abortWithError(context, domain.ErrorBadRequest)
		return
	}

	createSessionRequest.IP = context.ClientIP()

	if !server.validate(createSessionRequest, context) {
		return
	}

	session, err := server.ucs.CreateSession(createSessionRequest)

	if err != nil {
		server.abortWithError(context, *err)
		return
	}

	context.JSON(201, session)
}

func (server *ServiceHTTPServer) requireSessionMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		_, exists := context.Get("session")
		if !exists {
			server.abortWithError(context, domain.ErrorAuthenticationTokenMissing)
			return
		}

		_, exists = context.Get("user")
	}
}
