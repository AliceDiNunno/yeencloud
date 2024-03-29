package domain

type LogLevel string

type LogField string
type LogFields map[LogField]interface{}

const (
	LogFieldMail                = LogField("mail")
	LogFieldEnvironmentVariable = LogField("envVar")
	LogFieldAudit               = LogField("audit")

	LogFieldTraceTrigger = LogField("trace.trigger")
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
