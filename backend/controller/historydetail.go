package controller

import (
	"net/http"
	"backendproject/config"
	"backendproject/entity"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateHistoryDetail creates a new HistoryDetail entry
func CreateHistoryDetail(c *gin.Context) {
	var inputHistoryDetail entity.HistoryDetail

	// Bind JSON input to inputHistoryDetail
	if err := c.ShouldBindJSON(&inputHistoryDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB()

	// Create HistoryDetail instance
	historyDetail := entity.HistoryDetail{
		ProductName:  inputHistoryDetail.ProductName,
		Quantity:     inputHistoryDetail.Quantity,
		PricePerUnit: inputHistoryDetail.PricePerUnit,
		SubTotal:     inputHistoryDetail.SubTotal,
		StockID:      inputHistoryDetail.StockID,
		HistoryID:    inputHistoryDetail.HistoryID,
	}

	// Save to the database
	if err := db.Create(&historyDetail).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Send response
	c.JSON(http.StatusCreated, gin.H{
		"message": "Created success",
		"data":    historyDetail,
	})
}

// ListHistoryDetail lists all HistoryDetail entries
func ListHistoryDetail(c *gin.Context) {
	var historyDetails []entity.HistoryDetail

	db := config.DB()

	// Retrieve all HistoryDetail entries and preload related data
	result := db.Preload("Stock").
		Preload("History").
		Find(&historyDetails)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Send the result as JSON
	c.JSON(http.StatusOK, historyDetails)
}

// ListHistoryDetailByID retrieves a specific HistoryDetail by ID
func ListHistoryDetailByID(c *gin.Context) {
	// Get ID from URL parameter
	idParam := c.Param("id")

	// Convert ID to uint
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var historyDetail entity.HistoryDetail

	db := config.DB()

	// Retrieve HistoryDetail by ID and preload related data
	result := db.Preload("Stock").
		Preload("History").
		First(&historyDetail, uint(id))

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	// Send the result as JSON
	c.JSON(http.StatusOK, historyDetail)
}
