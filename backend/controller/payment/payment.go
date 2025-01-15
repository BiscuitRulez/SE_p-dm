package controller

import (
	"net/http"
	"strconv"

	// "fmt"
	"backendproject/config"
	"backendproject/entity"
	"time"

	"github.com/gin-gonic/gin"
)


///////////////////////
func GetPayments(c *gin.Context) {
	// สร้าง slice สำหรับเก็บผลลัพธ์
	var payments []entity.Payment

	// ใช้การเชื่อมต่อกับฐานข้อมูล
	db := config.DB()

	// ดึงข้อมูลทั้งหมดจากตาราง payments
	result := db.Find(&payments)

	// ตรวจสอบข้อผิดพลาด
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// ส่งผลลัพธ์กลับในรูปแบบ JSON
	c.JSON(http.StatusOK, payments)
}

func GetPaymentByID(c *gin.Context) {


	ID := c.Param("id")
 
	var payment entity.Payment
 
 
	db := config.DB()
 
	results := db.First(&payment, ID)
 
	if results.Error != nil {
 
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
 
		return
 
	}
 
	if payment.ID == 0 {
 
		c.JSON(http.StatusNoContent, gin.H{})
 
		return
 
	}
 
	c.JSON(http.StatusOK, payment)
 
 
 }

func CreatePayment(c *gin.Context) {
	// ดึง UserID จาก URL และแปลงเป็น integer
	UserID := c.Param("id")
	userIDInt, err := strconv.Atoi(UserID) // แปลง UserID จาก string เป็น int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	// โครงสร้างสำหรับรับ JSON
	var request struct {
		PaymentMethodID uint 
		PaymentStatusID uint 
	}

	// Bind JSON
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ตรวจสอบว่า PaymentMethodID และ PaymentStatusID มีอยู่ในฐานข้อมูล
	db := config.DB().Debug()
	var paymentMethod entity.PaymentMethod
	var paymentStatus entity.PaymentStatus

	if err := db.First(&paymentMethod, request.PaymentMethodID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Payment Method"})
		return
	}

	if err := db.First(&paymentStatus, request.PaymentStatusID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Payment Status"})
		return
	}

	// สร้าง entity Payment
	newPayment := entity.Payment{
		PaymentMethodID: request.PaymentMethodID,
		PaymentStatusID: request.PaymentStatusID,
		Date:            time.Now(),
		UserID:          uint(userIDInt),
	}

	// บันทึกข้อมูลลง DB
	if err := db.Create(&newPayment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ส่งข้อมูลกลับ
	c.JSON(http.StatusCreated, gin.H{"message": "Created success", "data": newPayment})
}

func DeletePayment(c *gin.Context) {

	id := c.Param("id")
	db := config.DB()
	if tx := db.Exec("DELETE FROM payments WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})

}

func UpdatePaymentByID(c *gin.Context) {
	var payment entity.Payment
	ID := c.Param("id") // รับค่า UserID จาก URL

	db := config.DB()

	// ค้นหาข้อมูล Payment ที่ตรงกับ UserID
	result := db.Where("id = ?", ID).First(&payment)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found for this user"})
		return
	}

	// ตรวจสอบว่าข้อมูลที่ส่งมามีปัญหาหรือไม่
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}

	// อัปเดตข้อมูล Payment
	result = db.Save(&payment)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update payment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated successfully", "data": payment})
}

