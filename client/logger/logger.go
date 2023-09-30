package logger

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/srnewbie/sample-server/config"
)

func New(config *config.LoggerConfig) *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	level, _ := logrus.ParseLevel(config.Level)
	logger.SetLevel(level)
	logger.SetOutput(os.Stdout)
	return logger
}
