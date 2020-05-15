package api

import (
	Service "giligili/service"
	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	var service = Service.SearchService{}
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		searchType := c.Query("searchType")
		res := service.DoSearch(searchType)
		c.JSON(200, res)
	}
}
