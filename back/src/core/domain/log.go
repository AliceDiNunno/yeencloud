package domain

type LogLevel string

type LogField struct {
	Parent *LogField
	Name   string
}
type LogFields map[LogField]interface{}

var (
	LogFieldConfig           = LogField{Name: "config"}
	LogFieldConfigVersion    = LogField{Parent: &LogFieldConfig, Name: "version"}
	LogFieldConfigDatabase   = LogField{Parent: &LogFieldConfig, Name: "database"}
	LogFieldConfigHTTP       = LogField{Parent: &LogFieldConfig, Name: "http"}
	LogFieldConfigRunContext = LogField{Parent: &LogFieldConfig, Name: "run_context"}

	LogFieldProfile     = LogField{Name: "profile"}
	LogFieldProfileID   = LogField{Parent: &LogFieldProfile, Name: "id"}
	LogFieldProfileMail = LogField{Parent: &LogFieldProfile, Name: "mail"}
	LogFieldProfileName = LogField{Parent: &LogFieldProfile, Name: "name"}

	LogFieldUser   = LogField{Name: "user"}
	LogFieldUserID = LogField{Parent: &LogFieldUser, Name: "id"}

	LogFieldSession = LogField{Name: "session"}

	LogFieldSessionRequest     = LogField{Parent: &LogFieldSession, Name: "request"}
	LogFieldSessionRequestMail = LogField{Parent: &LogFieldSessionRequest, Name: "mail"}

	LogFieldTrace          = LogField{Name: "trace"}
	LogFieldTraceID        = LogField{Parent: &LogFieldTrace, Name: "id"}
	LogFieldTraceDump      = LogField{Parent: &LogFieldTrace, Name: "dump"}
	LogFieldTraceResult    = LogField{Parent: &LogFieldTrace, Name: "result"}
	LogFieldTraceTrigger   = LogField{Parent: &LogFieldTrace, Name: "trigger"}
	LogFieldTraceStepCount = LogField{Parent: &LogFieldTrace, Name: "step_count"}

	LogFieldTraceStep        = LogField{Parent: &LogFieldTrace, Name: "%d"}
	LogFieldTraceStepCaller  = LogField{Parent: &LogFieldTraceStep, Name: "caller"}
	LogFieldTraceStepDetails = LogField{Parent: &LogFieldTraceStep, Name: "details"}
	LogFieldTraceStepStart   = LogField{Parent: &LogFieldTraceStep, Name: "start"}
	LogFieldTraceStepEnd     = LogField{Parent: &LogFieldTraceStep, Name: "end"}

	LogFieldStep   = LogField{Name: "step"}
	LogFieldStepID = LogField{Parent: &LogFieldStep, Name: "id"}

	LogFieldError = LogField{Name: "error"}

	LogFieldTime        = LogField{Name: "time"}
	LogFieldTimeStarted = LogField{Parent: &LogFieldTime, Name: "started"}
	LogFieldTimeEnded   = LogField{Parent: &LogFieldTime, Name: "ended"}
	LogFieldDuration    = LogField{Parent: &LogFieldTime, Name: "duration"}
)

const (
	LogLevelDebug = LogLevel("DEBUG")
	LogLevelInfo  = LogLevel("INFO")
	LogLevelSQL   = LogLevel("SQL")
	LogLevelWarn  = LogLevel("WARN")
	LogLevelError = LogLevel("ERROR")
	LogLevelPanic = LogLevel("PANIC")
	LogLevelFatal = LogLevel("FATAL")
)

func (l LogLevel) String() string {
	return string(l)
}

func (l LogField) String() string {
	if l.Parent == nil {
		return l.Name
	}
	return l.Parent.String() + "." + l.Name
}
