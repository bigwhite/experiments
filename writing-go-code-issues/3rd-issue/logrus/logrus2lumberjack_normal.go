package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/natefinch/lumberjack"
)

func main() {
	//    isColored := (f.ForceColors || f.isTerminal) && !f.DisableColors
	customFormatter := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
	}
	logger := logrus.New()
	logger.Formatter = customFormatter

	rotateLogger := &lumberjack.Logger{
		Filename: "./foo.log",
	}
	logger.Out = rotateLogger
	logger.Info("logrus log to lumberjack in normal text formatter")
}
