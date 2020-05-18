package model

import "github.com/jinzhu/gorm"

type Notice struct {
	gorm.Model
	title   string
	content string
}
