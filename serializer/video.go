package serializer

import "giligili/model"

// Video 视频序列化器
type Video struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Info        string `json:"info"`
	URL         string `json:"url"`
	Avatar      string `json:"avatar"`
	View        uint64 `json:"view"`
	User        User   `json:"user"`
	CreatedAt   int64  `json:"created_at"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	FromOutside bool   `json:"from_outside"`
}

// BuildVideo 序列化视频
func BuildVideo(item model.Video) Video {
	user, _ := model.GetUser(item.UserId)
	videoURL := item.URL
	if !item.FromOutside {
		videoURL = item.VideoURL()
	}
	return Video{
		ID:          item.ID,
		Title:       item.Title,
		Info:        item.Info,
		URL:         videoURL,
		Avatar:      item.AvatarURL(),
		View:        item.View(),
		User:        BuildUser(user),
		CreatedAt:   item.CreatedAt.Unix(),
		Status:      item.Status,
		Type:        item.Type,
		FromOutside: item.FromOutside,
	}
}

// BuildVideos 序列化视频列表
func BuildVideos(items []model.Video) (videos []Video) {
	for _, item := range items {
		video := BuildVideo(item)
		videos = append(videos, video)
	}
	return videos
}
