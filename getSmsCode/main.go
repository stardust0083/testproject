package main

import (
	"getSmsCode/handler"
	pb "getSmsCode/proto"

	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/asim/go-micro/plugins/transport/grpc/v4"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "go.micro.srv.GetSmsCode"
	version = "latest"
)

func main() {
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
	pb.RegisterGetSmsCodeHandler(srv.Server(), new(handler.GetSmsCode))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
