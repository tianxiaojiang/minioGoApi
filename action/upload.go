package action

import (
	"context"
	"crypto/md5"
	"fmt"
	"mime/multipart"
	"storeObj/global"
	"storeObj/minioClient"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func UploadEndpoint(c *gin.Context) {
	// Multipart form
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	bucket := c.Param("bucket")

	res := make([]string, len(files))
	for i, file := range files {
		// 上传文件至minio

		resItem, err := upload(file, bucket)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 500,
				"msg":  err.Error(),
				"data": nil,
			})
			return
		}

		res[i] = resItem
	}

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": res,
	})
}

func upload(file *multipart.FileHeader, bucketName string) (ret string, err error) {

	gconfig := global.Conf

	minioClient := minioClient.Init(gconfig)

	ret = ""
	s := fmt.Sprintf("%v%v", file.Filename, time.Now().Unix())
	objectName := fmt.Sprintf("%x", md5.Sum([]byte(s))) + "." + strings.Split(file.Filename, ".")[1]
	// contentType := file.Header.Get("Content-Type")

	fileReader, err := file.Open()
	if err != nil {
		return
	}
	defer fileReader.Close()

	// fmt.Println(minioClient)

	// return
	uploadInfo, err := minioClient.PutObject(
		context.Background(),
		bucketName,
		objectName,
		fileReader,
		file.Size,
		minio.PutObjectOptions{
			ContentType: "application/octet-stream",
			// ContentType: contentType,
		})

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully uploaded bytes: ", uploadInfo)

	ret = fmt.Sprintf(gconfig.App.AppUrl + "/" + uploadInfo.Bucket + "/" + uploadInfo.Key)

	return
}
