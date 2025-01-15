package entity

import (
	"gorm.io/gorm"
)

type ClaimStatus struct {
	gorm.Model
	ClaimStatus string `valid:"required~ClaimStatus is required"`

	Claim []Claim `gorm:"foreignKey:ClaimStatusID"`
}
