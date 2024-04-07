package domain

// MARK: - Objects

type LogLevel string

type LogField struct {
	Scope LogScope

	Identifier string
}
type LogFields map[LogField]interface{}

var (
	// MARK: - Config

	LogFieldConfig             = LogScope{Identifier: "config"}
	LogFieldConfigVersion      = LogField{Scope: LogFieldConfig, Identifier: "version"}
	LogFieldConfigDatabase     = LogField{Scope: LogFieldConfig, Identifier: "database"}
	LogFieldConfigHTTP         = LogField{Scope: LogFieldConfig, Identifier: "http"}
	LogFieldConfigRunContext   = LogField{Scope: LogFieldConfig, Identifier: "run_context"}
	LogFieldConfigLocalization = LogField{Scope: LogFieldConfig, Identifier: "localization"}

	// MARK: - Time

	LogFieldTime        = LogScope{Identifier: "time"}
	LogFieldTimeStarted = LogField{Scope: LogFieldTime, Identifier: "started"}
	LogFieldTimeEnded   = LogField{Scope: LogFieldTime, Identifier: "ended"}
	LogFieldDuration    = LogField{Scope: LogFieldTime, Identifier: "duration"}
)

// MARK: - Log Levels

const (
	LogLevelDebug = LogLevel("DEBUG")
	LogLevelInfo  = LogLevel("INFO")
	LogLevelSQL   = LogLevel("SQL")
	LogLevelWarn  = LogLevel("WARN")
	LogLevelError = LogLevel("ERROR")
	LogLevelPanic = LogLevel("PANIC")
	LogLevelFatal = LogLevel("FATAL")
)

// MARK: - Functions

func (l LogLevel) String() string {
	return string(l)
}

func (l LogField) String() string {
	return l.Scope.String() + logSeparator + l.Identifier
}
