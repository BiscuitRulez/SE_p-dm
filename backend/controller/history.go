package controller

import (
	"net/http"
	"backendproject/config"
	"backendproject/entity"
	"time"
	"strconv"
	"github.com/gin-gonic/gin"
)



func CreateHistory(c *gin.Context) {
	var inputHistory entity.History

	// Bind JSON เข้ากับตัวแปร inputHistory
	if err := c.ShouldBindJSON(&inputHistory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB()

	// สร้าง History instance
	history := entity.History{
		OrderDate:      time.Now(), // กำหนด OrderDate เป็นเวลาปัจจุบัน
		UserID:         inputHistory.UserID,
		OrderID:        inputHistory.OrderID,
		PointsEarned:   inputHistory.PointsEarned,
		PointsRedeemed: inputHistory.PointsRedeemed,
		TotalAmount:    inputHistory.TotalAmount,
	}

	// บันทึกข้อมูลลงในฐานข้อมูล
	if err := db.Create(&history).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ส่ง Response
	c.JSON(http.StatusCreated, gin.H{
		"message": "Created success",
		"data":    history,
	})
}

func ListHistory(c *gin.Context) {
    // สร้าง slice สำหรับเก็บผลลัพธ์
    var histories []entity.History

    // ใช้การเชื่อมต่อกับฐานข้อมูล
    db := config.DB()

    // ดึงข้อมูลทั้งหมดจากตาราง History พร้อมโหลดความสัมพันธ์ที่เกี่ยวข้อง
    result := db.Preload("User").
        Preload("Order").
        Preload("HistoryDetails.Stock"). // โหลด Stock ที่สัมพันธ์กับ HistoryDetail
        Find(&histories)

    // ตรวจสอบข้อผิดพลาด
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    // ส่งผลลัพธ์กลับในรูปแบบ JSON
    c.JSON(http.StatusOK, histories)
}


func ListHistoryByID(c *gin.Context) {
    // รับค่า ID จากพารามิเตอร์ใน URL
    idParam := c.Param("id")

    // แปลงค่า ID เป็น uint
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    var history entity.History

    db := config.DB()

    // ดึงข้อมูล History พร้อมความสัมพันธ์ User, Order, และ HistoryDetails
    results := db.Preload("User").
        Preload("Order").
        Preload("HistoryDetails.Stock"). // โหลด Stock ที่สัมพันธ์กับ HistoryDetail
        First(&history, uint(id))

    // ตรวจสอบข้อผิดพลาด
    if results.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
        return
    }

    // ส่งคืนข้อมูล
    c.JSON(http.StatusOK, history)
}

 