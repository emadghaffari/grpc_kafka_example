package bootstrap

import (
	"github.com/emadghaffari/grpc_kafka_example/config"
)

// Init bootstrap the project before start application
func Init() {
	config.ViperConfig.Set()
	config.LogrusConfig.Set()
	config.KafkaConfig.Set()
}
