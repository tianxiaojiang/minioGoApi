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

	// 打上标签 status => init
	// 标签值约定：init 初始化上传、confirm 业务确认使用
	// 		init 生命周期只有3天
	// 		confirm 生命周期不限
	// 业务确认使用时，删除此标签。生命周期则不限
	userTags := map[string]string{
		"status": "init",
	}

	// return
	uploadInfo, err := minioClient.PutObject(
		context.Background(),
		bucketName,
		objectName,
		fileReader,
		file.Size,
		minio.PutObjectOptions{
			// ContentType: "application/octet-stream",
			UserTags:    userTags,
			ContentType: file.Header.Get("ContentType"),
		})

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully uploaded bytes: ", uploadInfo)

	ret = fmt.Sprintf(gconfig.App.AppUrl + "/" + uploadInfo.Bucket + "/" + uploadInfo.Key)

	return
}
