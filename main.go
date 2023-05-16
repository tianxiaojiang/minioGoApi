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

	// 文件上传, 为了兼容小程序，使用POST
	router.POST("/:bucket", action.UploadEndpoint)

	// 文件读取
	router.GET("/:bucket/:filename", action.GetEndpoint)

	// 文件删除
	router.DELETE("/:bucket/:filename", action.DelEndpoint)

	// 对象打标签 tk=:tagname&tv=:tagvalue
	router.PATCH("/:bucket/:filename", action.TagEndpoint)

	router.Run(":" + global.Conf.App.Port)
}
