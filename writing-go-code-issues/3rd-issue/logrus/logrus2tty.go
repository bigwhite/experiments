package main

import (
	"github.com/Sirupsen/logrus"
)

func main() {
	customFormatter := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	logger := logrus.New()
	logger.Formatter = customFormatter

	logger.Info("logrus log to tty in normal text formatter")
}
