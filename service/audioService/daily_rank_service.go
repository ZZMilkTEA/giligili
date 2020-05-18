package audioService

import (
	"fmt"
	"giligili/cache"
	"giligili/model"
	"giligili/serializer"
	"strings"
)

// DailyRankService 每日排行的服务
type DailyRankService struct {
}

// Get 获取排行
func (service *DailyRankService) Get() serializer.Response {
	var audios []model.Audio

	// 从redis读取点击前十的音频
	vids, _ := cache.RedisClient.ZRevRange(cache.DailyRankKey, 0, 9).Result()

	if len(vids) > 1 {
		order := fmt.Sprintf("FIELD(id, %s)", strings.Join(vids, ","))
		err := model.DB.Where("id in (?)", vids).Order(order).Find(&audios).Error
		if err != nil {
			return serializer.Response{
				Status: 50000,
				Msg:    "数据库连接错误",
				Error:  err.Error(),
			}
		}
	}

	return serializer.Response{
		Data: serializer.BuildAudios(audios),
	}
}
