package main

import (
	"net"
	"os"

	"golang.org/x/exp/slog"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr)))
	slog.Info("hello", "name", "Al")
	slog.Error("oops", net.ErrClosed, "status", 500)
	slog.LogAttrs(slog.ErrorLevel, "oops",
		slog.Int("status", 500), slog.Any("err", net.ErrClosed))
}
