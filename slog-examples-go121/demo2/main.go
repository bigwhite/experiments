package main

import (
	"context"
	"net"
	"os"

	"log/slog"
)

func main() {
	opts := slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelError,
	}

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, &opts)))
	slog.Info("open file for reading", "name", "foo.txt", "path", "/home/tonybai/demo/foo.txt")
	slog.Error("open file error", "err", os.ErrNotExist, "status", 2)

	//slog.LogAttrs(context.Background(), slog.LevelError, "oops",
	//slog.Int("status", 500), slog.Any("err", net.ErrClosed))
	l := slog.Default().With("attr1", "attr1_value", "attr2", "attr2_value")
	l.Error("connect server error", "err", net.ErrClosed, "status", 500)
	l.Error("close conn error", "err", net.ErrClosed, "status", 501)
	l.LogAttrs(context.Background(), slog.LevelError, "log with attribute once", slog.String("attr3", "attr3_value"))
	l.Error("reconnect error", "err", net.ErrClosed, "status", 502)

	gl := l.WithGroup("response")
	gl.Error("http post response", "code", 403, "status", "server not response", "server", "10.10.121.88")
}
