package main

import (
	"net"
	"os"

	"golang.org/x/exp/slog"
)

func main() {
	var lvl = &slog.AtomicLevel{}
	lvl.Set(slog.DebugLevel)
	opts := slog.HandlerOptions{
		Level: lvl,
	}
	slog.SetDefault(slog.New(opts.NewJSONHandler(os.Stderr)))

	slog.Info("before resetting log level:")

	slog.Info("hello", "name", "Al")
	slog.Error("oops", net.ErrClosed, "status", 500)
	slog.LogAttrs(slog.ErrorLevel, "oops",
		slog.Int("status", 500), slog.Any("err", net.ErrClosed))

	slog.Info("after resetting log level to error level:")
	lvl.Set(slog.ErrorLevel)
	slog.Info("hello", "name", "Al")
	slog.Error("oops", net.ErrClosed, "status", 500)
	slog.LogAttrs(slog.ErrorLevel, "oops",
		slog.Int("status", 500), slog.Any("err", net.ErrClosed))

}
