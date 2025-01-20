package entity

import (
	"gorm.io/gorm"
)

type Catagory struct {
	gorm.Model
	Name   string `json:"Name"`

	Product []Product `gorm:"foreignKey:CatagoryID"`
}