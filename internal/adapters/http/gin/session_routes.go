package gin

import (
	"errors"

	"github.com/AliceDiNunno/yeencloud/internal/adapters/audit"
	"github.com/AliceDiNunno/yeencloud/internal/core/domain"
	"github.com/gin-gonic/gin"
)

func (server *ServiceHTTPServer) retrieveSessionMiddleware(ctx *gin.Context) {
	server.auditer.AddStep(server.getTrace(ctx), audit.DefaultSkip)

	token := ctx.GetHeader(HeaderAuthorization)
	if token == "" {
		return
	}

	session, err := server.usecases(ctx).GetSessionByToken(server.getTrace(ctx), token)
	if err != nil {
		server.abortWithError(ctx, errors.Join(&UnauthorizedError{}, err))
		return
	}

	ctx.Set(CtxSessionField, session)
}

func (server *ServiceHTTPServer) requireSessionMiddleware(ctx *gin.Context) {
	server.auditer.AddStep(server.getTrace(ctx), audit.DefaultSkip)

	_, exists := ctx.Get(CtxSessionField)
	if !exists {
		server.abortWithError(ctx, &BadRequestError{})
		return
	}

	_, exists = ctx.Get(CtxUserField)
	if !exists {
		server.abortWithError(ctx, &UnauthorizedError{})
		return
	}
}

func (server *ServiceHTTPServer) createSessionHandler(ctx *gin.Context) {
	var createSessionRequest domain.NewSession

	if err := ctx.ShouldBindJSON(&createSessionRequest); err != nil {
		server.abortWithError(ctx, &BadRequestError{})
		return
	}

	createSessionRequest.Origin = ctx.ClientIP()

	if !server.validate(ctx, createSessionRequest) {
		return
	}

	session, err := server.usecases(ctx).CreateSession(server.getTrace(ctx), createSessionRequest)

	if err != nil {
		server.abortWithError(ctx, errors.Join(&UnauthorizedError{}, err))
		return
	}

	server.created(ctx, session)
}
