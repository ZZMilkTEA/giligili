package model

import (
	"github.com/jinzhu/gorm"
)

type ReviewLog struct {
	gorm.Model
	VideoId        uint   `gorm:"not null"`
	ReviewerId     uint   `gorm:"not null"`
	StatusForward  string `gorm:"type:enum('pending_review','passed','not_passed'); default:'pending_review'"`
	StatusBackward string `gorm:"type:enum('pending_review','passed','not_passed'); default:'pending_review'"`
	Remark         string
}
