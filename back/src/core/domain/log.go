package domain

type LogLevel string

type LogField string
type LogFields map[LogField]interface{}

const (
	LogFieldConfigVersion    = LogField("config.version")
	LogFieldConfigDatabase   = LogField("config.database")
	LogFieldConfigHTTP       = LogField("config.http")
	LogFieldConfigRunContext = LogField("config.run_context")

	LogFieldProfile     = LogField("profile")
	LogFieldProfileID   = LogFieldProfile + ".id"
	LogFieldProfileMail = LogFieldProfile + ".mail"
	LogFieldProfileName = LogFieldProfile + ".name"

	LogFieldTrace          = LogField("trace")
	LogFieldTraceID        = LogFieldTrace + ".id"
	LogFieldTraceDump      = LogFieldTrace + ".dump"
	LogFieldTraceResult    = LogFieldTrace + ".result"
	LogFieldTraceTrigger   = LogFieldTrace + ".trigger"
	LogFieldTraceStepCount = LogFieldTrace + ".step_count"

	LogFieldTraceStep        = LogFieldTrace + ".%d"
	LogFieldTraceStepCaller  = LogFieldTraceStep + ".caller"
	LogFieldTraceStepDetails = LogFieldTraceStep + ".details"
	LogFieldTraceStepStart   = LogFieldTraceStep + ".start"
	LogFieldTraceStepEnd     = LogFieldTraceStep + ".end"

	LogFieldStep   = LogField("step")
	LogFieldStepID = LogFieldStep + ".id"

	LogFieldError = LogField("error")

	LogFieldTime        = LogField("time")
	LogFieldTimeStarted = LogFieldTime + ".started"
	LogFieldTimeEnded   = LogFieldTime + ".ended"
	LogFieldDuration    = LogFieldTime + ".duration"
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
	return string(l)
}
