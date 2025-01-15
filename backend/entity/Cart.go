package entity

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Quantity 	uint 		`json:"quantity"`

	
	UserID		uint 	`json:"UserID"`
	User   		*Users 	`gorm:"foreignKey:UserID"`

	StockID	 	uint 	`json:"StockID"`
	Stock   	Stock 	`gorm:"foreignKey:StockID"`
}