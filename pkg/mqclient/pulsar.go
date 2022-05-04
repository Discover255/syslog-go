package mqclient

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/discover255/syslog-go/pkg/config"
)

// func init() {
// 	ProducerInit()
// }

var ProducerMap map[string]*pulsar.Producer
var PulsarClientMap map[string]*pulsar.Client

func PulsarClientFactory(hosts []string) *pulsar.Client {
	if PulsarClientMap == nil {
		PulsarClientMap = make(map[string]*pulsar.Client)
	}
	hosts_str := strings.Join(hosts, ",")
	if client, ok := PulsarClientMap[hosts_str]; ok {
		return client
	} else {
		newClient, err := pulsar.NewClient(pulsar.ClientOptions{
			URL:               fmt.Sprintf("pulsar://%s", hosts[0]),
			OperationTimeout:  30 * time.Second,
			ConnectionTimeout: 30 * time.Second,
		})
		if err != nil {
			log.Fatalf("Could not instantiate Pulsar client: %v", err)
		}
		PulsarClientMap[hosts_str] = &newClient
		return &newClient
	}
}

// var PulsarProducerMap map[string]*pulsar.Producer

func PulsarProducerFactory(mqconfig *config.MQConfig) *pulsar.Producer {
	if ProducerMap == nil {
		ProducerMap = make(map[string]*pulsar.Producer)
	}
	client := *PulsarClientFactory(mqconfig.Hosts)
	if producer, ok := ProducerMap[mqconfig.Name]; ok {
		return producer
	} else {
		newProducer, err := client.CreateProducer(pulsar.ProducerOptions{
			Topic: mqconfig.Topic,
		})
		if err != nil {
			log.Fatal(err)
		}
		ProducerMap[mqconfig.Name] = &newProducer
		return &newProducer
	}
}

func ProducerInit() {
	for _, mqconfig := range config.GlobaConfigMap.MQs {
		if mqconfig.Kind == "pulsar" {
			PulsarProducerFactory(&mqconfig)
		}
	}
}

func Close() {
	for _, producer := range ProducerMap {
		(*producer).Close()
	}
	for _, client := range PulsarClientMap {
		(*client).Close()
	}
}

func SendToPulsar(name string, msg string) {
	producer := *ProducerMap[name]
	producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte(msg),
	})
}
