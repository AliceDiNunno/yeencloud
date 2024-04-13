package gin

import (
	"github.com/AliceDiNunno/yeencloud/src/adapters/audit"
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/gin-gonic/gin"
)

func (server *ServiceHTTPServer) getUserProfileMiddleware(ctx *gin.Context) {
	server.auditer.AddStep(server.getTrace(ctx), audit.DefaultSkip)

	session, exists := ctx.Get(CtxSessionField)

	if !exists {
		return
	}

	// This is necessary to make sure the user is still valid
	user, err := server.usecases(ctx).GetUserByID(server.getTrace(ctx), session.(domain.Session).UserID)

	if err != nil {
		server.abortWithError(ctx, *err)
		return
	}

	ctx.Set(CtxUserField, user.ID)

	profile, err := server.usecases(ctx).GetProfileByUserID(server.getTrace(ctx), user.ID)

	if err != nil {
		server.abortWithError(ctx, *err)
		return
	}

	ctx.Set(CtxProfileMailField, user.Email)
	ctx.Set(CtxProfileField, profile)
}

func (server *ServiceHTTPServer) getProfile(ctx *gin.Context) (domain.Profile, bool) {
	field, found := ctx.Get(CtxProfileField)

	if !found {
		return domain.Profile{}, false
	}

	profile, ok := field.(domain.Profile)

	if !ok {
		return domain.Profile{}, false
	}

	return profile, true
}
