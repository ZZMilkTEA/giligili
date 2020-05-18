package userService

import (
	"giligili/model"
	"giligili/serializer"
)

type ChangeUserInfoService struct {
	Permission uint   `form:"permission" json:"permission"`
	Status     string `form:"status" json:"status"`
	Avatar     string `form:"avatar" json:"avatar"`
}

//更改用户信息
func (service *ChangeUserInfoService) ChangeUserInfo(userId uint) serializer.Response {

	var user model.User
	err := model.DB.First(&user, userId).Error
	if err != nil {
		return serializer.Response{
			Status: 40004,
			Msg:    "用户不存在",
			Error:  err.Error(),
		}
	}

	err = model.DB.Model(&user).Updates(model.User{Permission: service.Permission, Status: service.Status, Avatar: service.Avatar}).Error
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "用户信息更改失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildUser(user),
	}
}

func (service *ChangeUserInfoService) ChangeUserAvatar(userId uint) serializer.Response {

	var user model.User
	err := model.DB.First(&user, userId).Error
	if err != nil {
		return serializer.Response{
			Status: 40004,
			Msg:    "用户不存在",
			Error:  err.Error(),
		}
	}

	err = model.DB.Model(&user).Updates(model.User{Avatar: service.Avatar}).Error
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "用户头像更改失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildUser(user),
	}
}
