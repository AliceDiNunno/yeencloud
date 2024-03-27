package domain

type LogLevel string

const (
	LogFieldMail                = "mail"
	LogFieldEnvironmentVariable = "envVar"
	LogFieldAudit               = "audit"
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
