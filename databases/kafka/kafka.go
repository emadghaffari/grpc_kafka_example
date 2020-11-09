package kafka

import (
	"github.com/Shopify/sarama"
)

var (
	// Kafka is instance of kafkaInterface
	Kafka kafkaInterface = &STR{}
)

type kafkaInterface interface {
	Validate()
}

// STR struct
type STR struct {
	Brokers  []string
	Version  string
	Group    string
	Topics   []string
	Assignor string
	Oldest   *bool
	Verbose  *bool
	Config   *sarama.Config
}

// Set func
func Set(ks *STR) {
	ks.Validate()
	Kafka = ks
}
