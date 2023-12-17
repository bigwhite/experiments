package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/testcontainers/testcontainers-go/modules/kafka"
)

func TestKafkaSetup(t *testing.T) {
	ctx := context.Background()

	kafkaContainer, err := kafka.RunContainer(ctx, kafka.WithClusterID("test-cluster"))
	if err != nil {
		panic(err)
	}

	// Clean up the container
	defer func() {
		if err := kafkaContainer.Terminate(ctx); err != nil {
			panic(err)
		}
	}()

	state, err := kafkaContainer.State(ctx)
	if err != nil {
		panic(err)
	}

	if kafkaContainer.ClusterID != "test-cluster" {
		t.Errorf("want test-cluster, actual %s", kafkaContainer.ClusterID)
	}
	if state.Running != true {
		t.Errorf("want true, actual %t", state.Running)
	}
	brokers, _ := kafkaContainer.Brokers(ctx)
	fmt.Printf("%q\n", brokers)
}
