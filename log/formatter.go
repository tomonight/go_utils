package log

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"sort"
	"sync"
)

const (
	WithFieldSkip = 7
)

type KeyField struct {
	key   string
	field interface{}
}

type SortKeyFields []KeyField

func (s SortKeyFields) Len() int           { return len(s) }
func (s SortKeyFields) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortKeyFields) Less(i, j int) bool { return s[i].key < s[j].key }

type SimpleFormatter struct {
	mu            sync.Mutex
	maxFileLength int64
}

func (f *SimpleFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer

	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	f.appendKeyValue(b, "time", entry.Time.Format("2006-01-02 15:04:05"))
	level := entry.Level.String()
	if level == "warning" {
		level = "warn"
	}
	f.appendKeyValue(b, "level", fmt.Sprintf("%-5s", level))
	if _, ok := entry.Data[fieldFile]; ok {
		f.appendKeyValue(b, fieldFile, entry.Data[fieldFile])
	}

	fields := make(SortKeyFields, 0, len(entry.Data))
	for k, f := range entry.Data {
		if k == fieldFile || k == FiledSkip {
			continue
		}
		fields = append(fields, KeyField{key: k, field: f})
	}

	sort.Sort(fields)
	for i := range fields {
		f.appendKeyValue(b, fields[i].key, fields[i].field)
	}

	if entry.Message != "" {
		f.appendKeyValue(b, "msg", entry.Message)
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}

func (f *SimpleFormatter) appendKeyValue(b *bytes.Buffer, key string, value interface{}) {
	if b.Len() > 0 {
		b.WriteByte(' ')
	}
	b.WriteString(key)
	b.WriteByte('=')
	if key == fieldFile {
		file := fmt.Sprint(value)
		len := f.updateMaxFileLength(int64(len(file)))
		f.appendValue(b, fmt.Sprintf("%-"+fmt.Sprintf("%d", len)+"s", file))
	} else {
		f.appendValue(b, value)
	}
}

func (f *SimpleFormatter) appendValue(b *bytes.Buffer, value interface{}) {
	stringVal, ok := value.(string)
	if !ok {
		stringVal = fmt.Sprint(value)
	}

	b.WriteString(stringVal)
}

func (f *SimpleFormatter) updateMaxFileLength(len int64) int64 {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.maxFileLength < len {
		f.maxFileLength = len
	}
	return f.maxFileLength
}
