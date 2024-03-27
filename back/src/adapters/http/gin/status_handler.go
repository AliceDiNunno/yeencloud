package gin

import (
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/gin-gonic/gin"
)

func (server *ServiceHTTPServer) getStatusHandler(context *gin.Context) {
	auditID := server.getTrace(context)

	status := gin.H{
		"message":   "OK",
		"version":   server.versionConfig,
		"languages": server.ucs.GetAvailableLanguages(),
	}

	stepID := server.auditer.AddStep(auditID, status)
	server.auditer.Log(auditID, stepID).WithLevel(domain.LogLevelInfo).Msg("Status request")

	// TODO: use a struct for the response
	server.success(context, status)

	server.auditer.EndStep(auditID, stepID)
}
