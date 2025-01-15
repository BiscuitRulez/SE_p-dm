package entity

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	Rating		uint	
	
	ProductID	uint 	
	Product   	Product 	

	ReviewLike []ReviewLike 
}