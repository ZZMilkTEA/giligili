package model

import "github.com/jinzhu/gorm"

type ReviewLog struct {
	gorm.Model
	Contributor    uint
	VideoID        uint
	ReviewerId     uint
	statusForward  string
	statusBackward string
}
