package errors

import (
	"runtime"
	"strings"
)

type trace struct {
	FileName     string `json:"fileName"`
	Line         int    `json:"line"`
	FunctionName string `json:"functionName"`
}

func Trace() (i trace) {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		i.FileName = file
		i.Line = line
		i.FunctionName = "?"
	}

	i.FileName = file
	i.Line = line
	i.FunctionName = fn.Name()

	return
}

func (i trace) GetFunctionName() (rs string) {
	f := strings.Split(i.FunctionName, "/")
	if len(f) > 0 {
		rs = f[len(f)-1]
	}

	f = strings.Split(rs, ".")
	if len(f) > 0 {
		rs = f[len(f)-1]
	}

	return
}

func (i trace) GetPackagePath() (rs string) {
	f := strings.Split(i.FunctionName, "/")
	if len(f) > 0 {
		rs = f[len(f)-1]
	}

	f = strings.Split(rs, ".")
	if len(f) > 0 {
		rs = f[len(f)-1]
	}

	return
}
