package entity


import (

   "gorm.io/gorm"

)

type Stock struct {
	gorm.Model
	ID               uint     `json:"id"`
	Price            float32  `json:"price"`
	Quantity  uint     `json:"quantity_in_stock"`
	Color            string   `json:"color"`
	ShapeSize        string   `json:"shape_size"`
	Image            string   `json:"image"`
	ProductID        uint     `json:"product_id"`
	Product          *Product `gorm:"foreignKey:ProductID" json:"product"` // แก้ foreignKey
}