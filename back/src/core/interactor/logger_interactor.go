package interactor

import (
	"context"
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
)

type LogMessage interface {
	WithLevel(Level domain.LogLevel) LogMessage

	At(time int64) LogMessage

	WithContext(ctx context.Context) LogMessage

	WithField(key string, value interface{}) LogMessage
	WithFields(fields map[string]interface{}) LogMessage

	Msg(message string) // Commits the log
}

type Logger interface {
	Log(Level domain.LogLevel) LogMessage
}
