package serializer

import "giligili/model"

// audio 音频序列化器
type audio struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Info        string `json:"info"`
	URL         string `json:"url"`
	Avatar      string `json:"avatar"`
	View        uint64 `json:"view"`
	User        User   `json:"user"`
	CreatedAt   int64  `json:"created_at"`
	Status      string `json:"status"`
	Kind        string `json:"kind"`
	FromOutside bool   `json:"from_outside"`
}

// BuildAudio 序列化音频
func BuildAudio(item model.Audio) audio {
	user, _ := model.GetUser(item.UserId)
	audioURL := item.URL
	if !item.FromOutside {
		audioURL = item.AudioURL()
	}
	return audio{
		ID:          item.ID,
		Title:       item.Title,
		Info:        item.Info,
		URL:         audioURL,
		Avatar:      item.AvatarURL(),
		View:        item.View(),
		User:        BuildUser(user),
		CreatedAt:   item.CreatedAt.Unix(),
		Status:      item.Status,
		Kind:        item.Kind,
		FromOutside: item.FromOutside,
	}
}

// BuildAudios 序列化音频列表
func BuildAudios(items []model.Audio) (audios []audio) {
	for _, item := range items {
		audio := BuildAudio(item)
		audios = append(audios, audio)
	}
	return audios
}
