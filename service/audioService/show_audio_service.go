package audioService

import (
	"giligili/model"
	"giligili/serializer"
)

// ShowAudioService 投稿详情的服务
type ShowAudioService struct {
}

// Show 音频
func (service *ShowAudioService) Show(id string) serializer.Response {
	var audio model.Audio
	err := model.DB.First(&audio, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "音频不存在",
			Error:  err.Error(),
		}
	}

	if audio.Status != "passed" {
		return serializer.Response{
			Status: 403,
			Msg:    "音频未审核通过",
			Error:  "audio error",
		}
	}
	//处理音频被观看的一系列问题
	audio.AddView()

	return serializer.Response{
		Data: serializer.BuildAudio(audio),
	}
}

func (service *ShowAudioService) GetAudio(id string) serializer.Response {
	var audio model.Audio
	err := model.DB.First(&audio, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "音频不存在",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildAudio(audio),
	}
}
