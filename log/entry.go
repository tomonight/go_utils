package log

import (
	"github.com/sirupsen/logrus"
)

type Entry struct {
	*logrus.Entry
}

func (entry *Entry) Debugf(format string, args ...interface{}) {
	entry.printf(logrus.DebugLevel, format, args...)
}

func (entry *Entry) Debug(format string, args ...interface{}) {
	entry.printf(logrus.DebugLevel, format, args...)
}

func (entry *Entry) Infof(format string, args ...interface{}) {
	entry.printf(logrus.InfoLevel, format, args...)
}

func (entry *Entry) Info(format string, args ...interface{}) {
	entry.printf(logrus.InfoLevel, format, args...)
}

func (entry *Entry) Warnf(format string, args ...interface{}) {
	entry.printf(logrus.WarnLevel, format, args...)
}

func (entry *Entry) Warn(format string, args ...interface{}) {
	entry.printf(logrus.WarnLevel, format, args...)
}

func (entry *Entry) Warning(format string, args ...interface{}) {
	entry.printf(logrus.WarnLevel, format, args...)
}

func (entry *Entry) Errorf(format string, args ...interface{}) {
	entry.printf(logrus.ErrorLevel, format, args...)
}

func (entry *Entry) Error(format string, args ...interface{}) {
	entry.printf(logrus.ErrorLevel, format, args...)
}

func (entry *Entry) WithField(key string, value interface{}) *Entry {
	return &Entry{Entry: entry.Entry.WithField(key, value)}
}

func (entry *Entry) WithFields(fields logrus.Fields) *Entry {
	return &Entry{Entry: entry.Entry.WithFields(fields)}
}

func (entry *Entry) printf(level logrus.Level, format string, args ...interface{}) {
	var skip int
	if v, ok := entry.Data[FiledSkip]; ok {
		skip, _ = v.(int)
	}
	e := entry.WithField(fieldFile, fileAndLine(Skip+skip))
	switch level {
	case logrus.DebugLevel:
		e.Entry.Debugf(format, args...)
	case logrus.InfoLevel:
		e.Entry.Infof(format, args...)
	case logrus.WarnLevel:
		e.Entry.Warnf(format, args...)
	case logrus.ErrorLevel:
		e.Entry.Errorf(format, args...)
	}
}
