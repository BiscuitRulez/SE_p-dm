package entity


import (

   "gorm.io/gorm"

)

type OrderItem struct {

   gorm.Model

   ID        uint     `json:"id"`

   Quantity  uint    `json:"quantity"`

   Price     float32    `json:"price"`

   OrderID  uint      

   Order    Order  

   StockID  uint     

   Stock    Stock  
}