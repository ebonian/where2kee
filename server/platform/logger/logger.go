package logger

import (
	"github.com/sirupsen/logrus"
)

func GetLogger() *logrus.Logger {
	logger := logrus.New()
	logger.Level = logrus.InfoLevel
	return logger
}
