package domain

import "time"

type AuditID string
type AuditSaveFunc func([]byte)

func (t AuditID) String() string {
	return string(t)
}

type Request struct {
	ID        AuditID `json:"id"`
	StartedAt int64   `json:"startedAt"`
	EndedAt   int64   `json:"endedAt"`

	Trigger     string        `json:"trigger"`
	TriggerData []interface{} `json:"triggerData"`

	Result []interface{} `json:"result"`

	Content *Step `json:"content"`
}

type StepID string

func (t StepID) String() string {
	return string(t)
}

type Step struct {
	ID      StepID                 `json:"id"`
	Next    *Step                  `json:"next"`
	Caller  map[string]interface{} `json:"caller"`
	Details []interface{}          `json:"details"`

	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}
