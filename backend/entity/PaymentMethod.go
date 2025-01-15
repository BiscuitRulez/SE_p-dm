package entity

import (
	"gorm.io/gorm"
)

type PaymentMethod struct {
	gorm.Model
	PaymentMethod  	string 		`valid:"required~PaymentStatus is required"` 
	
	Payment []Payment `gorm:"foreignKey:PaymentMethodID"`
}

