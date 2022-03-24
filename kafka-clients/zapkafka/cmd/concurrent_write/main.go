package main

import (
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"time"

	log "github.com/bigwhite/zapkafka"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SaramaProducer() {
	p, err := log.NewKafkaAsyncProducer([]string{"localhost:29092"})
	if err != nil {
		panic(err)
	}
	logger := log.New(log.NewKafkaSyncer(p, "test", zapcore.AddSync(os.Stderr)), int8(0))
	var wg sync.WaitGroup
	var cnt int64

	for j := 0; j < 10; j++ {
		wg.Add(1)
		go func(j int) {
			var value string
			for i := 0; i < 10000; i++ {
				now := time.Now()
				value = fmt.Sprintf("%02d-%04d-%s", j, i, now.Format("15:04:05"))
				logger.Info("log message:", zap.String("value", value))
				atomic.AddInt64(&cnt, 1)
			}
			wg.Done()
		}(j)
	}

	wg.Wait()
	logger.Sync()
	println("cnt =", atomic.LoadInt64(&cnt))
	time.Sleep(10 * time.Second)
}

func main() {
	SaramaProducer()
}
