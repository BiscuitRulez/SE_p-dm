package controller


import (
    "errors"
    "gorm.io/gorm"
   "net/http"


   "github.com/gin-gonic/gin"


   "backendproject/config"

   "backendproject/entity"

)

func GetStocksByProductID(c *gin.Context) {
	productID := c.Query("product_id")

	var stocks []entity.Stock
	if err := config.DB().Where("product_id = ?", productID).Find(&stocks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"stocks": stocks})
}

func GetAll(c *gin.Context) {


   var stocks []entity.Stock


   db := config.DB()

   results := db.Find(&stocks)

   if results.Error != nil {

       c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})

       return

   }

   c.JSON(http.StatusOK, stocks)


}


func Get(c *gin.Context) {


   ID := c.Param("id")

   var stock entity.Stock


   db := config.DB()

   results := db.First(&stock, ID)

   if results.Error != nil {

       c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})

       return

   }

   if stock.ID == 0 {

       c.JSON(http.StatusNoContent, gin.H{})

       return

   }

   c.JSON(http.StatusOK, stock)


}


func Update(c *gin.Context) {


   var stock entity.Stock


   StockID := c.Param("id")


   db := config.DB()

   result := db.First(&stock, StockID)

   if result.Error != nil {

       c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})

       return

   }


   if err := c.ShouldBindJSON(&stock); err != nil {

       c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})

       return

   }


   result = db.Save(&stock)

   if result.Error != nil {

       c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})

       return

   }


   c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})

}


func Delete(c *gin.Context) {
	id := c.Param("id")
	db := config.DB()
  
	result := db.Delete(&entity.Stock{}, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// ตรวจสอบเฉพาะกรณีที่ไม่พบ record
			c.JSON(http.StatusNotFound, gin.H{"error": "Stock with ID not found"})
		} else {
			// จัดการกรณีข้อผิดพลาดอื่นๆ เช่น ข้อผิดพลาดฐานข้อมูล
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}
  
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})
}
  

func Create(c *gin.Context) {
    var stock entity.Stock
    if err := c.ShouldBindJSON(&stock); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
        return
    }
    db := config.DB()
    result := db.Create(&stock)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create stock"})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "Created successful", "stock": stock})
 }