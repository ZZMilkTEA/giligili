package audioService

import (
	"giligili/model"
	"giligili/serializer"
)

// ListAudioService 音频列表服务
type ListAudioService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

// ListAll 音频列表
func (service *ListAudioService) ListAll() serializer.Response {
	audios := []model.Audio{}
	total := 0

	if service.Limit <= 0 {
		service.Limit = 6
	}

	if err := model.DB.Model(model.Audio{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Start).Find(&audios).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildAudios(audios), uint(total))
}

// List 审核通过的音频列表
func (service *ListAudioService) ListPassed() serializer.Response {
	audios := []model.Audio{}
	total := 0

	if service.Limit <= 0 {
		service.Limit = 6
	}

	if err := model.DB.Model(model.Audio{}).Where("status = 'passed'").Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Start).Where("status = 'passed'").Find(&audios).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildAudios(audios), uint(total))
}

// List 审核未通过的音频列表
func (service *ListAudioService) ListNotPassed() serializer.Response {
	audios := []model.Audio{}
	total := 0

	if service.Limit <= 0 {
		service.Limit = 6
	}

	if err := model.DB.Model(model.Audio{}).Where("status = 'pending_review'").Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Start).Where("status = 'pending_review'").Find(&audios).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildAudios(audios), uint(total))
}

// List 用户通过的音频列表
func (service *ListAudioService) ListPassedByUser(uid string) serializer.Response {
	audios := []model.Audio{}
	total := 0

	if service.Limit <= 0 {
		service.Limit = 6
	}

	if err := model.DB.Model(model.Audio{}).Where("user_id = ? AND status = 'passed'", uid).
		Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Start).Where("user_id = ? AND status = 'passed'", uid).
		Find(&audios).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildAudios(audios), uint(total))
}

// List 用户的音频列表
func (service *ListAudioService) ListByUser(uid string) serializer.Response {
	audios := []model.Audio{}
	total := 0

	if service.Limit <= 0 {
		service.Limit = 6
	}

	if err := model.DB.Model(model.Audio{}).Where("user_id = ?", uid).Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Start).Where("user_id = ?", uid).Find(&audios).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildAudios(audios), uint(total))
}
