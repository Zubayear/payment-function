package main

import (
	"Payment/handler"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/cmd"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "payment"
	version = "latest"
)

func main() {
	// Create function
	fnc := micro.NewFunction(
		micro.Name(service),
		micro.Version(version),
	)
	fnc.Init()

	// Handle function
	fnc.Handle(new(handler.Payment))

	cmd.Init()

	if err := broker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	go handler.ConsumeMessage()

	// Run function
	if err := fnc.Run(); err != nil {
		log.Fatal(err)
	}
}
