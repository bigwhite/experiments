package main

import (
	"log/slog"
	"os"
)

func main() {
	f, err := os.Create("foo.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	logger := slog.New(slog.NewJSONHandler(f, nil))
	slog.SetDefault(logger)
	slog.Info("greeting", "say", "hello")
}
