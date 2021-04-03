package logging

import (
	"github.com/sirupsen/logrus"
)

func GetLogger() *logrus.Logger {
	logger := logrus.New()

	Formatter := new(logrus.TextFormatter)
	Formatter.FullTimestamp = true
	logger.SetFormatter(Formatter)
	logger.SetLevel(logrus.InfoLevel)

	return logger
}
