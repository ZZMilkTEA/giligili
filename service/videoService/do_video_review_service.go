package videoService

import (
	"giligili/model"
	"giligili/serializer"
)

type DoVideoReviewService struct {
	VideoID        string `form:"video_id" json:"video_id,string" `
	StatusBackward string `form:"status_backward" json:"status_backward,string" `
}

func (service *DoVideoReviewService) ChangeVideoStatus(reviewerID uint) serializer.Response {
	video, err := model.GetVideoById(service.VideoID)
	if err != nil {
		return serializer.Response{
			Status: 40001,

			Msg:   "审核的视频不存在",
			Error: err.Error(),
		}
	}
	if video.Status == service.StatusBackward {
		return serializer.Response{
			Status: 50002,
			Msg:    "视频已是该状态，不需更改",
			Error:  "status error",
		}
	}

	reviewLog := model.ReviewLog{
		VideoId:        video.ID,
		ReviewerId:     reviewerID,
		StatusForward:  video.Status,
		StatusBackward: service.StatusBackward,
	}

	if err := model.ChangeVideoStatusBusiness(video, reviewLog); err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "视频状态更改失败",
			Error:  err.Error(),
		}
	} else {
		return serializer.Response{
			Data: serializer.BuildReviewLog(reviewLog),
		}
	}
}
