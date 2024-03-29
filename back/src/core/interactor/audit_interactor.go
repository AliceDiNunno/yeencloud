package interactor

import "github.com/AliceDiNunno/yeencloud/src/core/domain"

type Audit interface {
	NewTrace(trigger string, data ...interface{}) domain.AuditID
	AddStep(id domain.AuditID, details ...interface{}) domain.StepID

	EndStep(id domain.AuditID, step domain.StepID)
	EndTrace(id domain.AuditID, result ...interface{}) domain.Request

	Log(id domain.AuditID, step domain.StepID) LogMessage
	DumpTrace(id domain.AuditID) *domain.Request
}
