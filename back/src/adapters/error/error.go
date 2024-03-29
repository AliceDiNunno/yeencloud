package error

import (
	"back/src/core/domain"
	"fmt"
	"hash/crc32"
	"os"
	"runtime"
	"strings"
)

func stackFingerprint(s domain.Stack) string {
	hash := ""
	for _, frame := range s {
		hash = fmt.Sprintf("%s%s%s%d", hash, frame.File, frame.Method, frame.Line)
	}
	return fingerprint(hash)
}

func functionName(pc uintptr) string {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "<unknown>"
	}
	name := fn.Name()
	end := strings.LastIndex(name, string(os.PathSeparator))
	return name[end+1:]
}

func getStack() domain.Stack {
	stack := make(domain.Stack, 0)

	for i := 0; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		functionName := functionName(pc)

		stack = append(stack, domain.Frame{file, functionName, line})
	}

	// TODO: find the best way to display the stack
	//stack = stack[:len(stack)-1]

	stack = stack[2 : len(stack)-1]

	return stack
}

func fingerprint(str string) string {
	hash := crc32.NewIEEE()

	fmt.Fprintf(hash, str)

	return fmt.Sprintf("%x", hash.Sum32())
}

func fingerprintError(trace *domain.Error) string {
	hash := ""
	currentTrace := trace
	for {
		hash = fmt.Sprintf("%s%s", hash, stackFingerprint(currentTrace.Stack))
		currentTrace = currentTrace.Child
		if currentTrace == nil {
			break
		}
	}
	return fingerprint(hash)
}

func Trace(err domain.ErrorDescription) *domain.Error {
	error := domain.Error{
		ErrorDescription: err,
		Stack:            getStack(),
	}

	error.Fingerprint = fingerprintError(&error)

	return &error
}
