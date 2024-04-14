package audit

import (
	"github.com/AliceDiNunno/yeencloud/internal/core/domain"
)

func (a *Audit) DumpTrace(id domain.AuditTraceID) *domain.AuditTrace {
	trace, exists := a.currentTraces[id]
	if !exists {
		return nil
	}

	return trace
}
