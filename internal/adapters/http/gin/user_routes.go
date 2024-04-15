package gin

import (
	"github.com/AliceDiNunno/yeencloud/internal/core/domain"
	"github.com/gin-gonic/gin"
)

func (server *ServiceHTTPServer) createUserHandler(ctx *gin.Context) {
	var createUserRequest domain.NewUser

	if err := ctx.ShouldBindJSON(&createUserRequest); err != nil {
		server.abortWithError(ctx, &BadRequestError{})
		return
	}

	if !server.validate(ctx, createUserRequest) {
		return
	}

	language := ctx.GetString(CtxLanguageField)

	profile, err := server.usecases(ctx).CreateUser(server.getTrace(ctx), createUserRequest, language)

	if err != nil {
		server.abortWithError(ctx, err)
		return
	}

	server.created(ctx, profile)
}

func (server *ServiceHTTPServer) retrieveCurrentUserHandler(ctx *gin.Context) {
	id, found := ctx.Get(CtxProfileField)

	if !found {
		return
	}

	profile, ok := id.(domain.Profile)

	if !ok {
		server.abortWithError(ctx, &BadRequestError{})
		return
	}

	server.success(ctx, profile)
}
