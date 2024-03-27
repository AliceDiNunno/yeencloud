package gin

import (
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/gin-gonic/gin"
)

func (server *ServiceHTTPServer) getOrganizationsHandler(ctx *gin.Context) {
	id, found := ctx.Get(CtxProfileField)

	if !found {
		return
	}

	profile, ok := id.(domain.Profile)

	if !ok {
		return
	}

	organization, err := server.ucs.GetOrganizationsByProfileID(server.getTrace(ctx), profile.ID)
	if err != nil {
		server.abortWithError(ctx, *err)
		return
	}

	server.success(ctx, organization)
}
