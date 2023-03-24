package subscribe

import (
	scenarios "bigwhite/autotester/scenarios"
	"testing"
)

func Test_Subscribe_S0001_SubscribeOK(t *testing.T) {
	t.Parallel() // indicate the case can be ran in parallel mode

	tests := []struct {
		name  string
		topic string
		qos   byte
	}{
		{
			name:  "Case_001: Subscribe with QoS 0",
			topic: "a/b/c",
			qos:   0,
		},
		{
			name:  "Case_002: Subscribe with QoS 1",
			topic: "a/b/c",
			qos:   1,
		},
		{
			name:  "Case_003: Subscribe with QoS 2",
			topic: "a/b/c",
			qos:   2,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // indicate the case can be ran in parallel mode
			client, testCaseTeardown, err := scenarios.TestCaseSetup(addr, nil)
			if err != nil {
				t.Errorf("want ok, got %v", err)
				return
			}
			defer testCaseTeardown()

			token := client.Subscribe(tt.topic, tt.qos, nil)
			token.Wait()

			// Check if subscription was successful
			if token.Error() != nil {
				t.Errorf("want ok, got %v", token.Error())
			}

			token = client.Unsubscribe(tt.topic)
			token.Wait()
			if token.Error() != nil {
				t.Errorf("want ok, got %v", token.Error())
			}
		})
	}
}

func Test_Subscribe_S0002_SubscribeFail(t *testing.T) {
}
