package postgres

import (
	"context"
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/AliceDiNunno/yeencloud/src/core/interactor"
	"gorm.io/gorm/logger"
	"time"
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
		for key, value := range v.(domain.LogFields) {
			fields[key] = value
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
		"begin":        begin,
		"end":          end,
		"rowsAffected": rows,
		"error":        err,
		"duration":     duration,
	})
}

func newGormLogger(logger interactor.Logger) *gormLogger {
	return &gormLogger{logger: logger}
}
