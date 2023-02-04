package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"mini_tiktok/pkg/configs/config"
)

var (
	MinioClient *minio.Client
)

func Init() {
	conf := config.GlobalConfigs.MinioConfig
	endpoint := conf.Endpoint
	accessKeyID := conf.AccessKeyID
	secretAccessKey := conf.SecretAccessKey
	useSSL := conf.UseSSL

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
	MinioClient = minioClient
}
