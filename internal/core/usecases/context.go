package usecases

import (
	"github.com/AliceDiNunno/yeencloud/internal/adapters/audit"
	"github.com/AliceDiNunno/yeencloud/internal/core/domain"
	"github.com/AliceDiNunno/yeencloud/internal/core/interactor"
)

func (self UCs) log(rc *domain.RequestContext, level domain.LogLevel) interactor.LogMessage {
	return self.i.Trace.Log(rc.TraceID, rc.StepID).WithLevel(level)
}

func (self UCs) traceRequest(rc *domain.RequestContext, req func()) {
	auditStepID := self.i.Trace.AddStep(rc.TraceID, audit.DefaultSkip+1)
	rc.StepID = auditStepID
	req()
	self.i.Trace.EndStep(rc.TraceID, auditStepID)
}

func (self UCs) requirePermission(rc *domain.RequestContext, req func(), permission domain.Permission) {
	err := self.checkPermissions(rc.TraceID, rc.Profile.ID, nil, permission)
	if err != nil {
		self.log(rc, domain.LogLevelError).WithField(domain.LogFieldError, err).Msg("Permission denied")
		rc.Done(nil, err)
		return
	}
	req()
}
