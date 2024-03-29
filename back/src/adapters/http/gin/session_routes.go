package gin

import (
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/gin-gonic/gin"
)

func (server *ServiceHTTPServer) getSessionMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		server.auditer.AddStep(server.getTrace(ctx))

		token := ctx.GetHeader(HeaderAuthorization)
		if token == "" {
			return
		}

		session, err := server.ucs.GetSessionByToken(server.getTrace(ctx), token)
		if err != nil {
			server.abortWithError(ctx, *err)
			return
		}

		ctx.Set(CtxSessionField, session)
	}
}

func (server *ServiceHTTPServer) requireSessionMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		server.auditer.AddStep(server.getTrace(ctx))

		_, exists := ctx.Get(CtxSessionField)
		if !exists {
			server.abortWithError(ctx, domain.ErrorAuthenticationTokenMissing)
			return
		}

		_, exists = ctx.Get(CtxUserField)
		if !exists {
			server.abortWithError(ctx, domain.ErrorUserNotFound)
			return
		}
	}
}

func (server *ServiceHTTPServer) createSessionHandler(ctx *gin.Context) {
	var createSessionRequest domain.NewSession
	if err := ctx.ShouldBindJSON(&createSessionRequest); err != nil {
		server.abortWithError(ctx, domain.ErrorBadRequest)
		return
	}

	createSessionRequest.Origin = ctx.ClientIP()

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
