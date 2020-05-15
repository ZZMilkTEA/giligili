package model

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	MediaType  string `gorm:"type:enum('video','audio');not null"`
	MediaId    uint   `gorm:"not null"`
	Content    string `gorm:"size:1024"`
	FromUserId uint   `gorm:"not null"`
}
