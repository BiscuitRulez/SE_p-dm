package entity

import (
	"time"

	"gorm.io/gorm"
)

type ReviewLike struct {
	gorm.Model
	Like_date 	time.Time 	`json:"like_date"`
	
	
	ReviewID		uint 	`json:"ReviewID"`
	Review   		Review 	`gorm:"foreignKey:ReviewID"`

	UserID	 	uint 	`json:"UserID"`
	User   		*Users 	`gorm:"foreignKey:UserID"`
}