package main

import (
	"getArea/handler"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "getarea"
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
	fnc.Handle(new(handler.GetArea))

	// Run function
	if err := fnc.Run(); err != nil {
		log.Fatal(err)
	}
}
