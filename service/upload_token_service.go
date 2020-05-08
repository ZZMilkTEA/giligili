package service

import (
	"giligili/serializer"
	"github.com/google/uuid"
	"mime"
	"os"
	"path/filepath"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// UploadTokenService 获得上传oss token的服务
type UploadTokenService struct {
	Filename string `form:"filename" json:"filename"`
}

// Post 创建token
func (service *UploadTokenService) Post() serializer.Response {
	client, err := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "OSS配置错误",
			Error:  err.Error(),
		}
	}

	// 获取存储空间。
	bucket, err := client.Bucket(os.Getenv("OSS_BUCKET"))
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "OSS配置错误",
			Error:  err.Error(),
		}
	}

	// 获取扩展名
	ext := filepath.Ext(service.Filename)
	mimeType := mime.TypeByExtension(ext)

	// 带可选参数的签名直传。
	options := []oss.Option{
		oss.ContentType(mimeType),
		//oss.ContentType("image/png"),
	}

	//根据扩展名分配上传路径
	var key string
	switch ext {
	case ".jpg", ".png":
		key = "upload/avatar/" + uuid.Must(uuid.NewRandom()).String() + ext
	case ".mp4":
		key = "upload/videos/" + uuid.Must(uuid.NewRandom()).String() + ext
	}

	// 签名直传。
	signedPutURL, err := bucket.SignURL(key, oss.HTTPPut, 600, options...)
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "OSS配置错误",
			Error:  err.Error(),
		}
	}
	// 查看资源
	signedGetURL, err := bucket.SignURL(key, oss.HTTPGet, 600)
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "OSS配置错误",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: map[string]string{
			"key": key,
			"put": signedPutURL,
			"get": signedGetURL,
		},
	}
}
