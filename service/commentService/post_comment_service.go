package commentService

import (
	"giligili/model"
	"giligili/serializer"
)

// PostCommentService 发表评论的服务
type PostCommentService struct {
	Content string `form:"content" json:"content" binding:"max=3000"`
	//FromUserId    string `form:"from_user_id" json:"from_user_id"`
}

// CreateCommentToVideo 创建评论
func (service *PostCommentService) CreateCommentToVideo(user *model.User, videoId uint) serializer.Response {
	comment := model.Comment{
		MediaId:    videoId,
		MediaType:  "video",
		Content:    service.Content,
		FromUserId: user.ID,
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
		Data: serializer.BuildComment(comment, comment.FromUserId),
	}
}
