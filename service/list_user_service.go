package service

import (
	"giligili/model"
	"giligili/serializer"
)

type ListUserService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

//List 所有用户
func (service *ListUserService) List() serializer.Response {
	users := []model.User{}
	total := 0

	if service.Limit == 0 {
		service.Limit = 9
	}

	if err := model.DB.Model(model.User{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Start).Find(&users).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildUsers(users), uint(total))
}
