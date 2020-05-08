package model

import (
	"github.com/jinzhu/gorm"
)

type ReviewLog struct {
	gorm.Model
	VideoID        uint
	ReviewerId     uint
	StatusForward  uint
	StatusBackward uint
}
