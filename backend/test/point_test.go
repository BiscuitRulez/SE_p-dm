// ใน test/point_test.go
package test

import (
    "testing"
    "backendproject/entity"  // ใช้ path ที่ถูกต้องสำหรับ package entity

)


func TestTotalPoint(t *testing.T) {
    t.Run("TotalPoint_is_invalid", func(t *testing.T) {
        point := entity.Point{
            TotalPoint: 0, 
        }
        err := validate.Struct(point)
        if err == nil || err.Error() != "Key: 'Point.TotalPoint' Error:Field validation for 'TotalPoint' failed on the 'gt' tag" {
            t.Errorf("Expected error, got %v", err)
        }
    })

    t.Run("TotalPoint_is_positive", func(t *testing.T) {
        point := entity.Point{
            TotalPoint: 10, // ค่านี้ควรผ่านการตรวจสอบ
        }
        err := validate.Struct(point)
        if err != nil {
            t.Errorf("Expected no error, got %v", err)
        }
    })
}
