package entity

import (
	"gorm.io/gorm"
)

type Address struct {
	
	gorm.Model

	FullAddress 	string 		`json:"full_address"`
	City  			string 		`json:"city"`
	Province 		string 		`json:"province"`
	PostalCode  	string 		`json:"postal_code"`

	UserID			uint		`json:"user_id"`
	User			*Users		`gorm:"foreignKey:user_id" json:"user"`

	Payment 		[]Payment 	`gorm:"foreignKey:AddressID"`
}