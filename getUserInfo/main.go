package main

import (
	"getUserInfo/handler"
	pb "getUserInfo/proto"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"

	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/asim/go-micro/plugins/transport/grpc/v4"
)

var (
	service = "go.micro.srv.GetUserInfo"
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
	pb.RegisterGetUserInfoHandler(srv.Server(), new(handler.GetUserInfo))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
