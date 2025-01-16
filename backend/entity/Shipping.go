package entity

import (
	"gorm.io/gorm"
)

type Shipping struct {
	gorm.Model
	Fee		int	
	Name	string
	
	ShippingStatusID	uint 
	ShippingStatus   	ShippingStatus 	`gorm:"foreignKey:ShippingStatusID"`

	
	
}