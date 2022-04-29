package main

import (
	"fmt"
	"mainproject/controller"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// 在路由之前调用上面代码
		c.Next()
		//在路由之后调用
		latency := time.Since(t)
		fmt.Println(latency)
	}
}
func main() {
	//初始化路由
	router := gin.Default()
	router.Use(Logger())
	//映射静态资源
	router.Static("/home", "./view")
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	r1 := router.Group("api/v1.0")
	{
		r1.GET("/areas", controller.GetArea)
		r1.GET("/test", controller.Test)
		r1.GET("/session", controller.GetSession)
		r1.GET("/imagecode/:uuid", controller.GetImageCode)
		r1.GET("/smscode/:mobile", controller.GetSmscd)
		r1.POST("/users", controller.PostRet)
	}

	//开启监听
	router.Run(":8080")
}
