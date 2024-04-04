package gin

import (
	"github.com/AliceDiNunno/yeencloud/src/adapters/audit"
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	nice "github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-gonic/gin"
)

func (server *ServiceHTTPServer) setPublicRoutes(r *gin.RouterGroup) {
	r.GET("/status", server.getStatusHandler)

	users := r.Group("/users")
	users.POST("", server.createUserHandler)

	session := r.Group("/session")
	session.POST("", server.createSessionHandler)
}

func (server *ServiceHTTPServer) setAuthenticatedRoutes(authenticated *gin.RouterGroup) {
	authenticated.GET("/me", server.retrieveCurrentUserHandler)

	organizations := authenticated.Group("/organizations")
	server.setOrganizationRoutes(organizations)
}

func (server *ServiceHTTPServer) SetRoutes() {
	r := server.engine

	// Middlewares
	// They are executed in the order they are declared and they will end in the reverse order
	r.Use(server.ginLogger) // Log all requests
	r.Use(nice.Recovery(server.recoverFromPanic))
	r.Use(server.traceHandlerMiddleware)   // Start tracing the request
	r.Use(server.RequestContextMiddleware) // Create a request context
	// r.Use(server.timeoutMiddleware())
	r.Use(server.noResponseHandlerMiddleware) // Ensure that all routes write a response // create a request context
	r.Use(server.retrieveSessionMiddleware)   // Retrieve a session if it exists
	r.Use(server.getUserProfileMiddleware)    // Retrieve the user profile a session exists
	r.Use(server.getLangMiddleware)           // Retrieve the language of the user if it exists or use the one set by the request

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

	if ctx.Writer.Written() {
		return
	}

	server.auditer.Log(server.getTrace(ctx), audit.NoStep).WithLevel(domain.LogLevelError).Msg("No response written")

	server.abortWithError(ctx, ErrorInternal)
}

func (server *ServiceHTTPServer) printRoutes(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	server.log.Log(domain.LogLevelInfo).WithFields(domain.LogFields{
		HttpHandlerField:      handlerName,
		HttpHandlerCountField: nuHandlers,
	}).Msg(httpMethod + " " + absolutePath)
}
