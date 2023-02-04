package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"io"
	"mini_tiktok/pkg/configs/config"
)

func UploadFile(ctx context.Context, objName string, reader io.Reader, fileSize int64) (minio.UploadInfo, error) {
	conf := config.GlobalConfigs.MinioConfig
	return MinioClient.PutObject(ctx, conf.BucketName, objName, reader, fileSize,
		minio.PutObjectOptions{ContentType: ""})
}
