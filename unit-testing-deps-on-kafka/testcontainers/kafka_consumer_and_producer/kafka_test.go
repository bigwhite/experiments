package main

import (
	"bytes"
	"context"
	"errors"
	"net"
	"strconv"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go/modules/kafka"

	kc "github.com/segmentio/kafka-go" // kafka client
)

func createTopics(brokers []string, topics ...string) error {
	// to create topics when auto.create.topics.enable='false'
	conn, err := kc.Dial("tcp", brokers[0])
	if err != nil {
		return err
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		return err
	}
	var controllerConn *kc.Conn
	controllerConn, err = kc.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		return err
	}
	defer controllerConn.Close()

	var topicConfigs []kc.TopicConfig
	for _, topic := range topics {
		topicConfig := kc.TopicConfig{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		}
		topicConfigs = append(topicConfigs, topicConfig)
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		return err
	}

	return nil
}

func newWriter(brokers []string, topic string) *kc.Writer {
	return &kc.Writer{
		Addr:                   kc.TCP(brokers...),
		Topic:                  topic,
		Balancer:               &kc.LeastBytes{},
		AllowAutoTopicCreation: true,
		RequiredAcks:           0,
	}
}

func newReader(brokers []string, topic string) *kc.Reader {
	return kc.NewReader(kc.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  "test-group",
		MaxBytes: 10e6, // 10MB
	})
}

func TestProducerAndConsumer(t *testing.T) {
	ctx := context.Background()

	kafkaContainer, err := kafka.RunContainer(ctx, kafka.WithClusterID("test-cluster"))
	if err != nil {
		t.Fatalf("want nil, actual %v\n", err)
	}

	// Clean up the container
	defer func() {
		if err := kafkaContainer.Terminate(ctx); err != nil {
			t.Fatalf("want nil, actual %v\n", err)
		}
	}()

	state, err := kafkaContainer.State(ctx)
	if err != nil {
		t.Fatalf("want nil, actual %v\n", err)
	}

	if state.Running != true {
		t.Errorf("want true, actual %t", state.Running)
	}

	brokers, err := kafkaContainer.Brokers(ctx)
	if err != nil {
		t.Fatalf("want nil, actual %v\n", err)
	}

	topic := "test-topic"
	w := newWriter(brokers, topic)
	defer w.Close()
	r := newReader(brokers, topic)
	defer r.Close()

	err = createTopics(brokers, topic)
	if err != nil {
		t.Fatalf("want nil, actual %v\n", err)
	}
	time.Sleep(5 * time.Second)

	messages := []kc.Message{
		{
			Key:   []byte("Key-A"),
			Value: []byte("Value-A"),
		},
		{
			Key:   []byte("Key-B"),
			Value: []byte("Value-B"),
		},
		{
			Key:   []byte("Key-C"),
			Value: []byte("Value-C"),
		},
		{
			Key:   []byte("Key-D"),
			Value: []byte("Value-D!"),
		},
	}

	const retries = 3
	for i := 0; i < retries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// attempt to create topic prior to publishing the message
		err = w.WriteMessages(ctx, messages...)
		if errors.Is(err, kc.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
			time.Sleep(time.Millisecond * 250)
			continue
		}

		if err != nil {
			t.Fatalf("want nil, actual %v\n", err)
		}
		break
	}

	var getMessages []kc.Message
	for i := 0; i < len(messages); i++ {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			t.Fatalf("want nil, actual %v\n", err)
		}
		getMessages = append(getMessages, m)
	}

	for i := 0; i < len(messages); i++ {
		if !bytes.Equal(getMessages[i].Key, messages[i].Key) {
			t.Errorf("want %s, actual %s\n", string(messages[i].Key), string(getMessages[i].Key))
		}
		if !bytes.Equal(getMessages[i].Value, messages[i].Value) {
			t.Errorf("want %s, actual %s\n", string(messages[i].Value), string(getMessages[i].Value))
		}
	}
}
