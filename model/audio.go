package model

import (
	"giligili/cache"
	"os"
	"strconv"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/jinzhu/gorm"
)

// Audio 视频模型
type Audio struct {
	gorm.Model
	Title       string
	Info        string
	URL         string
	Avatar      string
	Type        string `gorm:"type:enum('education','food','technology')"`
	UserId      uint
	Status      string `gorm:"type:enum('pending_review','passed','not_passed'); default:'pending_review'"`
	FromOutside bool   `gorm:"default:'false'"`
}

//	通过ID获取视频
func GetAudioById(Id interface{}) (Audio, error) {
	var Audio Audio
	result := DB.First(&Audio, Id)
	return Audio, result.Error
}

//	创建审核日志事务
func ChangeAudioStatusBusiness(Audio Audio, reviewLog ReviewLog) error {
	tx := DB.Begin()

	if err := tx.Model(&Audio).Update("status", reviewLog.StatusBackward).Error; err != nil {
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
func (Audio *Audio) AvatarURL() string {
	if Audio.Avatar == "" {
		return ""
	}
	client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	signedGetURL, _ := bucket.SignURL(Audio.Avatar, oss.HTTPGet, 600)
	return signedGetURL
}

// AudioURL 视频地址
func (Audio *Audio) AudioURL() string {
	client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	signedGetURL, _ := bucket.SignURL(Audio.URL, oss.HTTPGet, 600)
	return signedGetURL
}

// View 点击数
func (Audio *Audio) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.AudioViewKey(Audio.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// AddView 视频游览
func (Audio *Audio) AddView() {
	// 增加视频点击数
	cache.RedisClient.Incr(cache.AudioViewKey(Audio.ID))
	// 增加排行点击数
	cache.RedisClient.ZIncrBy(cache.DailyRankKey, 1, strconv.Itoa(int(Audio.ID)))
}
