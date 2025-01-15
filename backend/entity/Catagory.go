package entity


import (

   "gorm.io/gorm"

)

type Catagory struct {

   gorm.Model

   ID        uint     `json:"id"`

   Name string    `json:"name"`

}