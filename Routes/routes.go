package main

import (
	"gin/App/Controller"
	"github.com/gin-gonic/gin"
)

func main() {

	//创建路由
	router := gin.Default()

	//设置路由分组
	testGroup := router.Group("abc")
	{
		testGroup.GET("/info", Controller.Info)
		testGroup.GET("/all", Controller.All)
		testGroup.GET("/insert", Controller.Insert)
		testGroup.GET("/delete", Controller.Delete)
		testGroup.GET("/update/:id", Controller.Update)
	}

	//设置端口号，默认8080
	router.Run(":10086")
}
