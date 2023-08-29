package main

import (
	"log/slog"
)

func main() {
	slog.Info("my first slog msg", "greeting", "hello, slog")
}
