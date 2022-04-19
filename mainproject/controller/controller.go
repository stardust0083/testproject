package controller

import (
	"context"
	GETAREAPB "getArea/proto"
	"mainproject/utils"
	"net/http"

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
	)
	microService.Init()
	client := GETAREAPB.NewGetAreaService("go.micro.srv.GetArea", microService.Client())
	rsp, err := client.Call(context.Background(), &GETAREAPB.CallRequest{})
	if err != nil {
		print("fuck call", rsp)
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		// ctx.JSON(http.StatusNotAcceptable, gin.H{"status": http.StatusOK, "message": "fuck!"})
		// return
	}
	// print(rsp)
	// ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Im here!"})
	ctx.JSON(200, rsp)
}

func Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Im here!"})
}
