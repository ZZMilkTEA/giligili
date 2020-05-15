package commentService

import (
	"giligili/model"
	"giligili/serializer"
)

type ListCommentService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
	//MediaType string `form:"type" json:"type"`
	//MediaId   uint   `form:"mediaId" json:"media_id"`
}

//ListAll 列出所有评论
func (service *ListCommentService) ListAll() serializer.Response {
	comments := []model.Comment{}
	total := 0

	if service.Limit == 0 {
		service.Limit = 20
	}

	if err := model.DB.Model(model.Comment{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Start).Find(&comments).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildComments(comments), uint(total))
}

//ListByMediaId 通过媒体列出评论
func (service *ListCommentService) ListByVideoId(videoId string) serializer.Response {
	comments := []model.Comment{}
	total := 0

	if service.Limit <= 0 {
		service.Limit = 20
	}

	if err := model.DB.Model(model.Comment{}).
		Where("media_type = ? AND media_id = ?", "video", videoId).
		Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Start).
		Where("media_type = ? AND media_id = ?", "video", videoId).
		Find(&comments).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildComments(comments), uint(total))
}

//ListByMediaId 通过用户列出评论
func (service *ListCommentService) ListByUser(uid string) serializer.Response {
	comments := []model.Comment{}
	total := 0

	if service.Limit == 0 {
		service.Limit = 20
	}

	if err := model.DB.Model(model.Comment{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Start).Find(&comments).
		Where("from_user_id = ?", uid).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildComments(comments), uint(total))
}
