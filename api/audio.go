package api

import (
	"giligili/model"
	"giligili/service/audioService"
	"github.com/gin-gonic/gin"
)

// CreateAudio 音频投稿
func CreateAudio(c *gin.Context) {
	value, _ := c.Get("user")
	user, _ := value.(*model.User)
	service := audioService.CreateAudioService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(user)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ShowAudio 音频详情接口
func ShowAudio(c *gin.Context) {
	service := audioService.ShowAudioService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

func GetAudio(c *gin.Context) {
	service := audioService.ShowAudioService{}
	res := service.GetAudio(c.Param("id"))
	c.JSON(200, res)
}

// ListAllAudio 音频列表接口
func ListAllAudio(c *gin.Context) {
	service := audioService.ListAudioService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ListAll()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ListPassedAudios 已通过音频列表接口
func ListPassedAudios(c *gin.Context) {
	service := audioService.ListAudioService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ListPassed()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ListAllAudio 音频列表接口
func ListNotPassedAudios(c *gin.Context) {
	service := audioService.ListAudioService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ListNotPassed()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UpdateAudio 更新音频的接口
func UpdateAudio(c *gin.Context) {
	service := audioService.UpdateAudioService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteAudio 删除音频的接口
func DeleteAudio(c *gin.Context) {
	userValue, _ := c.Get("user")
	user, _ := userValue.(*model.User)
	service := audioService.DeleteAudioService{}
	res := service.Delete(user, c.Param("id"))
	c.JSON(200, res)
}

// 获取用户音频列表接口
func ListPassedAudioByUser(c *gin.Context) {
	userId := c.Param("id")
	service := audioService.ListAudioService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ListPassedByUser(userId)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
