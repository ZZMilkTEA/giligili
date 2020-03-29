package service

import (
	"giligili/model"
	"giligili/serializer"
)

// ShowUserService 投稿详情的服务
type ShowUserService struct {
}

// Show 用户
func (service *ShowUserService) Show(id string) serializer.Response {
	var user model.User
	err := model.DB.First(&user, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "用户不存在",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildUser(user),
	}
}
