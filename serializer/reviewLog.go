package serializer

import "giligili/model"

// ReviewLog 审核日志序列化器
type ReviewLog struct {
	ID             uint  `json:"id"`
	VideoID        uint  `json:"video_id"`
	StatusForward  uint  `json:"status_forward"`
	StatusBackward uint  `json:"status_backward"`
	CreatedAt      int64 `json:"created_at"`
}

// BuildReviewLog 序列化审核日志
func BuildReviewLog(item model.ReviewLog) ReviewLog {
	return ReviewLog{
		ID:             item.ID,
		VideoID:        item.VideoID,
		StatusForward:  item.StatusForward,
		StatusBackward: item.StatusBackward,
		CreatedAt:      item.CreatedAt.Unix(),
	}
}

// BuildReviewLogs 序列化审核日志列表
func BuildReviewLogs(items []model.ReviewLog) (reviewLogs []ReviewLog) {
	for _, item := range items {
		reviewLog := BuildReviewLog(item)
		reviewLogs = append(reviewLogs, reviewLog)
	}
	return reviewLogs
}
