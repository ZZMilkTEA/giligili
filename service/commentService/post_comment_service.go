package commentService

import (
	"giligili/model"
	"giligili/serializer"
)

// PostCommentService 发表评论的服务
type PostCommentService struct {
	MediaId   uint   `form:"media_id" json:"media_id" binding:"required"`
	MediaType string `form:"media_type" json:"media_type" binding:"required"`
	Content   string `form:"content" json:"content" binding:"max=3000"`
	//FromUserId    string `form:"from_user_id" json:"from_user_id"`
	//Avatar string `form:"avatar" json:"avatar"`
}

// Create 创建评论
func (service *PostCommentService) Create(user *model.User) serializer.Response {
	comment := model.Comment{
		MediaId:    service.MediaId,
		MediaType:  service.MediaType,
		Content:    service.Content,
		FromUserId: user.ID,
		Avatar:     user.Avatar,
	}

	err := model.DB.Create(&comment).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "评论发表失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildComment(comment),
	}
}
