package gin

import (
	"back/src/core/domain"
	"back/src/core/domain/requests"
	"github.com/gin-gonic/gin"
)

func (server *ServiceHTTPServer) getUserMiddleware() gin.HandlerFunc {
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
	}
}

func (server *ServiceHTTPServer) createUserHandler(ctx *gin.Context) {
	var createUserRequest requests.NewUser
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
	id, found := ctx.Get(CtxUserField)

	if !found {
		return
	}

	userID, ok := id.(domain.UserID)

	if !ok {
		return
	}

	profile, err := server.ucs.GetProfileByUserID(server.getTrace(ctx), userID)
	if err != nil {
		server.abortWithError(ctx, *err)
		return
	}

	server.success(ctx, profile)
}
