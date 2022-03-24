package zapkafka

import (
	"github.com/Shopify/sarama"
	"go.uber.org/zap/zapcore"
)

type kafkaWriteSyncer struct {
	topic          string
	producer       sarama.AsyncProducer
	fallbackSyncer zapcore.WriteSyncer
}

func NewKafkaAsyncProducer(addrs []string) (sarama.AsyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Errors = true
	return sarama.NewAsyncProducer(addrs, config)
}

func NewKafkaSyncer(producer sarama.AsyncProducer, topic string, fallbackWs zapcore.WriteSyncer) zapcore.WriteSyncer {
	w := &kafkaWriteSyncer{
		producer:       producer,
		topic:          topic,
		fallbackSyncer: zapcore.AddSync(fallbackWs),
	}

	go func() {
		for e := range producer.Errors() {
			val, err := e.Msg.Value.Encode()
			if err != nil {
				continue
			}

			fallbackWs.Write(val)
		}
	}()
	return w
}

func (ws *kafkaWriteSyncer) Write(b []byte) (n int, err error) {
	b1 := make([]byte, len(b))
	copy(b1, b) // b is reused, we must pass its copy b1 to sarama
	msg := &sarama.ProducerMessage{
		Topic: ws.topic,
		Value: sarama.ByteEncoder(b1),
	}
	ws.producer.Input() <- msg

	/*
		select {
		case ws.producer.Input() <- msg:
		default:
			// if producer block on input channel, write log entry to default fallbackSyncer
			return ws.fallbackSyncer.Write(b1)
		}
	*/

	return len(b1), nil
}

func (ws *kafkaWriteSyncer) Sync() error {
	ws.producer.AsyncClose()
	return ws.fallbackSyncer.Sync()
}
