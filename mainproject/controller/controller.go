package controller

import (
	"context"
	GETAREAPB "getArea/proto"
	GETIMAGECODEPB "getImageCode/proto"
	"image"
	"image/png"
	"mainproject/utils"
	"net/http"

	"github.com/afocus/captcha"
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

func GetSession(ctx *gin.Context) {
	resp := make(map[string]interface{})

	//初始化返回值
	resp["errno"] = utils.RECODE_SESSIONERR
	resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)

	//返回数据
	ctx.JSON(200, resp)
}
func GetImageCode(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	reg := consul.NewRegistry()
	ser := grpc.NewTransport()
	microService := micro.NewService(
		micro.Registry(reg),
		micro.Transport(ser),
	)
	microService.Init()
	client := GETIMAGECODEPB.NewGetImageCodeService("go.micro.srv.GetArea", microService.Client())
	rsp, err := client.Call(context.Background(), &GETIMAGECODEPB.CallRequest{Uuid: uuid})

	if err != nil {
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		ctx.JSON(500, rsp)
		return
	}

	//解析返回数据为image,传回前端
	var img image.RGBA
	img.Stride = int(rsp.Stride)

	img.Rect.Min.X = int(rsp.Min.X)
	img.Rect.Min.Y = int(rsp.Min.Y)
	img.Rect.Max.X = int(rsp.Max.X)
	img.Rect.Max.Y = int(rsp.Max.Y)

	img.Pix = []uint8(rsp.Pix)

	var captchaImage captcha.Image
	captchaImage.RGBA = &img

	// 将图片发送给前端
	png.Encode(ctx.Writer, captchaImage)
}
