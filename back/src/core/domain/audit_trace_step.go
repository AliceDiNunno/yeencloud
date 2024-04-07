package domain

import (
	"runtime"
	"time"
)

// MARK: - Objects

type AuditTraceStepID string

type AuditTraceStep struct {
	ID      AuditTraceStepID `json:"id"`
	Caller  runtime.Frame    `json:"caller"`
	Details []interface{}    `json:"details"`

	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

// MARK: - Functions

func (t AuditTraceStepID) String() string {
	return string(t)
}
