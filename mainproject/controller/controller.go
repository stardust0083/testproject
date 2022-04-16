package controller

import (
	"context"
	GETAREA "getArea"
	GETAREAPB "getArea/proto"
	"mainproject/utils"

	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/asim/go-micro/plugins/transport/grpc/v4"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
)

func GetArea(ctx *gin.Context) {
	reg := consul.NewRegistry()
	ser := grpc.NewTransport()
	microService := micro.NewService(
		micro.Registry(reg),
		micro.Transport(ser),
		micro.Name(GETAREA.Service),
		micro.Version(GETAREA.Version),
	)
	microService.Init()
	client := GETAREAPB.NewGetAreaService(GETAREA.Service, microService.Client())
	rsp, err := client.Call(context.Background(), &GETAREAPB.CallRequest{})
	if err != nil {

		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
	}
	ctx.JSON(200, rsp)
}
