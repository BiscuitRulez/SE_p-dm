package controller

import (
	"backendproject/config"
	"backendproject/entity"
	"net/http"
	//"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllProduct(c *gin.Context) {
    var product []entity.Product

    // Attempt to retrieve all airlines from the database
    if err := config.DB().Find(&product).Error; err != nil {
        // If there's an error, return a 500 status code with the error message
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // If successful, return the list of airlines with a 200 status code
    c.JSON(http.StatusOK, gin.H{"products": product})
}


func GetProductByID(c *gin.Context) {
	var product entity.Product
	id := c.Param("id")

	if err := config.DB().First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func CreateProduct(c *gin.Context) {
    var product entity.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
        return
    }
    db := config.DB()
    result := db.Create(&product)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "Created successful", "stock": product})
 }

// func CreateProduct(c *gin.Context) {

//     var newProduct entity.Product // สร้างตัวแปรสำหรับเก็บข้อมูล Code ใหม่

//     // ผูก JSON ที่ส่งมาจาก Request Body กับตัวแปร newCode
//     if err := c.ShouldBindJSON(&newProduct); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
//     }

//     // เชื่อมต่อกับฐานข้อมูล
//     db := config.DB()

// 	if newProduct.CatagoryID == 0 || newProduct.UserID == 0 {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "Category, User, fields must have valid IDs"})
//         return
//     }

// 	var catagory entity.Catagory
//     if err := db.First(&catagory, newProduct.CatagoryID).Error; err != nil {
//         c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
//         return
//     }

// 	var user entity.User
//     if err := db.First(&user, newProduct.UserID).Error; err != nil {
//         c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
//         return
//     }

// 	if newProduct.Name == "" || newProduct.Description == "" || newProduct.Quantity < 0 {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product details"})
//         return
//     }

// 	np := entity.Product{

// 		Name:	newProduct.Name,

// 		Description: newProduct.Description,
	
// 		Quantity:	newProduct.Quantity,

// 		Image: newProduct.Image,

// 	}

//     // บันทึกข้อมูลลงในฐานข้อมูล
//     if err := db.Create(&np).Error; err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }

//     // ส่งข้อมูลที่สร้างสำเร็จกลับไป
//     c.JSON(http.StatusCreated, gin.H {"message": "Product created successfully","product_id": np.ID,})
// }

// func UpdateProductByID(c *gin.Context) {
// 	var product entity.Product
 
// 	ProductID := c.Param("id")
 
// 	db := config.DB()
 
// 	result := db.First(&product, ProductID)
 
// 	if result.Error != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
// 		return
// 	}
 
// 	if err := c.ShouldBindJSON(&product); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
// 		return
// 	}
 
// 	result = db.Save(&product)
 
// 	if result.Error != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
// 		return
 
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})
 
// }


func UpdateProductByID(c *gin.Context) {


	var product entity.Product
 
 
	ProductID := c.Param("id")
 
 
	db := config.DB()
 
	result := db.First(&product, ProductID)
 
	if result.Error != nil {
 
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
 
		return
 
	}
 
 
	if err := c.ShouldBindJSON(&product); err != nil {
 
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
 
		return
 
	}
 
 
	result = db.Save(&product)
 
	if result.Error != nil {
 
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
 
		return
 
	}
 
 
	c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})
 
 }



func DeleteProduct(c *gin.Context) {

	id := c.Param("id")
	db := config.DB()
	if tx := db.Exec("DELETE FROM products WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})

}
