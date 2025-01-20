package entity

import (
	"gorm.io/gorm"
	"time"
)

type Payment struct {
	gorm.Model
	
	Date				time.Time		`valid:"required~Date is required"`

	UserID				uint 			`valid:"required~UserID is required"`
	User   				*Users 			`gorm:"foreignKey:UserID"`
	

	PaymentMethodID		uint 			`valid:"required~PaymentMethodID is required"`
	PaymentMethod   	PaymentMethod 	`gorm:"foreignKey:PaymentMethodID"`

	PaymentStatusID		uint 			`valid:"required~PaymentStatusID is required"`
	PaymentStatus   	PaymentStatus	`gorm:"foreignKey:PaymentStatusID"`
	AddressID			uint 			`valid:"required~AddressID is required"`
	Address   			Address			`gorm:"foreignKey:AddressID"`
}


