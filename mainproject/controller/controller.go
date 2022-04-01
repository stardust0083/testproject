package controller

import (
	"context"
	"fmt"
	"net/http"

	getArea "getArea"

	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/asim/go-micro/plugins/transport/grpc/v4"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func GetArea(ctx *gin.Context) {
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:8500"}
	})
	ser := grpc.NewTransport()
	micService := micro.NewService(
		micro.Registry(reg),
		micro.Transport(ser),
	)

	microClient := getArea.NewGetAreaService("go.micro.srv.getArea", micService.Client())
	//调用远程服务
	resp, err := microClient.MicroGetArea(context.TODO(), &getArea.Request{})
	if err != nil {
		fmt.Println(err)
		/*ctx.JSON(http.StatusOK,resp)
		return */
	}

	//把int 的0值  json的特性,如果字段是零值,不对这个字段做序列化

	ctx.JSON(http.StatusOK, resp)
}
