package service

import (
	"giligili/model"
	"giligili/serializer"
)

type ChangeUserPermissionService struct {
	UserID     uint `form:"user_id" json:"user_id,string" binding:"required"`
	Permission uint `form:"permission" json:"permission,string" binding:"required"`
}

//更改用户权限
func (service *ChangeUserPermissionService) ChangeUserPermission() serializer.Response {
	var user model.User
	err := model.DB.First(&user, service.UserID).Error
	if err != nil {
		return serializer.Response{
			Status: 40004,
			Msg:    "用户不存在",
			Error:  err.Error(),
		}
	}

	err = model.DB.Model(&user).Update("permission", service.Permission).Error
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "用户权限更改失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildUser(user),
	}
}
