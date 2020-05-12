package service

import (
	"giligili/model"
	"giligili/serializer"
)

type SearchService struct {
	Keyword string `form:"keyword" json:"keyword" binding:"required"`
	Limit   int    `form:"limit"`
	Start   int    `form:"start"`
}

func (service *SearchService) DoSearch(searchType string) serializer.Response {
	total := 0
	searchKeyword := "%" + service.Keyword + "%"
	if service.Limit <= 0 {
		service.Limit = 6
	}

	if searchType == "video" {
		var videos []model.Video

		//记录符合条件记录总数
		if err := model.DB.Table("videos").Where("title LIKE ? AND status = 'passed'", searchKeyword).Count(&total).Error; err != nil {
			return serializer.Response{
				Status: 50000,
				Msg:    "数据库连接错误",
				Error:  err.Error(),
			}
		}
		//根据请求查找复合条件中指定区间的记录
		if err := model.DB.Table("videos").Limit(service.Limit).Offset(service.Start).
			Where("title LIKE ? AND status = 'passed'", searchKeyword).Find(&videos).Error; err != nil {
			return serializer.Response{
				Status: 50000,
				Msg:    "数据库连接错误",
				Error:  err.Error(),
			}
		}
		res := serializer.BuildVideos(videos)
		return serializer.BuildListResponse(res, uint(total))
	}

	if searchType == "user" {
		var users []model.User

		//记录符合条件记录总数
		if err := model.DB.Table("users").Where("nickname LIKE ? AND status = 'active'", searchKeyword).Count(&total).Error; err != nil {
			return serializer.Response{
				Status: 50000,
				Msg:    "数据库连接错误",
				Error:  err.Error(),
			}
		}

		//根据请求查找复合条件中指定区间的记录
		if err := model.DB.Table("users").Limit(service.Limit).Offset(service.Start).
			Where("nickname LIKE ? AND status = 'active'", searchKeyword).Find(&users).Error; err != nil {
			return serializer.Response{
				Status: 50000,
				Msg:    "数据库连接错误",
				Error:  err.Error(),
			}
		}
		res := serializer.BuildUsers(users)
		return serializer.BuildListResponse(res, uint(total))
	}

	return serializer.Response{
		Status: 40001,
		Msg:    "查找的类型非法",
		Error:  "search error",
	}
}
