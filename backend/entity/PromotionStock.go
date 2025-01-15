package entity

import (
	"gorm.io/gorm"
)

type PromotionStock struct {
	gorm.Model
	
	PromotionID		uint 	`json:"PromotionID"`
	Promotion   	Promotion 	`gorm:"foreignKey:PromotionID"`

	StockID	 	uint 	`json:"StockID"`
	Stock   	Stock 	`gorm:"foreignKey:StockID"`
}