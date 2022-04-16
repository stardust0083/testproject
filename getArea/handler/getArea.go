package handler

import (
	"context"
	"encoding/json"

	"github.com/garyburd/redigo/redis"
	log "go-micro.dev/v4/logger"

	pb "getArea/proto"
	"mainproject/models"
	"mainproject/utils"
)

type GetArea struct{}

func (e *GetArea) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received GetArea.Call request: %v", req)
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Errno)
	// redisConfigMap := map[string]string{
	// 	"key":   utils.G_server_name,
	// 	"conn":  utils.G_redis_addr + ":" + utils.G_redis_port,
	// 	"dbNum": utils.G_redis_dbnum,
	// }
	// log.Info(redisConfigMap)
	// redisConfigMap, _ = json.Marshal(redisConfigMap)
	redisClient := models.InitRedis().Get()
	defer redisClient.Close()
	areaResp, err := redis.Bytes(redisClient.Do("get", "areas"))
	if err != nil {
		log.Infof("Read Redis Failed")
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Errno)
		return nil
	}
	var areas []models.Area
	if len(areaResp) == 0 {
		//第一次从mysql中获取数据 调用封装的函数
		mysqlClient, _ := models.InitDb()
		areas, err = models.GetAllArea(mysqlClient)
		if err != nil {
			rsp.Errno = utils.RECODE_DBERR
			rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
			return err
		}
		//对areas数据编码
		areaBytes, err := json.Marshal(areas)
		if err != nil {
			rsp.Errno = utils.RECODE_DATAERR
			rsp.Errno = utils.RecodeText(utils.RECODE_DATAERR)
			return err
		}

		//存储到redis中
		_, err = redis.String(redisClient.Do("set", "areas", areaBytes))
		if err != nil {
			rsp.Errno = utils.RECODE_DATAERR
			rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
			return err
		}
	} else {
		err = json.Unmarshal(areaResp, &areas)
		if err != nil {
			rsp.Errno = utils.RECODE_DATAERR
			rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
			return err
		}
	}
	for _, v := range areas {
		temp := pb.CallResponse_Address{
			Aid:   int32(v.Id),
			Aname: v.Name,
		}
		rsp.Data = append(rsp.Data, &temp)
	}

	return nil
}
