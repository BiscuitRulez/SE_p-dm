package controller

import (
	"net/http"
	"time"
	"fmt"
	"backendproject/config"
	"backendproject/entity"
	"github.com/gin-gonic/gin"
)

// GetPointsByUserID - ดึงข้อมูลแต้มสะสมของผู้ใช้ตาม ID
func GetPointsByUserID(c *gin.Context) {
	userID := c.Param("userId")

	var point entity.Point
	if err := config.DB().Where("id = ?", userID).First(&point).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Points not found"})
		return
	}

	c.JSON(http.StatusOK, point)
}

// RedeemPoints - แลกแต้มสะสม
func RedeemPoints(c *gin.Context) {
    var req struct {
        UserID      uint `json:"user_id"`
        RedeemPoint uint `json:"redeem_point"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // ดึงข้อมูล PointPolicy
    var pointPolicy entity.PointPolicy
    if err := config.DB().First(&pointPolicy).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Point policy not found"})
        return
    }

    // ดึงข้อมูล Point ของผู้ใช้
    var point entity.Point
    if err := config.DB().Where("id = ?", req.UserID).First(&point).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Points not found"})
        return
    }

    // ตรวจสอบว่าแต้มเพียงพอหรือไม่
    if point.TotalPoint < req.RedeemPoint {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough points"})
        return
    }

    // คำนวณส่วนลดจากแต้มแลก
    discount := float64(req.RedeemPoint) / float64(pointPolicy.RedeemRate)

    // หักแต้ม
    point.TotalPoint -= req.RedeemPoint
    config.DB().Save(&point)

	// บันทึกประวัติการแลกแต้มลงใน PointTran
	pointTran := entity.PointTran{
		PointsRedeemed:  req.RedeemPoint,
		TransactionType: "redeem",
		TransactionDate: time.Now(),
		Description:     fmt.Sprintf("Redeemed %d points", req.RedeemPoint),
		PointID:         point.ID,
	}
	config.DB().Create(&pointTran)

    c.JSON(http.StatusOK, gin.H{
        "message":          "Points redeemed successfully",
        "remaining_points": point.TotalPoint,
        "discount":         discount,
    })
}


// EarnPoints - เพิ่มแต้มสะสม
func EarnPoints(c *gin.Context) {
    var req struct {
        UserID      uint    `json:"user_id"`
        AmountSpent float64 `json:"amount_spent"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // ดึงข้อมูล PointPolicy
    var pointPolicy entity.PointPolicy
    if err := config.DB().First(&pointPolicy).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Point policy not found"})
        return
    }

    // คำนวณแต้มสะสมจาก EarnRate
    earnedPoints := uint(req.AmountSpent) * pointPolicy.EarnRate

    // ดึงข้อมูล Point ของผู้ใช้
    var point entity.Point
    if err := config.DB().Where("id = ?", req.UserID).First(&point).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Points not found"})
        return
    }

    // เพิ่มแต้มสะสม
    point.TotalPoint += earnedPoints
    config.DB().Save(&point)

	// บันทึกประวัติการสะสมแต้มลงใน PointTran
	pointTran := entity.PointTran{
		PointsEarned:    earnedPoints,
		TransactionType: "earn",
		TransactionDate: time.Now(),
		Description:     fmt.Sprintf("Earned %d points from spending %.2f", earnedPoints, req.AmountSpent),
		PointID:         point.ID,
	}
	config.DB().Create(&pointTran)

    c.JSON(http.StatusOK, gin.H{"message": "Points earned successfully", "total_points": point.TotalPoint})
}


// UpdatePoints - อัปเดตแต้มสะสม
func UpdatePoints(c *gin.Context) {
	pointID := c.Param("pointId")

	var req struct {
		TotalPoint uint `json:"total_point"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var point entity.Point
	if err := config.DB().Where("id = ?", pointID).First(&point).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Points not found"})
		return
	}

	point.TotalPoint = req.TotalPoint
	config.DB().Save(&point)

	c.JSON(http.StatusOK, gin.H{"message": "Points updated successfully", "total_points": point.TotalPoint})
}

// DeletePoints - ลบข้อมูลแต้มสะสม (ถ้าจำเป็น)
func DeletePoints(c *gin.Context) {
	pointID := c.Param("pointId")

	if err := config.DB().Delete(&entity.Point{}, pointID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete points"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Points deleted successfully"})
}
