package main

import (
	"context"
	"net"
	"os"

	"log/slog"
)

func main() {
	var lvl slog.LevelVar
	lvl.Set(slog.LevelDebug)
	opts := slog.HandlerOptions{
		Level: &lvl,
	}
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, &opts)))

	slog.Info("before resetting log level:")

	slog.Info("greeting", "name", "tony")
	slog.Error("oops", "err", net.ErrClosed, "status", 500)
	slog.LogAttrs(context.Background(), slog.LevelError, "oops",
		slog.Int("status", 500), slog.Any("err", net.ErrClosed))

	slog.Info("after resetting log level to error level:")
	lvl.Set(slog.LevelError)
	slog.Info("greeting", "name", "tony")
	slog.Error("oops", "err", net.ErrClosed, "status", 500)
	slog.LogAttrs(context.Background(), slog.LevelError, "oops",
		slog.Int("status", 500), slog.Any("err", net.ErrClosed))

}
