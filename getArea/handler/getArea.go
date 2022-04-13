package handler

import (
	"context"

	log "go-micro.dev/v4/logger"

	pb "getArea/proto"
	"mainproject/utils"
)

type GetArea struct{}

func (e *GetArea) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received GetArea.Call request: %v", req)
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)
	redisConfigMap := map[string]string{
		"key":   utils.G_server_name,
		"conn":  utils.G_redis_addr + ":" + utils.G_redis_port,
		"dbNum": utils.G_redis_dbnum,
	}
	return nil
}
