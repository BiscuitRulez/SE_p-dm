package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"backendproject/config"
	"backendproject/entity"
)

func GetAllShippingStatus(c *gin.Context) {
    var shippingstatus []entity.ShippingStatus

    // Attempt to retrieve all airlines from the database
    if err := config.DB().Find(&shippingstatus).Error; err != nil {
        // If there's an error, return a 500 status code with the error message
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // If successful, return the list of airlines with a 200 status code
    c.JSON(http.StatusOK, gin.H{"shippingstatuss": shippingstatus})
}


func GetShippingStatusByID(c *gin.Context) {
	var shippingstatus entity.ShippingStatus
	id := c.Param("id")

	// ดึงข้อมูลจากฐานข้อมูลตาม ID
	if err := config.DB().First(&shippingstatus, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ShippingStatus not found"})
		return
	}

	// ส่งข้อมูลกลับไปในรูป JSON
	c.JSON(http.StatusOK, gin.H{"data": shippingstatus})
}

func UpdateShippingStatus(c *gin.Context) {
	var shippingstatus entity.ShippingStatus
 
	ShippingStatusID := c.Param("id")
 
	db := config.DB()
 
	result := db.First(&shippingstatus, ShippingStatusID)
 
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}
 
	if err := c.ShouldBindJSON(&shippingstatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}
 
	result = db.Save(&shippingstatus)
 
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
 
	}
	c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})
 
}

func DeleteShippingStatus(c *gin.Context) {

	id := c.Param("id")
	db := config.DB()
	if tx := db.Exec("DELETE FROM ShippingStatus WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})

}