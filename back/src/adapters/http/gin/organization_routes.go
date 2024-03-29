package gin

import (
	"back/src/core/domain"
	"github.com/gin-gonic/gin"
)

func (server *ServiceHTTPServer) getOrganizationsHandler(context *gin.Context) {
	id, found := context.Get("user")

	if !found {
		return
	}

	userID, ok := id.(domain.UserID)

	if !ok {
		return
	}

	organization, err := server.ucs.GetOrganizationsByUserID(userID)
	if err != nil {
		server.abortWithError(context, *err)
		return
	}

	context.JSON(200, organization)
}
