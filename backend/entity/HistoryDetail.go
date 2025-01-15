package entity


import (

   "gorm.io/gorm"

)

type HistoryDetail struct {

   gorm.Model

   ID                uint    

   ProductName       string    

   Quantity          uint   

   PricePerUnit      float32   

   SubTotal  float32   

   StockID  uint   
   Stock    Stock  

   HistoryID  uint      
   History    History  
}