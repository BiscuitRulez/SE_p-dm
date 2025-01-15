package controller


import (

   "net/http"

    "github.com/gin-gonic/gin"

   "backendproject/config"

   "backendproject/entity"

)


func GetAllReviewLikes(c *gin.Context) {


   var reviewlikes []entity.ReviewLike


   db := config.DB()

   results := db.Find(&reviewlikes)

   if results.Error != nil {

       c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})

       return

   }

   c.JSON(http.StatusOK, reviewlikes)


}


func GetReviewLikes(c *gin.Context) {


   ID := c.Param("id")

   var reviewlike entity.ReviewLike


   db := config.DB()

   results := db.First(&reviewlike, ID)

   if results.Error != nil {

       c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})

       return

   }

   if reviewlike.ID == 0 {

       c.JSON(http.StatusNoContent, gin.H{})

       return

   }

   c.JSON(http.StatusOK, reviewlike)


}

func CreateReviewLikes(c *gin.Context) {
    var reviewlike entity.ReviewLike
    if err := c.ShouldBindJSON(&reviewlike); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
        return
    }
    db := config.DB()
    result := db.Create(&reviewlike)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cart"})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "Created successful", "cart": reviewlike})
 }

func UpdateReviewLikes(c *gin.Context) {


   var reviewlike entity.ReviewLike


   ReviewLikeID := c.Param("id")


   db := config.DB()

   result := db.First(&reviewlike, ReviewLikeID)

   if result.Error != nil {

       c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})

       return

   }


   if err := c.ShouldBindJSON(&reviewlike); err != nil {

       c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})

       return

   }


   result = db.Save(&reviewlike)

   if result.Error != nil {

       c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})

       return

   }


   c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})

}


func DeleteReviewLikes(c *gin.Context) {


   id := c.Param("id")

   db := config.DB()

   if tx := db.Exec("DELETE FROM reviewlikes WHERE id = ?", id); tx.RowsAffected == 0 {

       c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})

       return

   }

   c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})

}