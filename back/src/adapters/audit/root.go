package audit

import (
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
)

func (a *Audit) DumpTrace(id domain.AuditID) *domain.Request {
	trace, exists := a.currentTraces[id]
	if !exists {
		return nil
	}

	return trace
}
