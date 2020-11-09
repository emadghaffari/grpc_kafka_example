package main

import (
	"github.com/emadghaffari/grpc_kafka_example/app"         // start application
	_ "github.com/emadghaffari/grpc_kafka_example/bootstrap" // bootstrap for configs
)

func main() {
	app.StartApplication()
}
