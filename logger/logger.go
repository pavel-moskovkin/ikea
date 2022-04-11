package logger

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

const defaultLevel = logrus.DebugLevel

func New(loglevel string) *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	level, err := logrus.ParseLevel(loglevel)
	if err != nil {
		fmt.Printf("unable to parse logLevel=%s, setting to level Debug", loglevel)
		level = defaultLevel
	}
	logger.SetLevel(level)
	return logger
}
