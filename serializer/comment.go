package serializer

import "giligili/model"

// ReviewLog 审核日志序列化器
type Comment struct {
	ID        uint   `json:"id"`
	MediaType string `json:"media_type"`
	MediaId   uint   `json:"video_id"`
	Content   string `json:"content"`
	User      User   `json:"user"`
	CreatedAt int64  `json:"created_at"`
}

// BuildComment 序列化评论
func BuildComment(item model.Comment, userId uint) Comment {
	user, _ := model.GetUser(userId)
	return Comment{
		ID:        item.ID,
		MediaType: item.MediaType,
		MediaId:   item.MediaId,
		Content:   item.Content,
		User:      BuildUser(user),
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildComments 序列化评论列表
func BuildComments(items []model.Comment) (comments []Comment) {
	for _, item := range items {
		comment := BuildComment(item, item.FromUserId)
		comments = append(comments, comment)
	}
	return comments
}
