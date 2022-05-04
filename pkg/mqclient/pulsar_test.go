package mqclient

import (
	"context"
	"testing"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/discover255/syslog-go/pkg/config"
)

// func TestSendToPulsar(t *testing.T) {
// 	SendToPulsar("hello world!")
// }

func TestPulsarProducerFactory(t *testing.T) {
	filename := "./sample.yaml"
	config.ReadConfig(&filename)
	// fmt.Println(config.GlobaConfigMap.MQs[0])
	var producer pulsar.Producer
	for _, mq_config := range config.GlobaConfigMap.MQs {
		producer = *PulsarProducerFactory(&mq_config)
		producer.Send(context.Background(), &pulsar.ProducerMessage{
			Payload: []byte("hello world!"),
		})
	}
	defer producer.Close()
}

func TestInitProducer(t *testing.T) {
	filename := "./sample.yaml"
	config.ReadConfig(&filename)
	ProducerInit()
	defer Close()
	SendToPulsar("default", "hello world!")
}
