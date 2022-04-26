package main

import (
	"getImageCode/handler"
	pb "getImageCode/proto"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "getimagecode"
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
	pb.RegisterGetImageCodeHandler(srv.Server(), new(handler.GetImageCode))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
