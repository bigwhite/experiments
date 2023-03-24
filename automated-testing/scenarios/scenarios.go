package scenarios

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func TestCaseSetup(addr string, opts *mqtt.ClientOptions) (client mqtt.Client, testCaseTeardown func(), err error) {
	if opts == nil {
		opts = mqtt.NewClientOptions()
	}
	opts.AddBroker("tcp://" + addr)

	// Create and start MQTT client
	client = mqtt.NewClient(opts)
	token := client.Connect()
	token.Wait()

	err = token.Error()
	testCaseTeardown = func() {
		client.Disconnect(0)
	}
	return
}
