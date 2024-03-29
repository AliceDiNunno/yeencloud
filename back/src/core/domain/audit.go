package domain

import (
	"runtime"
	"time"
)

type AuditID string
type AuditSaveFunc func([]byte)

func (t AuditID) String() string {
	return string(t)
}

type Request struct {
	ID        AuditID `json:"id"`
	StartedAt int64   `json:"startedAt"`
	EndedAt   int64   `json:"endedAt"`

	Trigger     string            `json:"trigger"`
	TriggerData map[string]string `json:"triggerData"`
	Frame       runtime.Frame     `json:"frame"`

	Result []interface{} `json:"result"`

	Content []Step `json:"content"`
}

type StepID string

func (t StepID) String() string {
	return string(t)
}

type Step struct {
	ID      StepID        `json:"id"`
	Caller  runtime.Frame `json:"caller"`
	Details []interface{} `json:"details"`

	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}
