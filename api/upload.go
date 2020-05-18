package api

import (
	"giligili/service/ossService"

	"github.com/gin-gonic/gin"
)

// MediaUploadToken 上传授权
func MediaUploadToken(c *gin.Context) {
	service := ossService.UploadTokenService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.PostVideo()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
