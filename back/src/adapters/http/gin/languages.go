package gin

import "github.com/gin-gonic/gin"

// #YC-8 TODO: move to status route
func (server *ServiceHTTPServer) getLanguagesHandler(context *gin.Context) {
	languages := server.ucs.GetAvailableLanguages()

	context.JSON(200, languages)
}
