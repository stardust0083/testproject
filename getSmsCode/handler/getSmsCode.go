package handler

import (
	"context"
	"fmt"
	"io"
	"mainproject/models"
	"mainproject/utils"
	"math/rand"
	"time"

	"github.com/garyburd/redigo/redis"
	log "go-micro.dev/v4/logger"

	pb "getSmsCode/proto"
)

type GetSmsCode struct{}

func (e *GetSmsCode) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received GetSmsCode.Call request: %v", req)
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)

	log.Infof("Read SQL")
	mysqlClient, _ := models.InitDb()
	defer mysqlClient.Close()
	log.Infof("Init SQL")
	var user models.User
	err := mysqlClient.Find(&user, "Mobile = ?", req.Phone).Error
	if err == nil {
		log.Infof("User Exists")
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		return nil
	}

	redisCli := models.InitRedis().Get()
	defer redisCli.Close()
	capCode, err := redis.String(redisCli.Do("get", req.Uuid))
	log.Infof(capCode)
	if err != nil {
		log.Infof("Captcha Not Found")
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		return err
	}
	if capCode != req.Text {
		log.Infof("Captcha Error")
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		return err
	}

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06d", rnd.Int31n(1000000))
	redisCli.Do("setex", req.Phone+"_code", 60*5, vcode)
	return nil
}

func (e *GetSmsCode) ClientStream(ctx context.Context, stream pb.GetSmsCode_ClientStreamStream) error {
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

func (e *GetSmsCode) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.GetSmsCode_ServerStreamStream) error {
	log.Infof("Received GetSmsCode.ServerStream request: %v", req)
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

func (e *GetSmsCode) BidiStream(ctx context.Context, stream pb.GetSmsCode_BidiStreamStream) error {
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
