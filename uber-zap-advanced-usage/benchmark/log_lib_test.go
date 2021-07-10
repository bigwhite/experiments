package main

import (
	"io"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func BenchmarkLogrus(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	logger := logrus.New()
	logger.SetOutput(io.Discard)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		logger.WithFields(logrus.Fields{
			"url":     "http://foo.com",
			"attempt": 3,
			"backoff": time.Second,
		}).Info("failed to fetch URL")
	}
}

func BenchmarkZap(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	cfg := zap.NewProductionConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(io.Discard),
		zapcore.InfoLevel,
	)
	logger := zap.New(core)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("failed to fetch URL",
			zap.String("url", `http://foo.com`),
			zap.Int("attempt", 3),
			zap.Duration("backoff", time.Second),
		)
	}
}
