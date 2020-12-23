package log

import (
	"fmt"
	"io"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

var lg *RollFileLogger

const (
	fieldFile = "file"
	Skip      = 3
	FiledSkip = "__skip"
)

func Initialize(file string, maxSize int, backlog int, level string, console bool) error {
	lg = NewRollFileLogger(file, maxSize*1024*1024, backlog, console)
	lg.SetLevel(level)
	return nil
}

func Debugf(format string, args ...interface{}) {
	printf(logrus.DebugLevel, format, args...)
}

func Debug(format string, args ...interface{}) {
	printf(logrus.DebugLevel, format, args...)
}

func Infof(format string, args ...interface{}) {
	printf(logrus.InfoLevel, format, args...)
}

func Info(format string, args ...interface{}) {
	printf(logrus.InfoLevel, format, args...)
}

func Warnf(format string, args ...interface{}) {
	printf(logrus.WarnLevel, format, args...)
}

func Warn(format string, args ...interface{}) {
	printf(logrus.WarnLevel, format, args...)
}

func Warning(format string, args ...interface{}) {
	printf(logrus.WarnLevel, format, args...)
}

func Errorf(format string, args ...interface{}) {
	printf(logrus.ErrorLevel, format, args...)
}

func Error(format string, args ...interface{}) {
	printf(logrus.ErrorLevel, format, args...)
}

func SetLevel(level string) {
	lg.SetLevel(level)
}

func Done() {
	lg.Done()
}

func Flush() {
	lg.Flush()
}

func WithField(key string, value interface{}) *Entry {
	return &Entry{Entry: logrus.NewEntry(lg.Logger).WithField(key, value)}
}

func WithFields(fields logrus.Fields) *Entry {
	return &Entry{Entry: logrus.NewEntry(lg.Logger).WithFields(fields)}
}

func fileAndLine(skip int) string {
	_, file, line, _ := runtime.Caller(skip)
	return fmt.Sprintf("%s:%-3d", path.Base(file), line)
}

func printf(level logrus.Level, format string, args ...interface{}) {
	entry := lg.WithField(fieldFile, fileAndLine(Skip))
	switch level {
	case logrus.DebugLevel:
		entry.Debugf(format, args...)
	case logrus.InfoLevel:
		entry.Infof(format, args...)
	case logrus.WarnLevel:
		entry.Warnf(format, args...)
	case logrus.ErrorLevel:
		entry.Errorf(format, args...)
	}
}

func RawWriter() io.Writer {
	return lg
}
