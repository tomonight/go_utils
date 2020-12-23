package log

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func startup(console bool) error {
	return Initialize("log.log", 1, 1, "debug", console)
}

func TestLevel(t *testing.T) {
	if err := startup(true); err != nil {
		t.Errorf("init err:%s", err.Error())
		return
	}

	defer Done()

	SetLevel("debug")
	writeLogf("debug")

	SetLevel("info")
	writeLogf("info")

	SetLevel("warn")
	writeLogf("warn")

	SetLevel("error")
	writeLogf("error")
}

func writeLogf(format string, args ...interface{}) {
	Debugf(format, args...)
	Debug(format, args...)

	Infof(format, args...)
	Info(format, args...)

	Warnf(format, args...)
	Warn(format, args...)

	Errorf(format, args...)
	Error(format, args...)
}

func TestWithField(t *testing.T) {
	if err := startup(true); err != nil {
		t.Errorf("init err:%s", err.Error())
		return
	}

	defer Done()
	SetLevel("debug")
	entry := WithField("filed1", "value1")
	entry.Infof("with filed")

	entry = WithFields(logrus.Fields{"filed1": "value1", "filed2": "value2"})
	entry.Warnf("with filed")
	entry.Warn("with field 2")
}

func TestRollFile(t *testing.T) {
	if err := startup(false); err != nil {
		t.Errorf("init err:%s", err.Error())
		return
	}

	defer Done()

	SetLevel("debug")
	for i := 0; i < 1024*128; i++ {
		Debugf("rolling file %d", i+1)
	}
}
