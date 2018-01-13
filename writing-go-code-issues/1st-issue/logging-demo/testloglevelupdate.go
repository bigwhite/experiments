package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
)

func watchAndUpdateLoglevel(c chan os.Signal, logger *log.Logger) {
	for {
		select {
		case sig := <-c:
			if sig == syscall.SIGUSR1 {
				level := logger.Level
				if level == log.PanicLevel {
					fmt.Println("Raise log level: It has been already the most top log level: panic level")
				} else {
					logger.SetLevel(level - 1)
					fmt.Println("Raise log level: the current level is", logger.Level)
				}

			} else if sig == syscall.SIGUSR2 {
				level := logger.Level
				if level == log.DebugLevel {
					fmt.Println("Reduce log level: It has been already the lowest log level: debug level")
				} else {
					logger.SetLevel(level + 1)
					fmt.Println("Reduce log level: the current level is", logger.Level)
				}

			} else {
				fmt.Println("receive unknown signal:", sig)
			}
		}
	}
}

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

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1, syscall.SIGUSR2)
	go watchAndUpdateLoglevel(c, logger)

	for {
		logger.Debug("it is debug level log")
		logger.Info("it is info level log")
		logger.Warn("it is warning level log")
		logger.Error("it is warning level log")
		//logger.Fatal("it is fatal level log")
		//logger.Panic("it is panic level log")
		time.Sleep(5 * time.Second)
	}
}
