package action

import (
	"context"
	"fmt"
	"storeObj/global"
	"storeObj/minioClient"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/tags"
)

func TagEndpoint(c *gin.Context) {
	// Multipart form
	bucketName := c.Param("bucket")
	objName := c.Param("filename")
	// 要设置的tag key
	tagKey := c.Query("tk")
	// 要设置的tag value
	tagVal := c.Query("tv")

	gconfig := global.Conf
	minioClient := minioClient.Init(gconfig)

	fmt.Printf("tagkey: %s, tagval: %s", tagKey, tagVal)

	tags, err := tags.NewTags(map[string]string{
		tagKey: tagVal,
	}, true)

	if err != nil {
		c.JSON(500, gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	err = minioClient.PutObjectTagging(context.Background(), bucketName, objName, tags, minio.PutObjectTaggingOptions{})
	if err != nil {
		c.JSON(500, gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": nil,
	})
}
