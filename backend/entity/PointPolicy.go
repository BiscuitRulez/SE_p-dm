package entity


import (

   "gorm.io/gorm"

)

type PointPolicy struct {
	gorm.Model
	ID          uint       `json:"id"`
	EarnRate    uint     `json:"earn"`
	RedeemRate uint     `json:"redeem"`
	Description        string     `json:"description"`     
}