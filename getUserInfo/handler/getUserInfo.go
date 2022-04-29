package handler

import (
	"context"
	"io"
	"mainproject/models"
	"mainproject/utils"
	"time"

	log "go-micro.dev/v4/logger"

	pb "getUserInfo/proto"
)

type GetUserInfo struct{}

func (e *GetUserInfo) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received GetUserInfo.Call request: %v", req)
	db, err := models.InitDb()
	var usrdb models.User
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return err
	}
	err = db.Where("name=?", req.Name).First(&usrdb).Error
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return err
	}
	var usr pb.UserData
	usr.Name = usrdb.Name
	usr.RealName = usrdb.Real_name
	usr.Mobile = usrdb.Mobile
	usr.IdCard = usrdb.Id_card
	usr.AvatarUrl = usrdb.Avatar_url
	usr.UserId = int32(usrdb.ID)

	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	rsp.Data = &usr

	return nil
}

func (e *GetUserInfo) ClientStream(ctx context.Context, stream pb.GetUserInfo_ClientStreamStream) error {
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

func (e *GetUserInfo) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.GetUserInfo_ServerStreamStream) error {
	log.Infof("Received GetUserInfo.ServerStream request: %v", req)
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

func (e *GetUserInfo) BidiStream(ctx context.Context, stream pb.GetUserInfo_BidiStreamStream) error {
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
