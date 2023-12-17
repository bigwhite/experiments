package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"strconv"
	"testing"
	"time"

	kc "github.com/segmentio/kafka-go" // kafka client
)

func createTopics(brokers []string, tlsConfig *tls.Config, topics ...string) error {
	dialer := &kc.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
		TLS:       tlsConfig,
	}

	conn, err := dialer.DialContext(context.Background(), "tcp", brokers[0])
	if err != nil {
		fmt.Println("creating topic: dialer dial error:", err)
		return err
	}
	defer conn.Close()
	fmt.Println("creating topic: dialer dial ok")

	controller, err := conn.Controller()
	if err != nil {
		fmt.Println("creating topic: get controller error:", err)
		return err
	}
	fmt.Println("creating topic: get controller ok")

	var controllerConn *kc.Conn
	controllerConn, err = kc.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		fmt.Println("creating topic: dial control listener error:", err)
		return err
	}
	fmt.Println("creating topic: dial control listener ok")
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

func newWriter(brokers []string, tlsConfig *tls.Config, topic string) *kc.Writer {
	w := &kc.Writer{
		Addr:                   kc.TCP(brokers...),
		Topic:                  topic,
		Balancer:               &kc.LeastBytes{},
		AllowAutoTopicCreation: true,
		Async:                  true,
		//RequiredAcks:           0,
		Completion: func(messages []kc.Message, err error) {
			for _, message := range messages {
				if err != nil {
					fmt.Println("write message fail", err)
				} else {
					fmt.Println("write message ok", string(message.Topic), string(message.Value))
				}
			}
		},
	}

	if tlsConfig != nil {
		w.Transport = &kc.Transport{
			TLS: tlsConfig,
		}
	}
	return w
}

func newReader(brokers []string, tlsConfig *tls.Config, topic string) *kc.Reader {
	dialer := &kc.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
		TLS:       tlsConfig,
	}

	return kc.NewReader(kc.ReaderConfig{
		Dialer:   dialer,
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  "test-group",
		MaxBytes: 10e6, // 10MB
	})
}

func TestProducerAndConsumer(t *testing.T) {
	var err error
	_, err = UpDefaultKakfa()
	if err != nil {
		t.Fatalf("want nil, actual %v\n", err)
	}

	t.Cleanup(func() {
		DownDefaultKakfa()
	})

	brokers := []string{"localhost:9093"}
	topic := "test-topic"

	tlsConfig, _ := newTLSConfig()
	w := newWriter(brokers, tlsConfig, topic)
	defer w.Close()
	r := newReader(brokers, tlsConfig, topic)
	defer r.Close()
	err = createTopics(brokers, tlsConfig, topic)
	if err != nil {
		fmt.Printf("create topic error: %v, but it may not affect the later action, just ignore it\n", err)
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
		if err != nil {
			fmt.Printf("write message error: %v\n", err)
			time.Sleep(time.Second * 2)
			continue
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

func newTLSConfig() (*tls.Config, error) {
	/*
	   // 加载 CA 证书
	   caCert, err := ioutil.ReadFile("/path/to/ca.crt")
	   if err != nil {
	           return nil, err
	   }

	   // 加载客户端证书和私钥
	   cert, err := tls.LoadX509KeyPair("/path/to/client.crt", "/path/to/client.key")
	   if err != nil {
	           return nil, err
	   }

	   // 创建 CertPool 并添加 CA 证书
	   caCertPool := x509.NewCertPool()
	   caCertPool.AppendCertsFromPEM(caCert)
	*/
	// 创建并返回 TLS 配置
	return &tls.Config{
		//RootCAs:      caCertPool,
		//Certificates: []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}, nil
}
