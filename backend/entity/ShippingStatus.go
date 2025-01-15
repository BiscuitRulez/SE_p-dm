package entity

import (
	"gorm.io/gorm"
)

type ShippingStatus struct {
	gorm.Model
	Status   string `json:"status"`

	Shipping []Shipping `gorm:"foreignKey:ShippingStatusID"`
}