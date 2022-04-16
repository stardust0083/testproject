package main

import (
	"deleteSession/handler"
	pb "deleteSession/proto"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "deletesession"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
	)
	srv.Init()

	// Register handler
	pb.RegisterDeleteSessionHandler(srv.Server(), new(handler.DeleteSession))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
