package controller


import (

   "net/http"


   "github.com/gin-gonic/gin"


   "backendproject/config"

   "backendproject/entity"

)


func GetAll(c *gin.Context) {


   var stocks []entity.Stock


   db := config.DB()

   results := db.Find(&stocks)

   if results.Error != nil {

       c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})

       return

   }

   c.JSON(http.StatusOK, stocks)


}


func Get(c *gin.Context) {


   ID := c.Param("id")

   var stock entity.Stock


   db := config.DB()

   results := db.First(&stock, ID)

   if results.Error != nil {

       c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})

       return

   }

   if stock.ID == 0 {

       c.JSON(http.StatusNoContent, gin.H{})

       return

   }

   c.JSON(http.StatusOK, stock)


}


func Update(c *gin.Context) {


   var stock entity.Stock


   StockID := c.Param("id")


   db := config.DB()

   result := db.First(&stock, StockID)

   if result.Error != nil {

       c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})

       return

   }


   if err := c.ShouldBindJSON(&stock); err != nil {

       c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})

       return

   }


   result = db.Save(&stock)

   if result.Error != nil {

       c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})

       return

   }


   c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})

}


func Delete(c *gin.Context) {


   id := c.Param("id")

   db := config.DB()

   if tx := db.Exec("DELETE FROM stock WHERE id = ?", id); tx.RowsAffected == 0 {

       c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})

       return

   }

   c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})

}

func Create(c *gin.Context) {
    var stock entity.Stock
    if err := c.ShouldBindJSON(&stock); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
        return
    }
    db := config.DB()
    result := db.Create(&stock)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create stock"})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "Created successful", "stock": stock})
 }