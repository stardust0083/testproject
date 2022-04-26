package main

import (
	"getSession/handler"
	pb "getSession/proto"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "getsession"
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
	pb.RegisterGetSessionHandler(srv.Server(), new(handler.GetSession))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
