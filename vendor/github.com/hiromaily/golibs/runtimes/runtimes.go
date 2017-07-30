package runtimes

import (
	"bytes"
	"runtime"
	"strings"
)

// CurrentFunc is to get current func name
func CurrentFunc(skip int) string {
	programCounter, _, _, ok := runtime.Caller(skip)
	if !ok {
		return ""
	}
	sl := strings.Split(runtime.FuncForPC(programCounter).Name(), ".")
	return sl[len(sl)-1]
}

// CurrentFuncV2 is to get current func name
func CurrentFuncV2() []byte {
	b := make([]byte, 250)
	b = b[:runtime.Stack(b, false)]
	for i := 0; i < 3; i++ {
		j := bytes.IndexByte(b, '\n')
		if j < 0 {
			return nil
		}

		b = b[j+1:]
	}
	i := bytes.IndexByte(b, '(')
	if i < 0 {
		return nil
	}

	return b[:i]
}
