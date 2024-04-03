package gin

import (
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/gin-gonic/gin"
)

func (server *ServiceHTTPServer) setOrganizationMembersRoutes(specificOrganization *gin.RouterGroup) {
	organizationMembers := specificOrganization.Group("/members")
	organizationMembers.GET("", server.getOrganizationMembersHandler)    // List all members of an organization
	organizationMembers.POST("", server.inviteOrganizationMemberHandler) // Invite a new member to an organization

	specificOrganizationMember := organizationMembers.Group("/:member", server.retrieveOrganizationMemberMiddleware)
	specificOrganizationMember.PUT("", server.updateOrganizationMemberHandler)    // Update a member of an organization
	specificOrganizationMember.DELETE("", server.removeOrganizationMemberHandler) // Remove a member from an organization
}

func (server *ServiceHTTPServer) retrieveOrganizationMemberMiddleware(context *gin.Context) {

}

func (server *ServiceHTTPServer) getOrganizationMemberFromParam(ctx *gin.Context) (domain.OrganizationMember, bool) {
	return domain.OrganizationMember{}, false
}

func (server *ServiceHTTPServer) getOrganizationMembersHandler(ctx *gin.Context) {

}

func (server *ServiceHTTPServer) inviteOrganizationMemberHandler(ctx *gin.Context) {

}

func (server *ServiceHTTPServer) updateOrganizationMemberHandler(ctx *gin.Context) {

}

func (server *ServiceHTTPServer) removeOrganizationMemberHandler(ctx *gin.Context) {

}
