package minioClient

import (
	"fmt"
	"log"
	"storeObj/global"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var Instance *minio.Client

func Init(gconfig *global.Config) *minio.Client {

	if Instance != nil {
		return Instance
	}

	var err error
	Instance, err = minio.New(
		gconfig.Minio.Endpoint,
		&minio.Options{
			Creds:  credentials.NewStaticV4(gconfig.Minio.AccessKey, gconfig.Minio.SecretKey, ""),
			Secure: gconfig.Minio.UseSSL,
		},
	)
	if err != nil {
		log.Fatalf("初始化 minioClient 失败: %s", err)
	}

	fmt.Printf("初始化Minio成功，%#v\n", Instance) // minioClient is now set up

	return Instance
}
