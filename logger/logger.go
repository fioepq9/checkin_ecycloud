package logger

import "github.com/sirupsen/logrus"

var L *logrus.Logger

func NewLogger() *logrus.Logger {
	l := logrus.New()
	l.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	l.SetLevel(logrus.InfoLevel)
	return l
}

func init() {
	L = NewLogger()
}
