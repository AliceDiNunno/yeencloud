package audit

import "back/src/core/domain"

type Request struct {
	ID        domain.AuditID `json:"id"`
	StartedAt int64          `json:"startedAt"`
	EndedAt   int64          `json:"endedAt"`

	Trigger     string        `json:"trigger"`
	TriggerData []interface{} `json:"triggerData"`

	Result []interface{} `json:"result"`

	Content *Step `json:"content"`
}
