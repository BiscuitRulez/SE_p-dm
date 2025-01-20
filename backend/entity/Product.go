package entity

import (
	"fmt"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
    ID          uint       `json:"id"`
	Name string `json:"name" valid:"required~name is required"`
	// Modified validation message to match test expectation

	Description string `json:"description" valid:"required~description is required"`
	Image       string `json:"image" valid:"required~image_url is required,url~image must be a valid URL"`
	Stock       *Stock      `json:"stock,omitempty" gorm:"foreignKey:ProductID"`
	CartItems  []CartItem  `json:"-" gorm:"foreignKey:ProductID"`
	OrderItems []OrderItem `json:"-" gorm:"foreignKey:ProductID"`
	Reviews    []Review    `json:"reviews" gorm:"foreignKey:ProductID"`
	AvgRating  float64     `json:"avg_rating" gorm:"default:0" valid:"range(0|5)~average rating must be between 0 and 5"`
	UserID     uint        `json:"user_id"`
	User       *Users      `gorm:"foreignKey:UserID" json:"user"` // แก้ foreignKey
	CatagoryID uint        `json:"catagory_id"`
	Catagory   *Catagory   `gorm:"foreignKey:CatagoryID" json:"catagory"` // แก้ foreignKey
	ProductTags []ProductTags `gorm:"foreignKey:ProductID"`
}

// Validate validates the Product struct using custom validation
func (p *Product) Validate() error {
	if p.Name == "" {
		return fmt.Errorf("Name is required")
	}
	// if p.Price < 0 {
	// 	return fmt.Errorf("price must be positive")
	// }
	return nil
}
