package videoService

import (
	"giligili/model"
	"giligili/serializer"
)

// ListVideoService 视频列表服务
type ListVideoService struct {
	Limit int    `form:"limit"`
	Start int    `form:"start"`
	Kind  string `form:"kind"`
}

// ListAll 视频列表
func (service *ListVideoService) ListAll() serializer.Response {
	videos := []model.Video{}
	total := 0

	if service.Limit <= 0 {
		service.Limit = 6
	}

	if err := model.DB.Model(model.Video{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Start).Find(&videos).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildVideos(videos), uint(total))
}

// List 审核通过的视频列表
func (service *ListVideoService) ListPassed() serializer.Response {
	videos := []model.Video{}
	total := 0

	if service.Limit <= 0 {
		service.Limit = 6
	}

	if err := model.DB.Model(model.Video{}).Where(&model.Video{Kind: service.Kind, Status: "passed"}).
		Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Start).Where(&model.Video{Kind: service.Kind, Status: "passed"}).
		Find(&videos).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildVideos(videos), uint(total))
}

// List 审核未通过的视频列表
func (service *ListVideoService) ListNotPassed() serializer.Response {
	videos := []model.Video{}
	total := 0

	if service.Limit <= 0 {
		service.Limit = 6
	}

	if err := model.DB.Model(model.Video{}).Where("status = 'pending_review'").Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Start).Where("status = 'pending_review'").Find(&videos).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildVideos(videos), uint(total))
}

// List 用户通过的视频列表
func (service *ListVideoService) ListPassedByUser(uid string) serializer.Response {
	videos := []model.Video{}
	total := 0

	if service.Limit <= 0 {
		service.Limit = 6
	}

	if err := model.DB.Model(model.Video{}).Where("user_id = ? AND status = 'passed'", uid).
		Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Start).Where("user_id = ? AND status = 'passed'", uid).
		Find(&videos).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildVideos(videos), uint(total))
}

// List 用户的视频列表
func (service *ListVideoService) ListByUser(uid string) serializer.Response {
	videos := []model.Video{}
	total := 0

	if service.Limit <= 0 {
		service.Limit = 6
	}

	if err := model.DB.Model(model.Video{}).Where("user_id = ?", uid).Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Start).Where("user_id = ?", uid).Find(&videos).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildVideos(videos), uint(total))
}
