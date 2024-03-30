package gin

import (
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
)

func (server *ServiceHTTPServer) getOrganizationsHandler(ctx *gin.Context) {
	profile, found := server.getProfile(ctx)

	if !found {
		return
	}

	organization, err := server.ucs.GetOrganizationsByProfileID(server.getTrace(ctx), profile.ID)
	if err != nil {
		server.abortWithError(ctx, *err)
		return
	}

	server.success(ctx, organization)
}

func (server *ServiceHTTPServer) createOrganizationHandler(ctx *gin.Context) {
	var createOrganizationRequest domain.NewOrganization

	if err := ctx.ShouldBindJSON(&createOrganizationRequest); err != nil {
		spew.Dump(err)
		server.abortWithError(ctx, ErrorBadRequest)
		return
	}

	if !server.validate(ctx, createOrganizationRequest) {
		return
	}

	profile, found := server.getProfile(ctx)
	if !found {
		return
	}

	audit := server.getTrace(ctx)

	session, err := server.ucs.CreateOrganization(audit, profile.ID, createOrganizationRequest)

	if err != nil {
		server.abortWithError(ctx, *err)
		return
	}

	server.created(ctx, session)
}

func (server *ServiceHTTPServer) getOrganizationHandler(ctx *gin.Context) {

}

func (server *ServiceHTTPServer) updateOrganizationHandler(ctx *gin.Context) {

}

func (server *ServiceHTTPServer) deleteOrganizationHandler(ctx *gin.Context) {

}
