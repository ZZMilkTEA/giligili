package model

import (
	"giligili/cache"
	"os"
	"strconv"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/jinzhu/gorm"
)

// Video 视频模型
type Video struct {
	gorm.Model
	Title  string
	Info   string
	URL    string
	Avatar string
	Type   uint //视频分类（分区），暂时用uint类型，0 教育， 1 美食， 2 科技
	UserID uint
	Status uint //0 待审核，1 审核通过（发布），2 审核不通过
}

//	通过ID获取视频
func GetVideoById(Id interface{}) (Video, error) {
	var video Video
	result := DB.First(&video, Id)
	return video, result.Error
}

//	创建审核日志事务
func ChangeVideoStatusBusiness(video Video, reviewLog ReviewLog) error {
	tx := DB.Begin()

	if err := tx.Model(&video).Update("status", reviewLog.StatusBackward).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&reviewLog).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// AvatarURL 封面地址
func (video *Video) AvatarURL() string {
	client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	signedGetURL, _ := bucket.SignURL(video.Avatar, oss.HTTPGet, 600)
	return signedGetURL
}

// VideoURL 视频地址
func (video *Video) VideoURL() string {
	client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	signedGetURL, _ := bucket.SignURL(video.URL, oss.HTTPGet, 600)
	return signedGetURL
}

// View 点击数
func (video *Video) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.VideoViewKey(video.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// AddView 视频游览
func (video *Video) AddView() {
	// 增加视频点击数
	cache.RedisClient.Incr(cache.VideoViewKey(video.ID))
	// 增加排行点击数
	cache.RedisClient.ZIncrBy(cache.DailyRankKey, 1, strconv.Itoa(int(video.ID)))
}
