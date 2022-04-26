package handler

import (
	"context"
	"image/color"
	"io"
	"mainproject/models"
	"time"

	"github.com/afocus/captcha"
	log "go-micro.dev/v4/logger"

	pb "getImageCode/proto"
)

type GetImageCode struct{}

func (e *GetImageCode) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	//获取验证码操作对象
	cap := captcha.New()

	//设置字体库
	if err := cap.SetFont("Delttras.ttf"); err != nil {
		panic(err.Error())
	}

	//设置验证码属性
	cap.SetSize(128, 64)
	cap.SetDisturbance(captcha.NORMAL)
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})

	//生成图片
	img, str := cap.Create(6, captcha.ALL)
	//获取到验证码,把验证码结合uuid写入redis中
	redisPool := models.InitRedis()
	redisConn := redisPool.Get()
	redisConn.Do("setex", req.Uuid, 60*5, str)

	//把拿到的图片对象序列化成字节切片传递给web
	rsp.Max = &pb.CallResponse_Point{X: int64(img.Rect.Max.X), Y: int64(img.Rect.Max.Y)}
	rsp.Min = &pb.CallResponse_Point{X: int64(img.Rect.Min.X), Y: int64(img.Rect.Min.Y)}
	rsp.Stride = int64(img.Stride)
	rsp.Pix = []byte(img.Pix)
	return nil
}

func (e *GetImageCode) ClientStream(ctx context.Context, stream pb.GetImageCode_ClientStreamStream) error {
	var count int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Infof("Got %v pings total", count)
			return stream.SendMsg(&pb.ClientStreamResponse{Count: count})
		}
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		count++
	}
}

func (e *GetImageCode) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.GetImageCode_ServerStreamStream) error {
	log.Infof("Received GetImageCode.ServerStream request: %v", req)
	for i := 0; i < int(req.Count); i++ {
		log.Infof("Sending %d", i)
		if err := stream.Send(&pb.ServerStreamResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 250)
	}
	return nil
}

func (e *GetImageCode) BidiStream(ctx context.Context, stream pb.GetImageCode_BidiStreamStream) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&pb.BidiStreamResponse{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
