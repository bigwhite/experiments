package zapkafka

import (
	"io"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	l     *zap.Logger // zap ensure that zap.Logger is safe for concurrent use
	cfg   zap.Config
	level zap.AtomicLevel
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.l.Info(msg, fields...)
}

func New(writer io.Writer, level int8, opts ...zap.Option) *Logger {
	if writer == nil {
		panic("the writer is nil")
	}
	atomicLevel := zap.NewAtomicLevelAt(zapcore.Level(level))

	logger := &Logger{
		cfg:   zap.NewProductionConfig(),
		level: atomicLevel,
	}

	logger.cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(time.RFC3339)) // 2021-11-19 10:11:30.777
	}
	logger.cfg.EncoderConfig.TimeKey = "logtime"

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(logger.cfg.EncoderConfig),
		zapcore.AddSync(writer),
		atomicLevel,
	)
	logger.l = zap.New(core, opts...)
	return logger
}

// SetLevel alters the logging level on runtime
// it is concurrent-safe
func (l *Logger) SetLevel(level int8) error {
	l.level.SetLevel(zapcore.Level(level))
	return nil
}
