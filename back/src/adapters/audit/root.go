package audit

import "back/src/core/domain"

type root struct {
	ID        domain.AuditID
	StartedAt int64
	EndedAt   int64

	Trigger     string
	TriggerData []interface{}

	Result []interface{}

	Content *step
}
