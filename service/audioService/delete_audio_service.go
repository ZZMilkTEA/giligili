package audioService

import (
	"fmt"
	"giligili/model"
	"giligili/serializer"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

// DeleteAudioService 删除投稿的服务
type DeleteAudioService struct {
}

// Delete 删除音频
func (service *DeleteAudioService) Delete(user *model.User, id string) serializer.Response {
	//---------------------------前期准备----------------------------------
	var audio model.Audio
	err := model.DB.First(&audio, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "音频不存在",
			Error:  err.Error(),
		}
	}

	if user.Permission < 1 && user.ID != audio.UserId {
		return serializer.Response{
			Status: 40001,
			Msg:    "无法删除",
			Error:  "您没有删除此音频的权限",
		}
	}

	//---------------------------------执行删除----------------------------------------
	//数据库行删除
	err = model.DB.Delete(&audio).Error
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "音频删除失败",
			Error:  err.Error(),
		}
	}
	//创建OSSClient实例
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

	//执行OSS删除
	delRes, err := bucket.DeleteObjects([]string{audio.URL, audio.Avatar})
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	fmt.Println("Deleted Objects:", delRes.DeletedObjects)

	return serializer.Response{}
}
