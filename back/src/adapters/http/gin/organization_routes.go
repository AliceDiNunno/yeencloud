package gin

import (
	"back/src/core/domain"
	"github.com/gin-gonic/gin"
)

func (server *ServiceHTTPServer) getOrganizationsHandler(ctx *gin.Context) {
	id, found := ctx.Get("user")

	if !found {
		return
	}

	userID, ok := id.(domain.UserID)

	if !ok {
		return
	}

	organization, err := server.ucs.GetOrganizationsByUserID(server.getTrace(ctx), userID)
	if err != nil {
		server.abortWithError(ctx, *err)
		return
	}

	server.success(ctx, organization)
}
