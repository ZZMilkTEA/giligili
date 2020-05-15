package api

import (
	"giligili/model"
	"giligili/service/videoService"
	"github.com/gin-gonic/gin"
)

// CreateVideo 视频投稿
func CreateVideo(c *gin.Context) {
	value, _ := c.Get("user")
	user, _ := value.(*model.User)
	service := videoService.CreateVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(user)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ShowVideo 视频详情接口
func ShowVideo(c *gin.Context) {
	service := videoService.ShowVideoService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

func GetVideo(c *gin.Context) {
	service := videoService.ShowVideoService{}
	res := service.GetVideo(c.Param("id"))
	c.JSON(200, res)
}

// ListAllVideo 视频列表接口
func ListAllVideo(c *gin.Context) {
	service := videoService.ListVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ListAll()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ListPassedVideos 已通过视频列表接口
func ListPassedVideos(c *gin.Context) {
	service := videoService.ListVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ListPassed()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ListAllVideo 视频列表接口
func ListNotPassedVideos(c *gin.Context) {
	service := videoService.ListVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ListNotPassed()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UpdateVideo 更新视频的接口
func UpdateVideo(c *gin.Context) {
	service := videoService.UpdateVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteVideo 删除视频的接口
func DeleteVideo(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*model.User)
	service := videoService.DeleteVideoService{}
	res := service.Delete(user, c.Param("id"))
	c.JSON(200, res)
}

// 获取用户视频列表接口
func ListVideoByUser(c *gin.Context) {
	service := videoService.ListVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ListByUser(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
