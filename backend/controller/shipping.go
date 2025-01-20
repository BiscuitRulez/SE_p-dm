package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"backendproject/config"
	"backendproject/entity"
)

func GetAllShipping(c *gin.Context) {
    var shipping []entity.Shipping

    // Attempt to retrieve all airlines from the database
    if err := config.DB().Find(&shipping).Error; err != nil {
        // If there's an error, return a 500 status code with the error message
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // If successful, return the list of airlines with a 200 status code
    c.JSON(http.StatusOK, gin.H{"shippings": shipping})
}

func GetShippingByID(c *gin.Context) {
	var shipping entity.Shipping
	id := c.Param("id")

	// ดึงข้อมูลจากฐานข้อมูลตาม ID
	if err := config.DB().First(&shipping, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Shipping not found"})
		return
	}

	// ส่งข้อมูลกลับไปในรูป JSON
	c.JSON(http.StatusOK, gin.H{"data": shipping})
}

func CreateShipping(c *gin.Context) {

    var newShipping entity.Shipping // สร้างตัวแปรสำหรับเก็บข้อมูล Code ใหม่

    // ผูก JSON ที่ส่งมาจาก Request Body กับตัวแปร newCode
    if err := c.ShouldBindJSON(&newShipping); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // เชื่อมต่อกับฐานข้อมูล
    db := config.DB()

	ns := entity.Shipping{

		Name:	newShipping.Name,
		Fee:	newShipping.Fee,

	}

    // บันทึกข้อมูลลงในฐานข้อมูล
    if err := db.Create(&ns).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // ส่งข้อมูลที่สร้างสำเร็จกลับไป
    c.JSON(http.StatusCreated, gin.H {"message": "Shipping created successfully","shipping": ns})
}

func UpdateShippingByID(c *gin.Context) {
	var shipping entity.Shipping
 
	ShippingID := c.Param("id")
 
	db := config.DB()
 
	result := db.First(&shipping, ShippingID)
 
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}
 
	if err := c.ShouldBindJSON(&shipping); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}
 
	result = db.Save(&shipping)
 
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
 
	}
	c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})
 
}

func DeleteShipping(c *gin.Context) {

	id := c.Param("id")
	db := config.DB()
	if tx := db.Exec("DELETE FROM Shipping WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})

}