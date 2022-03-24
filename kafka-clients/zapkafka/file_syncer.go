package zapkafka

import (
	"io"

	"go.uber.org/zap/zapcore"
)

func NewFileSyncer(writer io.Writer) zapcore.WriteSyncer {
	if ws, ok := writer.(zapcore.WriteSyncer); ok {
		return ws
	}
	return zapcore.Lock(zapcore.AddSync(writer))
}
