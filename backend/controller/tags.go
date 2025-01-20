package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"backendproject/config"
	"backendproject/entity"
)

func GetAllTags(c *gin.Context) {
    var tag []entity.Tags

    // Attempt to retrieve all airlines from the database
    if err := config.DB().Find(&tag).Error; err != nil {
        // If there's an error, return a 500 status code with the error message
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // If successful, return the list of airlines with a 200 status code
    c.JSON(http.StatusOK, gin.H{"tags": tag})
}


func GetTagsByID(c *gin.Context) {
	var tag entity.Tags
	id := c.Param("id")

	// ดึงข้อมูลจากฐานข้อมูลตาม ID
	if err := config.DB().First(&tag, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tags not found"})
		return
	}

	// ส่งข้อมูลกลับไปในรูป JSON
	c.JSON(http.StatusOK, gin.H{"data": tag})
}

func CreateTags(c *gin.Context) {

    var newTags entity.Tags // สร้างตัวแปรสำหรับเก็บข้อมูล Code ใหม่

    // ผูก JSON ที่ส่งมาจาก Request Body กับตัวแปร newCode
    if err := c.ShouldBindJSON(&newTags); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // เชื่อมต่อกับฐานข้อมูล
    db := config.DB()

	nt := entity.Tags{

		Tag_Name:	newTags.Tag_Name,

	}

    // บันทึกข้อมูลลงในฐานข้อมูล
    if err := db.Create(&nt).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // ส่งข้อมูลที่สร้างสำเร็จกลับไป
    c.JSON(http.StatusCreated, gin.H {"message": "Tag created successfully","tag": nt})
}

func UpdateTagsByID(c *gin.Context) {
	var tag entity.Tags
 
	TagsID := c.Param("id")
 
	db := config.DB()
 
	result := db.First(&tag, TagsID)
 
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}
 
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}
 
	result = db.Save(&tag)
 
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
 
	}
	c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})
 
}

func DeleteTags(c *gin.Context) {

	id := c.Param("id")
	db := config.DB()
	if tx := db.Exec("DELETE FROM Tags WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})

}
