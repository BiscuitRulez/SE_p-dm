package entity

import (
	"gorm.io/gorm"
	"time"
)

type Promotion struct {
	gorm.Model
	Topic  			string 		
	Date_start   	time.Time 
	Date_end   		time.Time 	
	Status  		string 		

	PromotionStock []PromotionStock `gorm:"foreignKey:PromotionID"`
}