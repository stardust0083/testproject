package main

import (
	"postRet/handler"
	pb "postRet/proto"

	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/asim/go-micro/plugins/transport/grpc/v4"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "go.micro.srv.PostRet"
	version = "latest"
)

func main() {
	// Create service
	reg := consul.NewRegistry()
	ser := grpc.NewTransport()
	srv := micro.NewService(
		micro.Registry(reg),
		micro.Transport(ser),
		micro.Name(service),
		micro.Version(version),
	)
	srv.Init()

	// Register handler
	pb.RegisterPostRetHandler(srv.Server(), new(handler.PostRet))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
