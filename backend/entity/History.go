package entity


import (

   "time"

   "gorm.io/gorm"

)

type History struct {
   gorm.Model
   ID              uint            
   OrderDate       time.Time         
   PointsEarned    float32          
   PointsRedeemed  float32          
   TotalAmount     float32   

   UserID          uint          
   User            Users           `gorm:"foreignKey:UserID"`

   OrderID         uint             
   Order           Order           `gorm:"foreignKey:OrderID"`

   HistoryDetails  []HistoryDetail   
}