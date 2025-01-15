package controller

import (
	"net/http"

	"backendproject/config"
	"backendproject/entity"
	"time"

	"github.com/gin-gonic/gin"
)

// GetClaims - ดึงข้อมูล Claims ทั้งหมด
func GetClaims(c *gin.Context) {
	var claims []entity.Claim
	db := config.DB()

	if result := db.Preload("User").Preload("Order").Preload("Problem").Preload("ClaimStatus").Find(&claims); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, claims)
}

// GetClaimByID - ดึง Claim ตาม ID
func GetClaimByID(c *gin.Context) {
	id := c.Param("id")
	var claim entity.Claim
	db := config.DB()

	if result := db.Preload("User").Preload("Order").Preload("Problem").Preload("ClaimStatus").First(&claim, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, claim)
}

// CreateClaim - สร้าง Claim ใหม่
func CreateClaim(c *gin.Context) {
	var request struct {
		Date        time.Time `json:"date"`
		Photo       string    `json:"photo"`
		ProblemID   uint      `json:"problem_id"`
		ClaimStatusID uint    `json:"claim_status_id"`
		UserID      uint      `json:"user_id"`
		OrderID     uint      `json:"order_id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB()

	// ตรวจสอบ Problem, ClaimStatus, User, Order
	var problem entity.Problem
	var claimStatus entity.ClaimStatus
	var user entity.Users
	var order entity.Order

	if db.First(&problem, request.ProblemID).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Problem ID"})
		return
	}
	if db.First(&claimStatus, request.ClaimStatusID).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Claim Status ID"})
		return
	}
	if db.First(&user, request.UserID).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}
	if db.First(&order, request.OrderID).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Order ID"})
		return
	}

	// สร้าง Claim ใหม่
	newClaim := entity.Claim{
		Date:         request.Date,
		Photo:        request.Photo,
		ProblemID:    request.ProblemID,
		ClaimStatusID: request.ClaimStatusID,
		UserID:       request.UserID,
		OrderID:      request.OrderID,
	}

	if err := db.Create(&newClaim).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Claim created successfully", "data": newClaim})
}

// UpdateClaimByID - อัปเดต Claim
func UpdateClaimByID(c *gin.Context) {
	id := c.Param("id")
	var claim entity.Claim

	db := config.DB()
	if result := db.First(&claim, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Claim not found"})
		return
	}

	if err := c.ShouldBindJSON(&claim); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := db.Save(&claim); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Claim updated successfully", "data": claim})
}

// DeleteClaim - ลบ Claim
func DeleteClaim(c *gin.Context) {
	id := c.Param("id")
	db := config.DB()

	if result := db.Delete(&entity.Claim{}, id); result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Claim not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Claim deleted successfully"})
}
