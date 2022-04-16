package main

import "github.com/gin-gonic/gin"
import "mainproject/controller"
func main() {
	//初始化路由
	router := gin.Default()

	//映射静态资源
	router.Static("/home", "view")

	r1 := router.Group("api/v1.0")
	{
		r1.GET("/areas", controller.GetArea)
	}

	//开启监听
	router.Run(":8080")
}
