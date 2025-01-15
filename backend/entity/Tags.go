package entity

import (
	"gorm.io/gorm"
)

type Tags struct {
	gorm.Model
	Tag_name   string `json:"tag_name"`

	ProductTags []ProductTags `gorm:"foreignKey:TagsID"`
}