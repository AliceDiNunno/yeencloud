package gin

import "github.com/gin-gonic/gin"

// TODO: move to status route
func (server *ServiceHTTPServer) getLanguagesHandler(context *gin.Context) {
	languages := server.ucs.GetAvailableLanguages()

	context.JSON(200, languages)
}
