package utils

import "time"

var (
	G_redis_maxidel     = 20
	G_redis_maxactive   = 50
	G_redis_idletimeout = time.Duration(60 * 5)
	G_redis_addr        = "127.0.0.1"
	G_redis_port        = "6379"
)
