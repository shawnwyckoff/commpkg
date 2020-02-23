package logger

import (
	"strings"
	"time"
)

type Logger interface {
	Debgf(format string, a ...interface{})
	Infof(format string, a ...interface{})
	Warnf(format string, a ...interface{})
	Errof(format string, a ...interface{})
	Fataf(format string, a ...interface{})

	Erro(err error, wrapMsg ...string)
	Fata(err error, wrapMsg ...string)
	AssertOk(err error, wrapMsg ...string)
	AssertTrue(express bool, wrapMsg ...string)
}

type (
	Level string

	Log struct {
		Time     time.Time
		Level    Level
		Text     string
		Reverse1 string            `json:"Reverse1,omitempty" bson:"Reverse1,omitempty"`
		Reverse2 string            `json:"Reverse2,omitempty" bson:"Reverse2,omitempty"`
		Reverse3 string            `json:"Reverse3,omitempty" bson:"Reverse3,omitempty"`
		Reverse4 string            `json:"Reverse4,omitempty" bson:"Reverse4,omitempty"`
		Reverse5 string            `json:"Reverse5,omitempty" bson:"Reverse5,omitempty"`
		ExtTags  map[string]string `json:"ExtTags,omitempty" bson:"ExtTags,omitempty"`
	}
)

const (
	LevelDebg    Level = "DEBG"
	LevelInfo    Level = "INFO"
	LevelWarn    Level = "WARN"
	LevelErro    Level = "ERRO"
	LevelFata    Level = "FATA"
	LevelSuccess Level = "OK"
)

func NewLog(level Level, time time.Time, text string) Log {
	return Log{Level: level, Time: time, Text: text}
}

func (l *Log) SetExtTag(key, value string) *Log {
	if l.ExtTags == nil {
		l.ExtTags = make(map[string]string)
	}
	l.ExtTags[key] = value
	return l
}

func (l *Log) SetReverse1(value string) *Log {
	l.Reverse1 = value
	return l
}

func (l *Log) SetReverse2(value string) *Log {
	l.Reverse2 = value
	return l
}

func (l *Log) SetReverse3(value string) *Log {
	l.Reverse3 = value
	return l
}

func (l *Log) SetReverse4(value string) *Log {
	l.Reverse4 = value
	return l
}

func (l *Log) SetReverse5(value string) *Log {
	l.Reverse5 = value
	return l
}

func (l *Log) GetExtTagEx(key string) (string, bool) {
	if l.ExtTags == nil {
		return "", false
	}
	value, ok := l.ExtTags[key]
	if !ok {
		return "", false
	}
	return value, true
}

func (l *Log) GetExtTag(key string) string {
	value, ok := l.GetExtTagEx(key)
	if !ok {
		return ""
	}
	return value
}

// 清理日志中不需要出现的信息
func (l *Log) clear() {
	l.Text = strings.Trim(l.Text, "\n")
	lines := strings.Split(l.Text, "\n")
	var result []string
	removeTags := []string{
		"runtime/asm_amd64",
		"runtime.goexit",
		"support/xerror",
		"runtime.main",
	}
	for _, ln := range lines {
		jump := false
		for _, tag := range removeTags {
			if strings.Contains(ln, tag) {
				jump = true
				break
			}
		}
		if jump {
			continue
		}
		/*if strings.Contains(ln, "runtime.main") {
			break
		}*/
		ln = strings.Replace(ln, "/usr/local/go/", "$ROOT/", -1)
		ln = strings.Replace(ln, "github.com", "gitlab.com", -1)
		ln = strings.Replace(ln, "bitbucket.org", "gitlab.com", -1)
		ln = strings.Replace(ln, "golang.org", "", -1)
		ln = strings.Replace(ln, "gopkg.in", "", -1)
		ln = strings.Replace(ln, "v2ray.com", "", -1)
		ln = strings.Replace(ln, "go", "c", -1)
		ln = strings.Replace(ln, "taci", "rafw", -1)
		ln = strings.Replace(ln, "shawnwyckoff", "rafael", -1)
		ln = strings.Replace(ln, "kcp", "qtp", -1)
		result = append(result, ln)
	}

	l.Text = strings.Join(result, "\n")
}
