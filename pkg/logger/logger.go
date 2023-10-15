package logger

import (
	f "fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/tel-io/tel/v2"
)

func Debugf(format string, args ...interface{}) {
	tel.Global().Debug(f.Sprintf(format, args...))
}

func Infof(format string, args ...interface{}) {
	tel.Global().Info(f.Sprintf(format, args...))
}

func Errorf(action string, err interface{}) {
	var filename, function string
	pc, file, _, ok := runtime.Caller(1)
	if ok {
		filename = filepath.Base(file)
		toSplit := runtime.FuncForPC(pc).Name()
		splited := strings.Split(toSplit, ".")
		function = splited[len(splited)-1]
	}
	message := f.Sprintf("%s --> func %s --> action %v --> error: %v", filename, function, action, err)
	tel.Global().Error(message)
}

func Fatalf(format string, args ...interface{}) {
	tel.Global().Fatal(f.Sprintf(format, args...))
}
