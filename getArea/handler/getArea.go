package handler

import (
	"context"

	log "go-micro.dev/v4/logger"

	pb "getArea/proto"
)

type GetArea struct{}

func (e *GetArea) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received GetArea.Call request: %v", req)
	rsp.Errmsg = "Hello " + req.Name
	return nil
}
