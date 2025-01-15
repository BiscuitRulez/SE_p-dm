package entity

import (
	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	HistoryChat   string `json:"HistoryChat"`

	UserAdminChat []UserAdminChat `gorm:"foreignKey:ChatID"`
}