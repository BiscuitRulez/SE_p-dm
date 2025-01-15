package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"backendproject/entity"
	"backendproject/config"
)

// Mock data store (you should replace this with database queries)
// var paymentStatus = []entity.Payment{}

// Get all payments
// func GetPaymentStatus(c *gin.Context) {
// 	// filter := c.Query("filter") // รับค่าพารามิเตอร์ 'filter' จาก query string

// 	// Get the database connection
// 	db := config.DB()

// 	// ใช้ GORM Query แบบปลอดภัย
// 	query := db.Select("paymentStatus.id, paymentStatus.payment_status")

// 	// ถ้ามี filter ให้เพิ่มเงื่อนไขการค้นหา
// 	// if filter != "" {
// 	// 	query = query.Where("Movie_name LIKE ?", "%"+filter+"%")
// 	// }

// 	// Execute the query
// 	result := query.Find(&paymentStatus)

// 	// Check for errors in the query
// 	if result.Error != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
// 		return
// 	}

// 	// Return the results as JSON
// 	c.JSON(http.StatusOK, paymentStatus)	
// }

func GetPaymentStatus(c *gin.Context) {
	// สร้าง slice สำหรับเก็บผลลัพธ์
	var paymentStatus []entity.PaymentStatus

	// ใช้การเชื่อมต่อกับฐานข้อมูล
	db := config.DB()

	// ดึงข้อมูลทั้งหมดจากตาราง payments
	result := db.Find(&paymentStatus)

	// ตรวจสอบข้อผิดพลาด
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// ส่งผลลัพธ์กลับในรูปแบบ JSON
	c.JSON(http.StatusOK, paymentStatus)
}