package codes

import (
	"net/http"
	"strconv"

	"backendproject/config"
	"backendproject/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAll(c *gin.Context) {

	db := config.DB()

	var codes []entity.Codes

	db.Find(&codes)

	c.JSON(http.StatusOK, &codes)
}

func GetCodeById(c *gin.Context) {
	// รับ ID จากพารามิเตอร์และตรวจสอบว่าเป็นตัวเลข
	ID := c.Param("id")
	codeID, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var code []entity.Codes

	// ดึงข้อมูลจากฐานข้อมูล
	db := config.DB()
	results := db.First(&code, codeID)
	if results.Error != nil {
		// ตรวจสอบว่าข้อผิดพลาดคือ record not found หรือไม่
		if results.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Code not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error: " + results.Error.Error()})
		}
		return
	}

	// ส่ง response พร้อมข้อมูล code
	c.JSON(http.StatusOK, code)
}


func CreateCode(c *gin.Context) {

    var newCode entity.Codes // สร้างตัวแปรสำหรับเก็บข้อมูล Code ใหม่

    // ผูก JSON ที่ส่งมาจาก Request Body กับตัวแปร newCode
    if err := c.ShouldBindJSON(&newCode); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // เชื่อมต่อกับฐานข้อมูล
    db := config.DB()

	nc := entity.Codes{

		CodeTopic:	newCode.CodeTopic,

		CodeDescription: newCode.CodeDescription,
	
		Discount:	newCode.Discount,

		Quantity: newCode.Quantity,

        Minimum: newCode.Minimum,

		DateStart: newCode.DateStart,

		DateEnd: newCode.DateEnd,

		CodeStatus: newCode.CodeStatus,

		CodePicture: newCode.CodePicture,
       
	}

    // บันทึกข้อมูลลงในฐานข้อมูล
    if err := db.Create(&nc).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // ส่งข้อมูลที่สร้างสำเร็จกลับไป
    c.JSON(http.StatusCreated, gin.H {"message": "Code created successfully","code": nc})
}


func UpdateCode(c *gin.Context) {

	var code entity.Codes

	CodeID := c.Param("id")

	db := config.DB()
	result := db.First(&code, CodeID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}

	if err := c.ShouldBindJSON(&code); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}

	result = db.Save(&code)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})
}


func DeleteCode(c *gin.Context) {

	id := c.Param("id")
	db := config.DB()
	if tx := db.Exec("DELETE FROM codes WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})

}

func UpdateCodeAfterCollect(c *gin.Context) {

    id := c.Param("id")
	db := config.DB()
    var code entity.Codes

    if err := db.First(&code, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Code not found"})
        return
    }

    // ลดจำนวนโค้ดหรืออัปเดตสถานะโค้ด
    code.Quantity -= 1
    if code.Quantity < 0 {
        code.Quantity = 0
    }

    db.Save(&code)

    c.JSON(http.StatusOK, code)
}

func AddCodeToCollect(c *gin.Context) {
    userIDStr := c.Param("userId")
    codeIDStr := c.Param("codeId")

    userID, err := strconv.ParseUint(userIDStr, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    codeID, err := strconv.ParseUint(codeIDStr, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid code ID"})
        return
    }

    db := config.DB()

    // ตรวจสอบว่ามีข้อมูลนี้อยู่ในระบบแล้วหรือไม่
    var existingCollect entity.CodeCollectors
    if err := db.Where("user_id = ? AND code_id = ?", uint(userID), uint(codeID)).First(&existingCollect).Error; err == nil {
        c.JSON(http.StatusConflict, gin.H{"message": "Code already collected"})
        return
    }

    collect := entity.CodeCollectors{
        UserID: uint(userID),
        CodeID: uint(codeID),
    }

    if err := db.Create(&collect).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save code collect"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Code collected successfully", "data": collect})
}

func GetCollectedCodesToShow(c *gin.Context) {
    // รับค่า userId จาก URL parameter
    userIdParam := c.Param("userId")

    // แปลง userId จาก string เป็น uint
    userId, err := strconv.Atoi(userIdParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    // ดึงฐานข้อมูล
    db := config.DB()

    // กำหนด struct สำหรับผลลัพธ์ที่ต้องการ
    var results []entity.Codes  // เปลี่ยนให้เป็น struct `Codes` แทนผลลัพธ์ที่แยกต่างหาก

    // ใช้ Gorm สำหรับ JOIN และดึงข้อมูลทั้งหมดจากตาราง codes
    if err := db.Table("code_collectors").
        Select("DISTINCT codes.*").  // ดึงข้อมูลทั้งหมดจากตาราง codes
        Joins("INNER JOIN codes ON code_collectors.code_id = codes.id").
        Where("code_collectors.user_id = ?", userId).
        Scan(&results).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch collected codes"})
        return
    }

    // ส่งผลลัพธ์กลับไป
    if len(results) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "No codes found for the given user"})
        return
    }

    // c.JSON(http.StatusOK, gin.H{"data": results})  // ส่งข้อมูลทั้งหมดของ `codes`
    c.JSON(http.StatusOK, results)
}


func GetCollectedCodes(c *gin.Context) {
    userId := c.Param("userId")
    
    // Query or process the collected codes for the given userId
    db := config.DB()
    var collectedCodes []entity.CodeCollectors
    if err := db.Where("user_id = ?", userId).Find(&collectedCodes).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "No codes collected for user"})
        return
    }

    c.JSON(http.StatusOK, collectedCodes)
}






 
