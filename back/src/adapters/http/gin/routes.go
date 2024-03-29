package gin

import (
	"back/src/core/domain"
	"github.com/gin-gonic/gin"
)

func (server *ServiceHTTPServer) SetRoutes() {
	r := server.engine

	//Get prerequisites
	r.Use(server.getSessionMiddleware())
	r.Use(server.getUserMiddleware())
	r.Use(server.getLangMiddleware())

	//Unauthenticated routes
	r.GET("/languages", server.getLanguagesHandler)
	r.GET("/status", server.getStatusHandler)

	users := r.Group("/users")
	users.POST("/", server.createUserHandler)

	session := r.Group("/session")
	session.POST("/", server.createSessionHandler)

	organizations := r.Group("/organizations")
	organizations.GET("/", server.getOrganizationsHandler)

	//Authenticated routes
	authenticated := r.Group("/")
	authenticated.Use(server.requireSessionMiddleware())

	authenticated.GET("/me", server.getCurrentUserHandler)

	server.SetErrors(r)
}

func (server *ServiceHTTPServer) SetErrors(r *gin.Engine) {
	r.NoRoute(func(context *gin.Context) {
		server.abortWithError(context, domain.ErrorNotFound)
	})

	r.NoMethod(func(context *gin.Context) {
		server.abortWithError(context, domain.ErrorNoMethod)
	})
}
