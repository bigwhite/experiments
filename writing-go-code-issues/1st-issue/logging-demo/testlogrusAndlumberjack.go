package main

import (
	"time"

	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
)

func main() {
	logger := log.New()
	logger.SetLevel(log.DebugLevel)
	logger.Formatter = &log.JSONFormatter{}

	logger.Out = &lumberjack.Logger{
		Filename:   "./app.log",
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     1,    //days
		Compress:   true, // disabled by default
		LocalTime:  true,
	}

	for {
		logger.Debug("this is an app log")
		time.Sleep(2 * time.Millisecond)
	}
}
