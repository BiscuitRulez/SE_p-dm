package entity

import (
	"gorm.io/gorm"
)

type Problem struct {
	gorm.Model
	Problem string `valid:"required~Problem is required"`

	Claim []Claim `gorm:"foreignKey:ProblemID"`
}
