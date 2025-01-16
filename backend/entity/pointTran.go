package entity


import (

	"time"
    "gorm.io/gorm"

)

type PointTran struct {
	gorm.Model
	ID          uint       `json:"id"`
	PointsEarned    uint     `json:"point_earned"`
	PointsRedeemed uint     `json:"point_redeemed"`
	TransactionType       string     `json:"type"`
	TransactionDate      time.Time       `json:"date"`
	Description        string     `json:"description"`         // แก้ foreignKey
	OrderID uint   `json:"order_id"`
    Order   *Order `gorm:"foreignKey:OrderID" json:"order"`
	PointID uint   `json:"point_id"`
    Point   *Point `gorm:"foreignKey:PointID" json:"point"`
}