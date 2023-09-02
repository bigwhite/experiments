package main

import (
	"log/slog"

	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	r := &lumberjack.Logger{
		Filename:   "./foo.log",
		LocalTime:  true,
		MaxSize:    1,
		MaxAge:     3,
		MaxBackups: 5,
		Compress:   true,
	}
	logger := slog.New(slog.NewJSONHandler(r, nil))
	slog.SetDefault(logger)

	for i := 0; i < 100000; i++ {
		slog.Info("greeting", "say", "hello")
	}
}
