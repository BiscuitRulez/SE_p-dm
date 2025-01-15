package controller


import (

   "net/http"

    "github.com/gin-gonic/gin"

   "backendproject/config"

   "backendproject/entity"

)


func GetAllProductTags(c *gin.Context) {


   var producttags []entity.ProductTags


   db := config.DB()

   results := db.Find(&producttags)

   if results.Error != nil {

       c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})

       return

   }

   c.JSON(http.StatusOK, producttags)


}


func GetProductTags(c *gin.Context) {


   ID := c.Param("id")

   var producttag entity.ProductTags


   db := config.DB()

   results := db.First(&producttag, ID)

   if results.Error != nil {

       c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})

       return

   }

   if producttag.ID == 0 {

       c.JSON(http.StatusNoContent, gin.H{})

       return

   }

   c.JSON(http.StatusOK, producttag)


}

func CreateProductTags(c *gin.Context) {
    var producttag entity.ProductTags
    if err := c.ShouldBindJSON(&producttag); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
        return
    }
    db := config.DB()
    result := db.Create(&producttag)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cart"})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "Created successful", "cart": producttag})
 }

func UpdateProductTags(c *gin.Context) {


   var producttag entity.ProductTags


   ProductTagsID := c.Param("id")


   db := config.DB()

   result := db.First(&producttag, ProductTagsID)

   if result.Error != nil {

       c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})

       return

   }


   if err := c.ShouldBindJSON(&producttag); err != nil {

       c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})

       return

   }


   result = db.Save(&producttag)

   if result.Error != nil {

       c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})

       return

   }


   c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})

}


func DeleteProductTags(c *gin.Context) {


   id := c.Param("id")

   db := config.DB()

   if tx := db.Exec("DELETE FROM producttags WHERE id = ?", id); tx.RowsAffected == 0 {

       c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})

       return

   }

   c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})

}