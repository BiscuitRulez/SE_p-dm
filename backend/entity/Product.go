package entity


import (

   "gorm.io/gorm"

)

type Product struct {
	gorm.Model
	ID          uint       `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Quantity	string     `json:"quantity"`
	Image       string     `json:"image"`
	UserID      uint       `json:"user_id"`
	User        *Users     `gorm:"foreignKey:UserID" json:"user"`         // แก้ foreignKey
	CatagoryID  uint       `json:"catagory_id"`
	Catagory    *Catagory  `gorm:"foreignKey:CatagoryID" json:"catagory"` // แก้ foreignKey
}