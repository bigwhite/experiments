package connection

import (
	"testing"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Test_Connection_S0001_ConnectOKWithoutAuth(t *testing.T) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://" + addr)

	// Create and start MQTT client
	client := mqtt.NewClient(opts)
	token := client.Connect()
	token.Wait()

	// Check if connection was successful
	if token.Error() != nil {
		t.Errorf("want ok, got %v", token.Error())
		return
	}
	client.Disconnect(0)
}

func Test_Connection_S0002_ConnectOKWithAuth(t *testing.T) {
	//... ...
}
