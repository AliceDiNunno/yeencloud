package gin

import (
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/gin-gonic/gin"
)

func (server *ServiceHTTPServer) setOrganizationRoutes(organizationGroup *gin.RouterGroup) {
	organizations := organizationGroup.Group("")
	organizations.GET("", server.listOrganizationsHandler)   // List all organizations for the user
	organizations.POST("", server.createOrganizationHandler) // Create a new organization

	specificOrganization := organizations.Group("/:organization", server.retrieveOrganizationMiddleware)
	specificOrganization.GET("", server.getOrganizationHandler)       // Get an organization details
	specificOrganization.PUT("", server.updateOrganizationHandler)    // Update an organization details
	specificOrganization.DELETE("", server.deleteOrganizationHandler) // Delete an organization
}

func (server *ServiceHTTPServer) retrieveOrganizationMiddleware(ctx *gin.Context) {
	server.auditer.AddStep(server.getTrace(ctx))

	profile, found := server.getProfile(ctx)

	if !found {
		return
	}

	organizationID := ctx.Param("organization")

	organization, err := server.ucs.GetOrganizationByID(server.getTrace(ctx), profile.ID, domain.OrganizationID(organizationID))
	if err != nil {
		server.abortWithError(ctx, *err)
		return
	}

	ctx.Set(CtxOrganizationField, organization)
}

func (server *ServiceHTTPServer) getOrganizationFromParam(ctx *gin.Context) (domain.Organization, bool) {
	organizationID, found := ctx.Get(CtxOrganizationField)

	if !found {
		return domain.Organization{}, false
	}

	organization, ok := organizationID.(domain.Organization)

	if !ok {
		return domain.Organization{}, false
	}

	return organization, true
}

func (server *ServiceHTTPServer) listOrganizationsHandler(ctx *gin.Context) {
	profile, found := server.getProfile(ctx)

	if !found {
		return
	}

	trace := server.getTrace(ctx)

	organizations, err := server.ucs.ListOrganizationsByProfile(trace, profile.ID)

	if err != nil {
		server.abortWithError(ctx, *err)
		return
	}

	server.success(ctx, organizations)
}

func (server *ServiceHTTPServer) createOrganizationHandler(ctx *gin.Context) {
	var createOrganizationRequest domain.NewOrganization

	if err := ctx.ShouldBindJSON(&createOrganizationRequest); err != nil {
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
	profile, found := server.getProfile(ctx)
	if !found {
		return
	}

	organization, found := server.getOrganizationFromParam(ctx)
	if !found {
		return
	}

	organization, err := server.ucs.GetOrganizationByID(server.getTrace(ctx), profile.ID, organization.ID)

	if err != nil {
		server.abortWithError(ctx, *err)
		return
	}

	server.success(ctx, organization)
}

func (server *ServiceHTTPServer) updateOrganizationHandler(ctx *gin.Context) {
	var updateOrganizationRequest domain.UpdateOrganization

	profile, found := server.getProfile(ctx)
	if !found {
		return
	}

	organization, found := server.getOrganizationFromParam(ctx)
	if !found {
		return
	}

	if err := ctx.ShouldBindJSON(&updateOrganizationRequest); err != nil {
		server.abortWithError(ctx, ErrorBadRequest)
		return
	}

	if !server.validate(ctx, updateOrganizationRequest) {
		return
	}

	organization, err := server.ucs.UpdateOrganization(server.getTrace(ctx), profile.ID, organization.ID, updateOrganizationRequest)

	if err != nil {
		server.abortWithError(ctx, *err)
		return
	}

	server.success(ctx, organization)
}

func (server *ServiceHTTPServer) deleteOrganizationHandler(ctx *gin.Context) {
	profile, found := server.getProfile(ctx)
	if !found {
		return
	}

	organization, found := server.getOrganizationFromParam(ctx)
	if !found {
		return
	}

	err := server.ucs.DeleteOrganization(server.getTrace(ctx), profile.ID, organization.ID)

	if err != nil {
		server.abortWithError(ctx, *err)
		return
	}

	server.success(ctx, organization)
}
