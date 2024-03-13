package gin

import (
	"back/src/core/domain"
	"back/src/core/domain/requests"
	"github.com/gin-gonic/gin"
)

func (server *ServiceHTTPServer) getSessionMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		server.auditer.AddStep(server.getTrace(ctx))

		token := ctx.GetHeader("Authorization")
		if token == "" {
			return
		}

		session, err := server.ucs.GetSessionByToken(server.getTrace(ctx), token)
		if err != nil {
			server.abortWithError(ctx, *err)
			return
		}

		ctx.Set("session", session)
	}
}

func (server *ServiceHTTPServer) requireSessionMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		server.auditer.AddStep(server.getTrace(ctx))

		_, exists := ctx.Get("session")
		if !exists {
			server.abortWithError(ctx, domain.ErrorAuthenticationTokenMissing)
			return
		}

		_, exists = ctx.Get("user")
	}
}

func (server *ServiceHTTPServer) createSessionHandler(ctx *gin.Context) {
	var createSessionRequest requests.NewSession
	if err := ctx.ShouldBindJSON(&createSessionRequest); err != nil {
		server.abortWithError(ctx, domain.ErrorBadRequest)
		return
	}

	createSessionRequest.IP = ctx.ClientIP()

	if !server.validate(createSessionRequest, ctx) {
		return
	}

	session, err := server.ucs.CreateSession(server.getTrace(ctx), createSessionRequest)

	if err != nil {
		server.abortWithError(ctx, *err)
		return
	}

	server.created(ctx, session)
}
