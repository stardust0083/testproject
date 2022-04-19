package main

import (
	"mainproject/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	//初始化路由
	router := gin.Default()

	//映射静态资源
	router.Static("/home", "mainproject/view")

	r1 := router.Group("api/v1.0")
	{
		r1.GET("/areas", controller.GetArea)
		r1.GET("/test", controller.Test)
	}

	//开启监听
	router.Run(":8080")
}
