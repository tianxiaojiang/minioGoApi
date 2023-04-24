package action

import (
	"context"
	"fmt"
	"io/ioutil"
	"storeObj/global"
	"storeObj/minioClient"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func ReadImgData(bucket string, filename string) (data []byte, err error) {

	data = nil

	minioClient := minioClient.Init(global.Conf)

	fmt.Println(bucket, filename)

	object, err := minioClient.GetObject(context.Background(), bucket, filename, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer object.Close()

	data, err = ioutil.ReadAll(object)

	return
}

func GetEndpoint(c *gin.Context) {
	bucket := c.Param("bucket")
	filename := c.Param("filename")

	data, err := ReadImgData(bucket, filename)
	if err != nil {
		c.Data(404, "plain/text", []byte(err.Error()))
	}

	c.Data(200, "image/jpeg", data)
}
