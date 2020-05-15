package serializer

import "giligili/model"

// ReviewLog 审核日志序列化器
type ReviewLog struct {
	ID             uint   `json:"id"`
	VideoId        uint   `json:"video_id"`
	StatusForward  string `json:"status_forward"`
	StatusBackward string `json:"status_backward"`
	CreatedAt      int64  `json:"created_at"`
	Remark         string `json:"remark"`
}

// BuildReviewLog 序列化审核日志
func BuildReviewLog(item model.ReviewLog) ReviewLog {
	return ReviewLog{
		ID:             item.ID,
		VideoId:        item.VideoId,
		StatusForward:  item.StatusForward,
		StatusBackward: item.StatusBackward,
		CreatedAt:      item.CreatedAt.Unix(),
		Remark:         item.Remark,
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
