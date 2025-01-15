package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"backendproject/config"
	"backendproject/entity"
)

// AddAddressController เพิ่มข้อมูลที่อยู่ใหม่
func AddAddressController(c *gin.Context) {
	var address entity.Address
	db := config.DB()

	// รับข้อมูล JSON และ Map กับ Struct Address
	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid JSON format",
			"details": err.Error(),
		})
		return
	}

	// ตรวจสอบว่าข้อมูลที่อยู่ถูกต้องหรือไม่
	if address.FullAddress == "" || address.City == "" || address.Province == "" || address.PostalCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "กรุณากรอกข้อมูลที่อยู่ให้ครบถ้วน",
		})
		return
	}

	// สร้างตัวแปร Address ใหม่สำหรับการบันทึก
	newAddress := entity.Address{
		FullAddress: address.FullAddress,
		City:        address.City,
		Province:    address.Province,
		PostalCode:  address.PostalCode,
		UserID:      address.UserID, // เชื่อมโยงกับ UserID ที่ได้รับมา
	}

	// บันทึกข้อมูล Address ในฐานข้อมูล
	if err := db.Create(&newAddress).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "ไม่สามารถเพิ่มที่อยู่ได้",
			"details": err.Error(),
		})
		return
	}

	// ส่งข้อมูลที่สร้างสำเร็จกลับไป
	c.JSON(http.StatusCreated, gin.H{
		"message": "เพิ่มที่อยู่สำเร็จ",
		"data":    newAddress,
	})
}


func GetAllAddress(c *gin.Context) {
    var address []entity.Address

    // Attempt to retrieve all airlines from the database
    if err := config.DB().Find(&address).Error; err != nil {
        // If there's an error, return a 500 status code with the error message
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // If successful, return the list of airlines with a 200 status code
    c.JSON(http.StatusOK, gin.H{"tags": address})
}


func GetAddressesByUserId(c *gin.Context) {
	var addresses []entity.Address // ใช้ slice เพื่อเก็บข้อมูลหลายรายการ
	id := c.Param("id")

	// ดึงข้อมูลทั้งหมดที่เกี่ยวข้องกับ user ID
	if err := config.DB().Where("user_id = ?", id).Find(&addresses).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Addresses not found"})
		return
	}

	// ส่งข้อมูลกลับไปในรูป JSON
	c.JSON(http.StatusOK, gin.H{"data": addresses})
}


func UpdateTagAddresss(c *gin.Context) {
	var address entity.Address
 
	AddressID := c.Param("id")
 
	db := config.DB()
 
	result := db.First(&address, AddressID)
 
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}
 
	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}
 
	result = db.Save(&address)
 
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
 
	}
	c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})
 
}

func DeleteAddress(c *gin.Context) {

	id := c.Param("id")
	db := config.DB()
	if tx := db.Exec("DELETE FROM Address WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})

}
