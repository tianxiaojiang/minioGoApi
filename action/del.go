package action

import (
	"context"
	"fmt"
	"storeObj/global"
	"storeObj/minioClient"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func DelImgData(bucket string, filename string) (err error) {

	minioClient := minioClient.Init(global.Conf)

	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
	}
	err = minioClient.RemoveObject(context.Background(), bucket, filename, opts)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func DelEndpoint(c *gin.Context) {
	bucket := c.Param("bucket")
	filename := c.Param("filename")

	err := DelImgData(bucket, filename)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "ok",
		})
	}
}
