package main

import (
	"getArea/handler"
	pb "getArea/proto"

	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/asim/go-micro/plugins/transport/grpc/v4"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	Service = "go.micro.srv.GetArea"
	Version = "latest"
)

func main() {
	reg := consul.NewRegistry()
	ser := grpc.NewTransport()
	microService := micro.NewService(
		micro.Registry(reg),
		micro.Transport(ser),
		micro.Name(Service),
		micro.Version(Version),
	)
	microService.Init()
	err := pb.RegisterGetAreaHandler(microService.Server(), new(handler.GetArea))
	if err != nil {
		log.Fatal(err)
	}
	// Run service
	if err := microService.Run(); err != nil {
		log.Fatal(err)
	}
}
