package gin

import (
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/gin-gonic/gin"
)

func (server *ServiceHTTPServer) setPublicRoutes(r *gin.RouterGroup) {
	r.GET("/status", server.getStatusHandler)

	users := r.Group("/users")
	users.POST("/", server.createUserHandler)

	session := r.Group("/session")
	session.POST("/", server.createSessionHandler)
}

func (server *ServiceHTTPServer) setAuthenticatedRoutes(authenticated *gin.RouterGroup) {
	authenticated.GET("/me", server.getCurrentUserHandler)

	// Organizations
	organizations := authenticated.Group("/organizations")
	organizations.GET("/", server.getOrganizationsHandler)
	organizations.POST("/", server.createOrganizationHandler)

	organization := authenticated.Group("/:organization")
	organization.GET("/", server.getOrganizationHandler)
	organization.PATCH("/", server.updateOrganizationHandler)
	organization.DELETE("/", server.deleteOrganizationHandler)
}

func (server *ServiceHTTPServer) SetRoutes() {
	r := server.engine

	// Get prerequisites
	r.Use(server.ginLogger)
	r.Use(gin.Recovery())
	r.Use(server.traceHandlerMiddleware)
	r.Use(server.noResponseHandlerMiddleware)
	r.Use(server.getSessionMiddleware)
	r.Use(server.getUserProfileMiddleware)
	r.Use(server.getLangMiddleware)

	// Unauthenticated routes
	server.setPublicRoutes(r.Group("/"))

	// Authenticated routes
	authenticated := r.Group("/")
	authenticated.Use(server.requireSessionMiddleware)
	server.setAuthenticatedRoutes(authenticated)

	server.SetErrors(r)
}

func (server *ServiceHTTPServer) SetErrors(r *gin.Engine) {
	r.NoRoute(func(ctx *gin.Context) {
		server.abortWithError(ctx, ErrorPageNotFound)
	})

	r.NoMethod(func(ctx *gin.Context) {
		server.abortWithError(ctx, ErrorMethodNotAllowed)
	})
}

func (server *ServiceHTTPServer) noResponseHandlerMiddleware(ctx *gin.Context) {
	// This should never happen, if a router does not write a response, we're returning an internal error
	ctx.Next()

	// TODO: add error to audit

	if ctx.Writer.Written() {
		return
	}

	server.abortWithError(ctx, ErrorInternal)
}

func (server *ServiceHTTPServer) printRoutes(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	server.log.Log(domain.LogLevelInfo).WithFields(domain.LogFields{
		HttpHandlerField:      handlerName,
		HttpHandlerCountField: nuHandlers,
	}).Msg(httpMethod + " " + absolutePath)
}
