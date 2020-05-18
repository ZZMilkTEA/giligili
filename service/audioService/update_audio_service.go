package audioService

import (
	"giligili/model"
	"giligili/serializer"
)

// UpdateAudioService 更新音频的服务
type UpdateAudioService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info  string `form:"info" json:"info" binding:"max=300"`
}

// Update 更新音频
func (service *UpdateAudioService) Update(id string) serializer.Response {
	var audio model.Audio
	err := model.DB.First(&audio, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "音频不存在",
			Error:  err.Error(),
		}
	}

	audio.Title = service.Title
	audio.Info = service.Info
	err = model.DB.Save(&audio).Error
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "音频保存失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildAudio(audio),
	}
}
