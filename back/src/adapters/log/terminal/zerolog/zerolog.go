package zerolog

import (
	"github.com/AliceDiNunno/yeencloud/src/adapters/log"
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	"os"
	"regexp"
	"strconv"
)

type ZeroLogMiddleware struct {
}

func (z *ZeroLogMiddleware) colorRestMethodPath(str string) string {
	regex, err := regexp.Compile(`^(?P<method>\w+)\s(?P<path>\/\S+)$`)
	if err != nil {
		return ""
	}

	if regex.MatchString(str) {
		method := regex.FindStringSubmatch(str)[1]
		path := regex.FindStringSubmatch(str)[2]
		color := "\x1b[0m"

		switch method {
		case "GET":
			color = "\x1b[32m"
		case "POST":
			color = "\x1b[34m"
		case "PUT":
			color = "\x1b[36m"
		case "DELETE":
			color = "\x1b[31m"
		case "PATCH":
			color = "\x1b[33m"
		default:
			color = "\x1b[0m"
		}

		return color + method + "\u001B[0m " + path
	}

	return str
}

func (z *ZeroLogMiddleware) colorRestMethodStatus(str string) string {
	regex, err := regexp.Compile(`^(?P<method>\w+)\s(?P<path>\/\S+)\s(?P<status>\d+)$`)
	if err != nil {
		return ""
	}

	if regex.MatchString(str) {
		method := regex.FindStringSubmatch(str)[1]
		path := regex.FindStringSubmatch(str)[2]
		status := regex.FindStringSubmatch(str)[3]

		statusCode, _ := strconv.Atoi(status)

		color := "\x1b[0m"

		if statusCode >= 200 && statusCode < 300 {
			color = "\x1b[32m"
		} else if statusCode >= 300 && statusCode < 400 {
			color = "\x1b[33m"
		} else if statusCode >= 400 && statusCode < 500 {
			color = "\x1b[31m"
		} else if statusCode >= 500 {
			color = "\x1b[95m"
		}

		return color + method + "\u001B[0m " + path + " " + color + status + "\u001B[0m"
	}

	return str
}

func (z *ZeroLogMiddleware) colorize(str string) string {
	str = z.colorRestMethodPath(str)
	str = z.colorRestMethodStatus(str)

	return str
}

func (z *ZeroLogMiddleware) Log(message log.Message) {
	level := zerolog.NoLevel

	switch message.Level {
	case domain.LogLevelDebug:
		level = zerolog.DebugLevel
	case domain.LogLevelInfo:
		level = zerolog.InfoLevel
	case domain.LogLevelWarn:
		level = zerolog.WarnLevel
	case domain.LogLevelError:
		level = zerolog.ErrorLevel
	case domain.LogLevelFatal:
		level = zerolog.FatalLevel
	case domain.LogLevelPanic:
		level = zerolog.PanicLevel
	case domain.LogLevelSQL:
		level = zerolog.TraceLevel
	}

	currentLog := zlog.WithLevel(level)

	for k, v := range message.Fields {
		currentLog = currentLog.Any(k, v)
	}

	msgStr := message.Message

	if message.Level == domain.LogLevelSQL {
		msgStr = "\x1b[35m" + msgStr + "\x1b[0m"
	}
	currentLog.Msg(z.colorize(msgStr))
}

func NewZeroLogMiddleware() *ZeroLogMiddleware {
	zlog.Logger = zlog.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	return &ZeroLogMiddleware{}
}
