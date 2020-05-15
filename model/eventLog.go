package model

import (
	"github.com/jinzhu/gorm"
)

type EventLog struct {
	gorm.Model
	ExecutiveId  uint   `gorm:"not null"`
	ExecutedType string `gorm:"type:enum('create','delete','modify');not null"`
	ExecutedId   uint   `gorm:"not null"`
	Action       string `gorm:"not null"`
	Remark       string
}

const (
	Create string = "create"
	Delete string = "delete"
	Modify string = "modify"
)
