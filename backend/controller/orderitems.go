package controller


import (

   "net/http"

    "github.com/gin-gonic/gin"

   "backendproject/config"

   "backendproject/entity"

)


func GetAllOrderItems(c *gin.Context) {


   var orderitems []entity.OrderItem


   db := config.DB()

   results := db.Find(& orderitems)

   if results.Error != nil {

       c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})

       return

   }

   c.JSON(http.StatusOK,  orderitems)


}


func GetOrderItems(c *gin.Context) {


   ID := c.Param("id")

   var orderitem entity.OrderItem


   db := config.DB()

   results := db.First(&orderitem, ID)

   if results.Error != nil {

       c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})

       return

   }

   if orderitem.ID == 0 {

       c.JSON(http.StatusNoContent, gin.H{})

       return

   }

   c.JSON(http.StatusOK, orderitem)


}

func CreateOrderItems(c *gin.Context) {
    var orderitem entity.OrderItem
    if err := c.ShouldBindJSON(&orderitem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
        return
    }
    db := config.DB()
    result := db.Create(&orderitem)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cart"})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "Created successful", "cart": orderitem})
 }

func UpdateOrderItems(c *gin.Context) {


   var orderitem entity.OrderItem


   OrderItemID := c.Param("id")


   db := config.DB()

   result := db.First(&orderitem, OrderItemID)

   if result.Error != nil {

       c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})

       return

   }


   if err := c.ShouldBindJSON(&orderitem); err != nil {

       c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})

       return

   }


   result = db.Save(&orderitem)

   if result.Error != nil {

       c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})

       return

   }


   c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})

}


func DeleteOrderItems(c *gin.Context) {


   id := c.Param("id")

   db := config.DB()

   if tx := db.Exec("DELETE FROM orderitems WHERE id = ?", id); tx.RowsAffected == 0 {

       c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})

       return

   }

   c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})

}