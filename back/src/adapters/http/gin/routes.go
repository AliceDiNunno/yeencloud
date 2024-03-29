package gin

import (
	"back/src/core/domain"
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

	organizations := authenticated.Group("/organizations")
	organizations.GET("/", server.getOrganizationsHandler)
}

func (server *ServiceHTTPServer) SetRoutes() {
	r := server.engine

	// Get prerequisites
	r.Use(server.traceHandlerMiddleware())
	r.Use(server.getSessionMiddleware())
	r.Use(server.getUserMiddleware())
	r.Use(server.getLangMiddleware())

	// Unauthenticated routes
	server.setPublicRoutes(r.Group("/"))

	// Authenticated routes
	authenticated := r.Group("/")
	authenticated.Use(server.requireSessionMiddleware())
	server.setAuthenticatedRoutes(authenticated)

	server.SetErrors(r)
}

func (server *ServiceHTTPServer) SetErrors(r *gin.Engine) {
	r.NoRoute(func(ctx *gin.Context) {
		server.abortWithError(ctx, domain.ErrorNotFound)
	})

	r.NoMethod(func(ctx *gin.Context) {
		server.abortWithError(ctx, domain.ErrorNoMethod)
	})
}
