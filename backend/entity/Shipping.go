package entity

import (
	"gorm.io/gorm"
)

type Shipping struct {
	gorm.Model
	Name            string		`json:"name" valid:"required~Name is required, stringlength(1|100)"`
	Fee				uint		`json:"fee" valid:"required~Fee is required,range(1|1000)"`
	
	ShippingStatusID	uint 	`json:"ShippingStatusID" valid:"required~Shipping Status is required"`
	ShippingStatus   	ShippingStatus 	`gorm:"foreignKey:ShippingStatusID"`

	// Order []Order `gorm:"foreignKey:ShippingID"`
}