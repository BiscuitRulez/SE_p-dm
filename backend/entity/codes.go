package entity

import (
	"time"

	"gorm.io/gorm"
)

type Codes struct {

	gorm.Model

	CodeTopic		string		`json:"code_topic" valid:"required~CodeTopic is required"`

	CodeDescription	string		`json:"code_description" valid:"required~CodeDescription is required"`

	Discount		int			`json:"discount" valid:"required~Discount is required"`

	Quantity		int			`json:"quantity" valid:"required~Quantity is required"`

	DateStart		time.Time	`json:"date_start" valid:"required~DateStart is required"`

	DateEnd			time.Time	`json:"date_end" valid:"required~DateEnd is required"`

	CodeStatus		string		`json:"code_status" valid:"required~CodeStatus is required"`

	CodePicture		string		`json:"code_picture" valid:"required~CodePicture is required"`

	Minimum 		int			`json:"minimum" valid:"required~Minimum is required"`


	CodeCollector []CodeCollectors `gorm:"foreignKey:code_id"`
}

 