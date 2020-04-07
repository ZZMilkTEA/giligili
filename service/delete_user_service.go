package service

import (
	"giligili/model"
	"giligili/serializer"
)

// DeleteUserService 删除投稿的服务
type DeleteUserService struct {
}

// Delete 删除视频
func (service *DeleteUserService) Delete(id string) serializer.Response {
	var user model.User
	err := model.DB.First(&user, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "用户不存在",
			Error:  err.Error(),
		}
	}

	err = model.DB.Delete(&user).Error
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "用户注销失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{}
}
