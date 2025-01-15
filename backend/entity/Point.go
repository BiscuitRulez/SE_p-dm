
package entity

import "gorm.io/gorm"

type Point struct {
    gorm.Model
    ID         uint `json:"id"`
    TotalPoint uint `json:"total" validate:"gt=0"` // ต้องมากกว่า 0
}
