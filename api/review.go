package api

import (
	"giligili/service"
	"giligili/token"
	"github.com/gin-gonic/gin"
)

func DoReview(c *gin.Context) {
	service := service.DoReviewService{}
	if err := c.ShouldBind(&service); err == nil {
		strToken := c.GetHeader("token")
		reviewerID, err := token.GetLoggedUserId(strToken)
		if err != nil {
			c.JSON(200, ErrorResponse(err))
		}
		res := service.ChangeVideoStatus(reviewerID)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
