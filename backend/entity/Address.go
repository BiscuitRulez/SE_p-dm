package entity

import (
	"gorm.io/gorm"
)

type Address struct {
	
	gorm.Model

	FullAddress 	string 		`json:"full_address" valid:"required~Full Address is required, stringlength(1|100)"`
	City  			string 		`json:"city" valid:"required~City is required, stringlength(1|50)"`
	Province 		string 		`json:"province" valid:"required~Province is required, stringlength(1|50)"`
	PostalCode  	string 		`json:"postal_code" valid:"required~Postal Code is required,range(1|99999)"`

	UserID			uint		`json:"user_id" valid:"required~User ID is required"`
	User			*Users		`gorm:"foreignKey:user_id" json:"user" valid:"-"`
}