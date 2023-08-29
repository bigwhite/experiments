package main

import (
	"net"
	"os"

	"log/slog"
)

func main() {
	h := slog.NewTextHandler(os.Stderr, nil)
	l := slog.New(h)
	l.Info("greeting", "name", "tony")
	l.Error("oops", "err", net.ErrClosed, "status", 500)

	h1 := slog.NewJSONHandler(os.Stderr, nil)
	l1 := slog.New(h1)
	l1.Info("greeting", "name", "tony")
	l1.Error("oops", "err", net.ErrClosed, "status", 500)

	slog.SetDefault(l)
	slog.Info("textHandler after setDefault", "name", "tony", "age", 30)
	slog.SetDefault(l1)
	slog.Info("jsonHandler after setDefault", "name", "tony", "age", 30)
}
