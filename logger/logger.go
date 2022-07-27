package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var L *logrus.Logger

func NewLogger(logpath string) *logrus.Logger {
	file, err := os.OpenFile(logpath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	l := logrus.New()
	l.SetOutput(file)
	l.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	l.SetLevel(logrus.InfoLevel)
	return l
}

func init() {
	L = NewLogger("./log.txt")
}
