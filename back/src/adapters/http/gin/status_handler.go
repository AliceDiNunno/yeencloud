package gin

import "github.com/gin-gonic/gin"

func (server *ServiceHTTPServer) getStatusHandler(context *gin.Context) {
	server.success(context, gin.H{
		"message":   "OK",
		"version":   server.versionConfig,
		"languages": server.ucs.GetAvailableLanguages(),
	})
}
