package entity

import (
	"gorm.io/gorm"
)

type ProductTags struct {
	gorm.Model

	ProductID	uint 		`json:"ProductID"`
	Product   	Product 	`gorm:"foreignKey:ProductID"`

	TagsID	 	uint 	`json:"TagsID"`
	Tags   		Tags 	`gorm:"foreignKey:TagsID"`
}