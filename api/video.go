package api

import (
	"giligili/model"
	"giligili/serializer"
	"giligili/service/ossService"
	"giligili/service/videoService"
	"github.com/gin-gonic/gin"
	"strconv"
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

// ShowVideo 播放视频用接口
func ShowVideo(c *gin.Context) {
	service := videoService.ShowVideoService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

// GetVideo 能对所有视频使用的接口，获取视频全部信息
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

// DeleteVideo 删除用户自己的视频的接口
func DeleteMyVideo(c *gin.Context) {
	userIdStr := c.Param("id")
	temp, _ := strconv.ParseUint(userIdStr, 10, 32)
	userId := uint(temp)

	userStr, _ := c.Get("user")
	user, _ := userStr.(*model.User)

	if user.ID != userId {
		errResponse := serializer.Response{
			Status: 40003,
			Data:   nil,
			Msg:    "只能修改自己的头像",
			Error:  "verify err",
		}
		c.JSON(200, errResponse)
	}

	service := videoService.DeleteVideoService{}
	res := service.Delete(user, c.Param("id"))
	c.JSON(200, res)
}

// 获取用户视频列表接口
func ListPassedVideoByUser(c *gin.Context) {
	userId := c.Param("id")
	service := videoService.ListVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ListPassedByUser(userId)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 获取视频雪碧图
func GetVideoSpritePic(c *gin.Context) {
	videoId := c.Param("id")
	service := ossService.GetVideoSpriteService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateSpritePic(videoId)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
