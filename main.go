package main

import (
	// start application
	"fmt"

	"github.com/emadghaffari/grpc_kafka_example/bootstrap" // bootstrap for configs
	"github.com/emadghaffari/grpc_kafka_example/databases/kafka"
)

func main() {
	bootstrap.Init()
	// j := kafka.Kafka{}
	fmt.Println(kafka.Kafka)
	kafka.Kafka.Validate()
}
