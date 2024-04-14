package interactor

import (
	"context"
	"github.com/AliceDiNunno/yeencloud/internal/core/domain"
)

type LogMessage interface {
	WithLevel(Level domain.LogLevel) LogMessage

	At(time int64) LogMessage

	WithContext(ctx context.Context) LogMessage

	WithField(key domain.LogField, value interface{}) LogMessage
	WithFields(fields domain.LogFields) LogMessage

	Msg(message string) // Commits the log
}

type Logger interface {
	Log(Level domain.LogLevel) LogMessage
}
