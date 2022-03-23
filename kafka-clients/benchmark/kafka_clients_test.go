package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/Shopify/sarama"
	ckafkago "github.com/confluentinc/confluent-kafka-go/kafka"
	kafkago "github.com/segmentio/kafka-go"
)

var m = []byte("this is benchmark for three mainstream kafka client")

func BenchmarkSaramaAsync(b *testing.B) {
	b.ReportAllocs()
	config := sarama.NewConfig()
	producer, err := sarama.NewAsyncProducer([]string{"localhost:29092"}, config)
	if err != nil {
		panic(err)
	}

	message := &sarama.ProducerMessage{Topic: "test", Value: sarama.ByteEncoder(m)}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		producer.Input() <- message
	}
}

func BenchmarkSaramaAsyncInParalell(b *testing.B) {
	b.ReportAllocs()
	config := sarama.NewConfig()
	producer, err := sarama.NewAsyncProducer([]string{"localhost:29092"}, config)
	if err != nil {
		panic(err)
	}

	message := &sarama.ProducerMessage{Topic: "test", Value: sarama.ByteEncoder(m)}

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			producer.Input() <- message
		}
	})
}

func BenchmarkKafkaGoAsync(b *testing.B) {
	b.ReportAllocs()
	w := &kafkago.Writer{
		Addr:     kafkago.TCP("localhost:29092"),
		Topic:    "test",
		Balancer: &kafkago.LeastBytes{},
		Async:    true,
	}

	c := context.Background()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w.WriteMessages(c, kafkago.Message{Value: []byte(m)})
	}
}

func BenchmarkKafkaGoAsyncInParalell(b *testing.B) {
	b.ReportAllocs()
	w := &kafkago.Writer{
		Addr:     kafkago.TCP("localhost:29092"),
		Topic:    "test",
		Balancer: &kafkago.LeastBytes{},
		Async:    true,
	}

	c := context.Background()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			w.WriteMessages(c, kafkago.Message{Value: []byte(m)})
		}
	})
}

func ReadConfig(configFile string) ckafkago.ConfigMap {
	m := make(map[string]ckafkago.ConfigValue)

	file, err := os.Open(configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open file: %s", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, "#") && len(line) != 0 {
			kv := strings.Split(line, "=")
			parameter := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])
			m[parameter] = value
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Failed to read file: %s", err)
		os.Exit(1)
	}

	return m

}

func BenchmarkConfluentKafkaGoAsync(b *testing.B) {
	b.ReportAllocs()
	conf := ReadConfig("./confluent-kafka-go.conf")

	topic := "test"
	p, _ := ckafkago.NewProducer(&conf)

	go func() {
		for _ = range p.Events() {
		}
	}()

	key := []byte("")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Produce(&ckafkago.Message{
			TopicPartition: ckafkago.TopicPartition{Topic: &topic, Partition: ckafkago.PartitionAny},
			Key:            key,
			Value:          m,
		}, nil)
	}
}

func BenchmarkConfluentKafkaGoAsyncInParalell(b *testing.B) {
	b.ReportAllocs()
	conf := ReadConfig("./confluent-kafka-go.conf")

	topic := "test"
	p, _ := ckafkago.NewProducer(&conf)

	go func() {
		for range p.Events() {
		}
	}()

	var mu sync.Mutex
	key := []byte("")
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu.Lock()
			p.Produce(&ckafkago.Message{
				TopicPartition: ckafkago.TopicPartition{Topic: &topic, Partition: ckafkago.PartitionAny},
				Key:            key,
				Value:          m,
			}, nil)
			mu.Unlock()
		}
	})
}
