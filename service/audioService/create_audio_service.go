package audioService

import (
	"giligili/model"
	"giligili/serializer"
)

// CreateAudioService 音频投稿的服务
type CreateAudioService struct {
	Title       string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Info        string `form:"info" json:"info" binding:"max=3000"`
	URL         string `form:"url" json:"url"`
	Avatar      string `form:"avatar" json:"avatar"`
	Kind        string `form:"kind" json:"kind"`
	UserId      uint   `form:"user_id" json:"user_id"`
	FromOutside bool   `form:"from_outside" json:"from_outside"`
}

// Create 创建音频
func (service *CreateAudioService) Create(user *model.User) serializer.Response {
	audio := model.Audio{
		Title:       service.Title,
		Info:        service.Info,
		URL:         service.URL,
		Avatar:      service.Avatar,
		Kind:        service.Kind,
		FromOutside: service.FromOutside,
		UserId:      user.ID,
	}

	err := model.DB.Create(&audio).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "音频保存失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildAudio(audio),
	}
}
