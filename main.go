package main

import (

	// "encoding/hex"
	"storeObj/action"
	"storeObj/bootstrap"
	"storeObj/global"
	"storeObj/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	bootstrap.InitializeConfig()

	router := gin.New()

	router.Use(handler.Auth())

	// 文件上传
	router.PUT("/:bucket", action.UploadEndpoint)

	// 文件读取
	router.GET("/:bucket/:filename", action.GetEndpoint)

	// 文件删除
	router.DELETE("/:bucket/:filename", action.DelEndpoint)

	router.Run(":" + global.Conf.App.Port)
}
