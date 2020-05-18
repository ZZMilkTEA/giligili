package api

import (
	"giligili/model"
	"giligili/service/commentService"
	"github.com/gin-gonic/gin"
	"strconv"
)

// PostComment 发表评论
func PostComment(c *gin.Context) {
	value, _ := c.Get("user")
	user, _ := value.(*model.User)

	videoIdStr := c.Param("id")
	temp, _ := strconv.ParseUint(videoIdStr, 10, 32)
	videoId := uint(temp)
	service := commentService.PostCommentService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateCommentToVideo(user, videoId)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//----------------------- ListCommentByMedia 从媒体列出评论列表----------------
func ListCommentsByMediaId(c *gin.Context) {
	videoId := c.Param("id")
	service := commentService.ListCommentService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ListByMeidaId(videoId)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ListCommentByUser 从用户列出评论列表
func ListCommentByUser(c *gin.Context) {
	userId := c.Param("id")
	service := commentService.ListCommentService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ListByUser(userId)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
