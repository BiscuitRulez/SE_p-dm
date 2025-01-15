package entity

import (
	"gorm.io/gorm"
)

type PaymentStatus struct {
	gorm.Model
	PaymentStatus string `valid:"required~PaymentStatus is required"`

	Payment []Payment `gorm:"foreignKey:PaymentStatusID"`
}
