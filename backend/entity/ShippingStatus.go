package entity

import (
	"gorm.io/gorm"
	"time"
)

type ShippingStatus struct {
	gorm.Model
	Status   string `json:"status"`
	ShippingDate	time.Time	`json:"Shipping"`

	Shipping []Shipping `gorm:"foreignKey:ShippingStatusID"`
}