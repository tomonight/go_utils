package log

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

type RollFileLogger struct {
	*logrus.Logger
	file        *os.File
	path        string
	maxFile     int
	maxSize     int
	curSize     int
	buf         *bytes.Buffer
	bufC        chan []byte
	doneC       chan struct{}
	exitC       chan struct{}
	flushC      chan struct{}
	writeStdout bool
}

func NewRollFileLogger(path string, maxSize int, maxFile int, writeStdout bool) *RollFileLogger {
	sep := -1
	for i, c := range path {
		if os.IsPathSeparator(uint8(c)) {
			sep = i
		}
	}

	if sep != -1 {
		os.MkdirAll(path[0:sep], os.ModePerm)
	}

	roll := RollFileLogger{}
	roll.Logger = logrus.New()
	roll.Logger.Out = &roll
	roll.Logger.Formatter = new(SimpleFormatter)

	roll.file = nil
	roll.path = path
	roll.maxSize = maxSize
	roll.maxFile = maxFile
	roll.curSize = 0
	roll.buf = bytes.NewBuffer([]byte{})
	roll.buf.Grow(16384 + 1024)
	roll.bufC = make(chan []byte, 4096)
	roll.doneC = make(chan struct{})
	roll.exitC = make(chan struct{})
	roll.writeStdout = writeStdout
	roll.flushC = make(chan struct{}, 8)

	go serveLogger(&roll)

	return &roll
}

func (roll *RollFileLogger) Write(buff []byte) (int, error) {
	tmp := make([]byte, len(buff))
	copy(tmp, buff)

	select {
	case roll.bufC <- tmp:
	case <-time.After(time.Second):
	}

	return len(buff), nil
}

func (roll *RollFileLogger) writeBuf(buff []byte, force bool) {
	if !roll.rollFile() {
		return
	}

	roll.buf.Write(buff)

	if !force && roll.buf.Len() < 16*1024 {
		return
	} else {
		bts := roll.buf.Bytes()
		roll.file.Write(bts)
		if roll.writeStdout {
			os.Stdout.Write(bts)
		}
		roll.curSize += roll.buf.Len()
		roll.buf.Reset()
	}
}

func (roll *RollFileLogger) rollFile() bool {
	if roll.file == nil {
		f, err := os.OpenFile(roll.genFilePath(0), os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			return false
		}

		st, err := f.Stat()
		if err != nil {
			f.Close()
			return false
		}

		f.Seek(0, io.SeekEnd)
		roll.file = f
		roll.curSize = int(st.Size())
	}

	if roll.curSize >= roll.maxSize {
		roll.file.Close()
		roll.file = nil
		roll.shiftFile()

		f, err := os.OpenFile(roll.genFilePath(0), os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			return false
		}

		roll.file = f
		roll.curSize = 0
	}

	return true
}

func (roll *RollFileLogger) shiftFile() bool {
	os.Remove(roll.genFilePath(roll.maxFile))
	for i := roll.maxFile - 1; i >= 0; i-- {
		os.Rename(roll.genFilePath(i), roll.genFilePath(i+1))
	}

	return true
}

func (roll *RollFileLogger) genFilePath(fileNum int) string {
	if fileNum == 0 {
		return roll.path
	} else {
		return fmt.Sprintf("%s.%d", roll.path, fileNum)
	}
}

func (roll *RollFileLogger) SetLevel(level string) {
	if level == "debug" {
		roll.Logger.SetLevel(logrus.DebugLevel)
	} else if level == "info" {
		roll.Logger.SetLevel(logrus.InfoLevel)
	} else if level == "warn" {
		roll.Logger.SetLevel(logrus.WarnLevel)
	} else if level == "error" {
		roll.Logger.SetLevel(logrus.ErrorLevel)
	} else if level == "fatal" {
		roll.Logger.SetLevel(logrus.FatalLevel)
	} else {
		roll.Logger.SetLevel(logrus.InfoLevel)
	}
}

func serveLogger(roll *RollFileLogger) {
	for {
		select {
		case <-roll.flushC:
			roll.flush()
		case <-roll.doneC:
			roll.flush()
			roll.writeBuf([]byte{}, true)
			close(roll.exitC)
			return
		case buf := <-roll.bufC:
			roll.writeBuf(buf, false)
		case <-time.After(time.Second * 1):
			roll.writeBuf([]byte{}, true)
		}
	}
}

func (roll *RollFileLogger) Done() {
	close(roll.doneC)
	select {
	case <-roll.exitC:
	case <-time.After(time.Second):
	}
}

func (roll *RollFileLogger) flush() {
	for {
		select {
		case buf := <-roll.bufC:
			roll.writeBuf(buf, false)
		default:
			return
		}
	}
}

func (roll *RollFileLogger) Flush() {
	select {
	case roll.flushC <- struct{}{}:
	default:
	}
}
