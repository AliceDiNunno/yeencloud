package domain

import (
	"runtime"
)

// MARK: - Objects

type AuditTraceID string
type AuditSaveFunc func([]byte)

type AuditTrace struct {
	ID        AuditTraceID `json:"id"`
	StartedAt int64        `json:"startedAt"`
	EndedAt   int64        `json:"endedAt"`

	Trigger     string            `json:"trigger"`
	TriggerData map[string]string `json:"triggerData"`
	Frame       runtime.Frame     `json:"frame"`

	Result []interface{} `json:"result"`

	Content []AuditTraceStep `json:"content"`
}

// MARK: - Logs
var (
	LogScopeTrace          = LogScope{Identifier: "trace"}
	LogFieldTraceID        = LogField{Scope: LogScopeTrace, Identifier: "id"}
	LogFieldTraceDump      = LogField{Scope: LogScopeTrace, Identifier: "dump"}
	LogFieldTraceResult    = LogField{Scope: LogScopeTrace, Identifier: "result"}
	LogFieldTraceTrigger   = LogField{Scope: LogScopeTrace, Identifier: "trigger"}
	LogFieldTraceStepCount = LogField{Scope: LogScopeTrace, Identifier: "step_count"}

	LogScopeStep   = LogScope{Parent: &LogScopeTrace, Identifier: "step"}
	LogFieldStepID = LogField{Scope: LogScopeStep, Identifier: "id"}

	LogScopeTraceStep        = LogScope{Parent: &LogScopeStep, Identifier: "%d"}
	LogFieldTraceStepCaller  = LogField{Scope: LogScopeTraceStep, Identifier: "caller"}
	LogFieldTraceStepDetails = LogField{Scope: LogScopeTraceStep, Identifier: "details"}
	LogFieldTraceStepStart   = LogField{Scope: LogScopeTraceStep, Identifier: "start"}
	LogFieldTraceStepEnd     = LogField{Scope: LogScopeTraceStep, Identifier: "end"}
)

// MARK: - Functions

func (t AuditTraceID) String() string {
	return string(t)
}
