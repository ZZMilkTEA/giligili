package api

import (
	"giligili/model"
	"giligili/service/videoService"
	"github.com/gin-gonic/gin"
)

func DoReview(c *gin.Context) {
	service := videoService.DoVideoReviewService{}
	videoId := c.Param("id")
	if err := c.ShouldBind(&service); err == nil {
		reviewer, _ := c.Get("user")
		reviewerId := reviewer.(*model.User).ID
		if err != nil {
			c.JSON(200, ErrorResponse(err))
		}
		res := service.ChangeVideoStatus(reviewerId, videoId)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
