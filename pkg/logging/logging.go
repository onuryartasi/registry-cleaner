package logging

import (
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	Formatter := new(logrus.TextFormatter)
	Formatter.FullTimestamp = true
	logger.SetFormatter(Formatter)
	logger.SetLevel(logrus.InfoLevel)
}

func GetLogger() *logrus.Logger {
	return logger
}

func SetLogger(l *logrus.Logger) {
	logger = l
}
