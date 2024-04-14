package gin

import (
	"github.com/AliceDiNunno/yeencloud/internal/adapters/audit"
	"github.com/AliceDiNunno/yeencloud/internal/core/domain"
	"github.com/gin-gonic/gin"
)

func (server *ServiceHTTPServer) validate(c *gin.Context, obj interface{}) bool {
	server.auditer.AddStep(server.getTrace(c), audit.DefaultSkip)

	succeeded, validationErrors := server.validator.Validate(obj)

	if succeeded {
		return true
	}
	lang := c.GetString(CtxLanguageField)

	translatedErrors := map[domain.ValidationFieldName][]string{}

	for name, validationError := range validationErrors {
		fields := []string{}

		for _, err := range validationError {
			fields = append(fields, server.localize.GetLocalizedText(lang, err))
		}

		translatedErrors[name] = fields
	}

	server.abortWithError(c, ErrorBadRequest, translatedErrors)

	return false
}
