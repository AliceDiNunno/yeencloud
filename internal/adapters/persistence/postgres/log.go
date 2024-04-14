package postgres

import (
	"context"
	"time"

	"github.com/AliceDiNunno/yeencloud/internal/core/domain"
	"github.com/AliceDiNunno/yeencloud/internal/core/interactor"
	"gorm.io/gorm/logger"
)

var (
	LogFieldSQL = domain.LogScope{Identifier: "sql"}

	LogFieldSQLRowsAffected = domain.LogField{Scope: LogFieldSQL, Identifier: "rows_affected"}
)

type gormLogger struct {
	logger interactor.Logger
}

func (g gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return g
}

func (g gormLogger) log(level domain.LogLevel, s string, i ...interface{}) {
	if g.logger == nil {
		return
	}

	fields := domain.LogFields{}
	for _, v := range i {
		logfields, valid := v.(domain.LogFields)
		if valid {
			for key, value := range logfields {
				fields[key] = value
			}
		}
	}
	g.logger.Log(level).WithFields(fields).Msg(s)
}

func (g gormLogger) Info(ctx context.Context, s string, i ...interface{}) {
	if g.logger == nil {
		return
	}

	g.log(domain.LogLevelInfo, s, i...)
}

func (g gormLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	if g.logger == nil {
		return
	}

	g.log(domain.LogLevelWarn, s, i...)
}

func (g gormLogger) Error(ctx context.Context, s string, i ...interface{}) {
	if g.logger == nil {
		return
	}

	g.log(domain.LogLevelError, s, i...)
}

func (g gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, rows := fc()

	end := time.Now()

	duration := time.Duration(end.UnixMilli() - begin.UnixMilli())

	if g.logger == nil {
		return
	}

	g.log(domain.LogLevelSQL, sql, domain.LogFields{
		domain.LogFieldTimeStarted: begin,
		domain.LogFieldTimeEnded:   end,
		// TODO: move duration to logger when timestarted and timeended are present
		domain.LogFieldDuration: duration,
		domain.LogFieldError:    err,
		LogFieldSQLRowsAffected: rows,
	})
}

func newGormLogger(logger interactor.Logger) *gormLogger {
	return &gormLogger{logger: logger}
}
