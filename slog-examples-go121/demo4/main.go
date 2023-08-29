package main

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"time"

	"log/slog"
)

type ChanHandler struct {
	slog.Handler
	ch  chan []byte
	buf *bytes.Buffer
}

func (h *ChanHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.Handler.Enabled(ctx, level)
}

func (h *ChanHandler) Handle(ctx context.Context, r slog.Record) error {
	err := h.Handler.Handle(ctx, r)
	if err != nil {
		return err
	}
	var nb = make([]byte, h.buf.Len())
	copy(nb, h.buf.Bytes())
	h.ch <- nb
	h.buf.Reset()
	return nil
}

func (h *ChanHandler) WithAttrs(as []slog.Attr) slog.Handler {
	return &ChanHandler{
		buf:     h.buf,
		ch:      h.ch,
		Handler: h.Handler.WithAttrs(as),
	}
}

func (h *ChanHandler) WithGroup(name string) slog.Handler {
	return &ChanHandler{
		buf:     h.buf,
		ch:      h.ch,
		Handler: h.Handler.WithGroup(name),
	}
}

func NewChanHandler(ch chan []byte) *ChanHandler {
	h := &ChanHandler{
		buf: bytes.NewBuffer(nil),
		ch:  ch,
	}

	h.Handler = slog.NewJSONHandler(h.buf, nil)

	return h
}

func main() {
	var ch = make(chan []byte, 100)
	attrs := []slog.Attr{
		{Key: "field1", Value: slog.StringValue("value1")},
		{Key: "field2", Value: slog.StringValue("value2")},
	}
	slog.SetDefault(slog.New(NewChanHandler(ch).WithAttrs(attrs)))
	go func() {
		for {
			b := <-ch
			fmt.Println(string(b))
		}
	}()

	slog.Info("hello", "name", "Al")
	slog.Error("oops", "err", net.ErrClosed, "status", 500)
	slog.LogAttrs(context.Background(), slog.LevelError, "oops",
		slog.Int("status", 500), slog.Any("err", net.ErrClosed))

	time.Sleep(3 * time.Second)
}
