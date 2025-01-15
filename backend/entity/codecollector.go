package entity

import (

	"gorm.io/gorm"
)

type CodeCollectors struct {

	gorm.Model

	UserID		uint	`json:"user_id"`

	User		*Users	`gorm:"foreignKey:user_id" json:"user"`

	CodeID		uint	`json:"code_id"`

	Code		*Codes	`gorm:"foreignKey:code_id" json:"code"`

	Order []Order `gorm:"foreignKey:CodeCollectID"`

}