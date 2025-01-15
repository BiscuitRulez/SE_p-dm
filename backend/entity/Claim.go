package entity

import (
	"time"
	"gorm.io/gorm"
)

type Claim struct {
	gorm.Model
	Date      			time.Time 	`valid:"required~Date is required"`
	Photo     			string    	`valid:"required~Photo is required"`

	ProblemID 			uint      	`valid:"required~ProblemID is required"`
	Problem   			Problem   	`gorm:"foreignKey:ProblemID"`

	ClaimStatusID 		uint      	`valid:"required~ClaimStatusID is required"`
	ClaimStatus 		ClaimStatus `gorm:"foreignKey:ClaimStatusID"`

	UserID    			uint      	`valid:"required~UserID is required"`
	User     		 	Users     	`gorm:"foreignKey:UserID"`

	OrderID   			uint      	`valid:"required~OrderID is required"`
	Order     			Order     	`gorm:"foreignKey:OrderID"`
}
