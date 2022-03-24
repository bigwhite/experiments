package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ReadConfig(configFile string) kafka.ConfigMap {

	m := make(map[string]kafka.ConfigValue)

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

func main() {
	conf := ReadConfig("./producer.conf")

	topic := "test"
	p, err := kafka.NewProducer(&conf)
	var mu sync.Mutex

	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}
	var wg sync.WaitGroup
	var cnt int64

	// Go-routine to handle message delivery reports and
	// possibly other event types (errors, stats, etc)
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
						*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
	}()

	for j := 0; j < 10; j++ {
		wg.Add(1)
		go func(j int) {
			var value string
			for i := 0; i < 10000; i++ {
				key := ""
				now := time.Now()
				value = fmt.Sprintf("%02d-%04d-%s", j, i, now.Format("15:04:05"))
				mu.Lock()
				p.Produce(&kafka.Message{
					TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
					Key:            []byte(key),
					Value:          []byte(value),
				}, nil)
				mu.Unlock()
				atomic.AddInt64(&cnt, 1)
			}
			wg.Done()
		}(j)
	}

	wg.Wait()
	// Wait for all messages to be delivered
	time.Sleep(10 * time.Second)
	p.Close()
}
