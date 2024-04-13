package audit

import (
	"runtime"
	"strings"
)

func (a *Audit) getFrame(skip int) runtime.Frame {
	pc, file, line, ok := runtime.Caller(skip)

	functionName := "<unknown>"
	fn := runtime.FuncForPC(pc)
	if fn != nil {
		fname := strings.Split(fn.Name(), "/")
		functionName = fname[len(fname)-1]
	}

	if ok {
		return runtime.Frame{
			PC:       pc,
			Function: functionName,
			File:     file,
			Line:     line,
		}
	}

	return runtime.Frame{}
}
