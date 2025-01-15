package controller


import (

   "net/http"

    "github.com/gin-gonic/gin"

   "backendproject/config"

   "backendproject/entity"

)


func GetAllReviews(c *gin.Context) {


   var reviews []entity.Review


   db := config.DB()

   results := db.Find(&reviews)

   if results.Error != nil {

       c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})

       return

   }

   c.JSON(http.StatusOK, reviews)


}


func GetReviews(c *gin.Context) {


   ID := c.Param("id")

   var review entity.Review


   db := config.DB()

   results := db.First(&review, ID)

   if results.Error != nil {

       c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})

       return

   }

   if review.ID == 0 {

       c.JSON(http.StatusNoContent, gin.H{})

       return

   }

   c.JSON(http.StatusOK, review)


}

func CreateReviews(c *gin.Context) {
    var review entity.Review
    if err := c.ShouldBindJSON(&review); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
        return
    }
    db := config.DB()
    result := db.Create(&review)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cart"})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "Created successful", "cart": review})
 }

func UpdateReviews(c *gin.Context) {


   var review entity.Review


   ReviewID := c.Param("id")


   db := config.DB()

   result := db.First(&review, ReviewID)

   if result.Error != nil {

       c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})

       return

   }


   if err := c.ShouldBindJSON(&review); err != nil {

       c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})

       return

   }


   result = db.Save(&review)

   if result.Error != nil {

       c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})

       return

   }


   c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})

}


func DeleteReviews(c *gin.Context) {


   id := c.Param("id")

   db := config.DB()

   if tx := db.Exec("DELETE FROM reviews WHERE id = ?", id); tx.RowsAffected == 0 {

       c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})

       return

   }

   c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})

}