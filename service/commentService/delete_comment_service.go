package commentService

import (
	"giligili/model"
	"giligili/serializer"
)

// DeleteUserService 删除投稿的服务
type DeleteCommentService struct {
}

// Delete 注销用户
func (service *DeleteCommentService) Delete(id string) serializer.Response {
	var comment model.Comment
	err := model.DB.First(&comment, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "评论不存在",
			Error:  err.Error(),
		}
	}

	err = model.DB.Delete(&comment).Error
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "评论删除失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{}
}
