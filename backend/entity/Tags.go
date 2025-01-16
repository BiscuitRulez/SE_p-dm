package entity

import (
	"gorm.io/gorm"
)

type Tags struct {
	gorm.Model
	Tag_Name   string `json:"Tag_Name"`

	ProductTags []ProductTags `gorm:"foreignKey:TagsID"`
}