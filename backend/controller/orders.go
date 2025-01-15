package controller


import (

   "net/http"

    "github.com/gin-gonic/gin"

   "backendproject/config"

   "backendproject/entity"

)


func GetAllOrders(c *gin.Context) {


   var orders []entity.Order


   db := config.DB()

   results := db.Find(& orders)

   if results.Error != nil {

       c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})

       return

   }

   c.JSON(http.StatusOK,  orders)


}


func GetOrders(c *gin.Context) {


   ID := c.Param("id")

   var order entity.Order


   db := config.DB()

   results := db.First(&order, ID)

   if results.Error != nil {

       c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})

       return

   }

   if order.ID == 0 {

       c.JSON(http.StatusNoContent, gin.H{})

       return

   }

   c.JSON(http.StatusOK, order)


}

func CreateOrders(c *gin.Context) {
    var order entity.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
        return
    }
    db := config.DB()
    result := db.Create(&order)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cart"})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "Created successful", "cart": order})
 }

func UpdateOrders(c *gin.Context) {


   var order entity.Order


   OrderID := c.Param("id")


   db := config.DB()

   result := db.First(&order, OrderID)

   if result.Error != nil {

       c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})

       return

   }


   if err := c.ShouldBindJSON(&order); err != nil {

       c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})

       return

   }


   result = db.Save(&order)

   if result.Error != nil {

       c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})

       return

   }


   c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})

}


func DeleteOrders(c *gin.Context) {


   id := c.Param("id")

   db := config.DB()

   if tx := db.Exec("DELETE FROM orders WHERE id = ?", id); tx.RowsAffected == 0 {

       c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})

       return

   }

   c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})

}