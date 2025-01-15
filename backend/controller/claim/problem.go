package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"backendproject/entity"
	"backendproject/config"
)

func GetProblem(c *gin.Context) {
    // สร้าง slice สำหรับเก็บผลลัพธ์
    var problem []entity.Claim

    // ใช้การเชื่อมต่อกับฐานข้อมูล
    db := config.DB()

    // ดึงข้อมูลทั้งหมดจากตาราง payment_methods
    if err := db.Find(&problem).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // ส่งผลลัพธ์กลับในรูปแบบ JSON
    c.JSON(http.StatusOK, gin.H{
        "message": "Fetched successfully",
        "data":    problem,
    })
}
