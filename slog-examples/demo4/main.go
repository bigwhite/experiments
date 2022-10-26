package main

import (
	"bytes"
	"fmt"
	"net"
	"time"

	"golang.org/x/exp/slog"
)

type ChanHandler struct {
	slog.Handler
	ch  chan []byte
	buf *bytes.Buffer
}

func (h *ChanHandler) Enabled(level slog.Level) bool {
	return h.Handler.Enabled(level)
}

func (h *ChanHandler) Handle(r slog.Record) error {
	err := h.Handler.Handle(r)
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
	var b = make([]byte, 256)
	h := &ChanHandler{
		buf: bytes.NewBuffer(b),
		ch:  ch,
	}

	h.Handler = slog.NewJSONHandler(h.buf)

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
	slog.Error("oops", net.ErrClosed, "status", 500)
	slog.LogAttrs(slog.ErrorLevel, "oops",
		slog.Int("status", 500), slog.Any("err", net.ErrClosed))

	time.Sleep(3 * time.Second)
}
