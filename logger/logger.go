package logger

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var L Logger

func init() {
	L = *NewLogger("./log.txt")
}

type Logger struct {
	InfoFile *os.File
}

func (l *Logger) Info(s string) {
	str := fmt.Sprintf("%s ---- %s\n", time.Now().Format("2006/01/02-15:04:05"), s)
	l.InfoFile.WriteString(str)
}

func (l *Logger) Error(s string, err error) {
	str := fmt.Sprintf("%s ---- %s\n\tError->\n\t\t%s\n\t--->End",
		time.Now().Format("2006/01/02-15:04:05"),
		s,
		strings.ReplaceAll(err.Error(), "\n", "\n\t\t"))
	l.InfoFile.WriteString(str)
}

func NewLogger(logpath string) *Logger {
	file, err := os.OpenFile(logpath, os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	return &Logger{
		InfoFile: file,
	}
}
