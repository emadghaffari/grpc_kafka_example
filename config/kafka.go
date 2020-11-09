package config

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	"github.com/emadghaffari/grpc_kafka_example/databases/kafka"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	// KafkaConfig for set configs
	KafkaConfig configInterface = &kafkaconfig{}
)

// configInterface
type configInterface interface {
	Set()
}

// config struct
type kafkaconfig struct{}

func (c *kafkaconfig) Set() {
	brokers := viper.GetStringSlice("kafka.brokers")
	version := viper.GetString("kafka.version")
	group := viper.GetString("kafka.group")
	topics := viper.GetStringSlice("kafka.topics")
	assignor := viper.GetString("kafka.assignor")
	oldest := viper.GetBool("kafka.oldest")
	verbose := viper.GetBool("kafka.verbose")

	// configs
	config := sarama.NewConfig()
	config.ClientID = "go-kafka"

	// config.Consumer
	config.Consumer.Return.Errors = true
	config.Consumer.MaxProcessingTime = time.Second
	switch assignor {
	case "sticky":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	case "roundrobin":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	case "range":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	default:
		log.WithFields(log.Fields{
			"error": fmt.Sprintf("Unrecognized consumer group partition assignor: %s", assignor),
		}).Fatal(fmt.Sprintf("Unrecognized consumer group partition assignor: %s", assignor))
	}

	// config.Net
	config.Net.MaxOpenRequests = 1

	// config.Producer
	config.Producer.Idempotent = true
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Compression = sarama.CompressionSnappy
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Retry.Backoff = time.Duration(time.Second * 5)
	config.Producer.Retry.Max = 5
	config.Producer.Compression = sarama.CompressionLZ4
	config.Producer.Timeout = time.Duration(time.Second * 50)

	kafka.Set(&kafka.STR{
		Assignor: assignor,
		Brokers:  brokers,
		Group:    group,
		Config:   config,
		Oldest:   &oldest,
		Verbose:  &verbose,
		Version:  version,
		Topics:   topics,
	})
}
