package entity

import (
	"gorm.io/gorm"
)

type PointPolicy struct {
	gorm.Model
	Earn_rate  		float32 	
	Redeem_rate   	float32 		
	Description   	string 		
}