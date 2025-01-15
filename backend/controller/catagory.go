package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"backendproject/config"
	"backendproject/entity"
)

func GetAllCatagory(c *gin.Context) {
    var catagory []entity.Catagory

    // Attempt to retrieve all airlines from the database
    if err := config.DB().Find(&catagory).Error; err != nil {
        // If there's an error, return a 500 status code with the error message
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // If successful, return the list of airlines with a 200 status code
    c.JSON(http.StatusOK, gin.H{"catagorys": catagory})
}


func GetCatagoryByID(c *gin.Context) {
	var catagory entity.Catagory
	id := c.Param("id")

	// ดึงข้อมูลจากฐานข้อมูลตาม ID
	if err := config.DB().First(&catagory, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Catagory not found"})
		return
	}

	// ส่งข้อมูลกลับไปในรูป JSON
	c.JSON(http.StatusOK, gin.H{"data": catagory})
}

func UpdateCatagory(c *gin.Context) {
	var catagory entity.Catagory
 
	CatagoryID := c.Param("id")
 
	db := config.DB()
 
	result := db.First(&catagory, CatagoryID)
 
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}
 
	if err := c.ShouldBindJSON(&catagory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}
 
	result = db.Save(&catagory)
 
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
 
	}
	c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})
 
}

func DeleteCatagory(c *gin.Context) {

	id := c.Param("id")
	db := config.DB()
	if tx := db.Exec("DELETE FROM Catagory WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})

}
