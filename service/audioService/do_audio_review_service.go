package audioService

import (
	"giligili/model"
	"giligili/serializer"
)

type DoAudioReviewService struct {
	StatusBackward string `form:"status_backward" json:"status_backward" `
	Remark         string `form:"remark" json:"remark" binding:"max=40"`
}

func (service *DoAudioReviewService) ChangeAudioStatus(reviewerId uint, audioId string) serializer.Response {
	audio, err := model.GetAudioById(audioId)
	if err != nil {
		return serializer.Response{
			Status: 40001,

			Msg:   "审核的音频不存在",
			Error: err.Error(),
		}
	}
	if audio.Status == service.StatusBackward {
		return serializer.Response{
			Status: 50002,
			Msg:    "音频已是该状态，不需更改",
			Error:  "status error",
		}
	}

	reviewLog := model.ReviewLog{
		MediaId:        audio.ID,
		MediaType:      "audio",
		ReviewerId:     reviewerId,
		StatusForward:  audio.Status,
		StatusBackward: service.StatusBackward,
		Remark:         service.Remark,
	}

	if err := model.ChangeAudioStatusBusiness(audio, reviewLog); err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "音频状态更改失败",
			Error:  err.Error(),
		}

	} else {
		return serializer.Response{
			Data: serializer.BuildReviewLog(reviewLog),
		}
	}
}
