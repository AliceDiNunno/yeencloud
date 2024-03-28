package gin

import (
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/gin-gonic/gin"
)

func (server *ServiceHTTPServer) getUserProfileMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		server.auditer.AddStep(server.getTrace(ctx))

		session, exists := ctx.Get(CtxSessionField)

		if !exists {
			return
		}

		// This is necessary to make sure the user is still valid
		user, err := server.ucs.GetUserByID(server.getTrace(ctx), session.(domain.Session).UserID)

		if err != nil {
			server.abortWithError(ctx, *err)
			return
		}

		ctx.Set(CtxUserField, user.ID)

		profile, err := server.ucs.GetProfileByUserID(server.getTrace(ctx), user.ID)

		if err != nil {
			server.abortWithError(ctx, *err)
			return
		}

		ctx.Set("mail", user.Email)
		ctx.Set(CtxProfileField, profile)
	}
}

func (server *ServiceHTTPServer) createUserHandler(ctx *gin.Context) {
	var createUserRequest domain.NewUser
	if err := ctx.ShouldBindJSON(&createUserRequest); err != nil {
		server.abortWithError(ctx, domain.ErrorBadRequest)
		return
	}

	if !server.validate(createUserRequest, ctx) {
		return
	}

	language := ctx.GetString(CtxLanguageField)

	profile, err := server.ucs.CreateUser(server.getTrace(ctx), createUserRequest, language)

	if err != nil {
		server.abortWithError(ctx, *err)
		return
	}

	server.created(ctx, profile)
}

func (server *ServiceHTTPServer) getCurrentUserHandler(ctx *gin.Context) {
	id, found := ctx.Get(CtxProfileField)

	if !found {
		return
	}

	profile, ok := id.(domain.Profile)

	if !ok {
		server.abortWithError(ctx, domain.ErrorProfileNotFound)
		return
	}

	server.success(ctx, profile)
}
