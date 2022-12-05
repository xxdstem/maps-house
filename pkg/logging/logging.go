package logging

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/siruspen/logrus"
)

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func NewLogger() Logger {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: time.Kitchen,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
	}

	if err := os.Mkdir("logs", 0644); err != nil && !os.IsExist(err) {
		panic(err)
	}
	f, err := os.OpenFile("./logs/log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	if err != nil {
		panic(err)
	}
	l.SetOutput(io.MultiWriter(f, os.Stdout))
	l.SetLevel(logrus.DebugLevel)
	e = logrus.NewEntry(l)
	return Logger{e}
}

func GetLogger() Logger {
	return Logger{e}
}
