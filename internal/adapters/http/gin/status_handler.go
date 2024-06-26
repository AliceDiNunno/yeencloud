package gin

import (
	"github.com/AliceDiNunno/yeencloud/internal/adapters/audit"
	"github.com/AliceDiNunno/yeencloud/internal/core/domain"
	"github.com/AliceDiNunno/yeencloud/internal/core/domain/config"
	"github.com/gin-gonic/gin"
)

type statusReply struct {
	Message   string               `json:"message"`
	Version   config.VersionConfig `json:"version"`
	Languages []domain.Language    `json:"languages"`
}

func (server *ServiceHTTPServer) getStatusHandler(context *gin.Context) {
	auditID := server.getTrace(context)

	status := statusReply{
		Message:   "OK",
		Version:   server.versionConfig,
		Languages: server.localize.GetAvailableLanguages(),
	}
	stepID := server.auditer.AddStep(auditID, audit.DefaultSkip, status)
	server.auditer.Log(auditID, stepID).WithLevel(domain.LogLevelInfo).Msg("Status request")

	server.auditer.EndStep(auditID, stepID)
	server.success(context, status)
}
