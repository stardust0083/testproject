package handler

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"mainproject/models"
	"mainproject/utils"
	"time"

	"github.com/garyburd/redigo/redis"
	log "go-micro.dev/v4/logger"

	pb "postRet/proto"
)

type PostRet struct{}

func (e *PostRet) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	redisConn := models.InitRedis().Get()
	code, err := redis.String(redisConn.Do("get", req.Mobile+"_code"))
	if err != nil {
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		return err
	}

	if code != req.SmsCode {
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		return errors.New("captcha doesn't match")
	}

	db, err := models.InitDb()
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return errors.New("init SQL Failed")
	}
	var user models.User
	user.Name = req.Mobile
	user.Mobile = req.Mobile

	m5Pwd := fmt.Sprintf("%x", md5.Sum([]byte(req.Password)))
	user.Password_hash = m5Pwd

	err = db.Create(&user).Error
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return err
	}
	return nil
}

func (e *PostRet) ClientStream(ctx context.Context, stream pb.PostRet_ClientStreamStream) error {
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

func (e *PostRet) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.PostRet_ServerStreamStream) error {
	log.Infof("Received PostRet.ServerStream request: %v", req)
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

func (e *PostRet) BidiStream(ctx context.Context, stream pb.PostRet_BidiStreamStream) error {
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
