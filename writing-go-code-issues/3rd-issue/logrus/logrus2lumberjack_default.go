package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/natefinch/lumberjack"
)

func main() {
	logger := logrus.New()
	rotateLogger := &lumberjack.Logger{
		Filename: "./foo.log",
	}
	logger.Out = rotateLogger
	logger.Info("logrus log to lumberjack in default text formatter")
}
