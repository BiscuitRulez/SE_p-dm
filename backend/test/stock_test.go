package test

import (
    "testing"
    "backendproject/entity" // ปรับให้เป็นที่อยู่ของ package model ของคุณ
)


func TestAllStock(t *testing.T) {
    t.Run("All_Correct", func(t *testing.T) {
        stock := entity.Stock{
            ID:          123,
            Price:       100,
            Quantity:    10,
            Color:       "Red",
            ShapeSize:   "Medium",
            Image:       "/assets/sofa.jpg",
            ProductID:  1,
        }
        err := validate.Struct(stock)
        if err != nil {
            t.Errorf("Expected no error, got %v", err)
        }
    })
}

func TestID(t *testing.T) {
    t.Run("ID_is_invalid", func(t *testing.T) {
        stock := entity.Stock{
            ID:          0,
            Price:       100,
            Quantity:    10,
            Color:       "Red",
            ShapeSize:   "Medium",
            Image:       "/assets/sofa.jpg",
            ProductID:  1,
        }
        err := validate.Struct(stock)
        if err != nil {
            if err.Error() != "ID cannot be 0" {
                t.Errorf("Expected ID cannot be empty, got %v", err)
            }
        }
    })
}

func TestPrice(t *testing.T) {
    t.Run("Price_is_invalid", func(t *testing.T) {
        stock := entity.Stock{
            ID:          123,
            Price:       0,
            Quantity:    10,
            Color:       "Red",
            ShapeSize:   "Medium",
            Image:       "/assets/sofa.jpg",
            ProductID:  1,
        }
        err := validate.Struct(stock)
        if err != nil {
            if err.Error() != "Price cannot be 0" {
                t.Errorf("Expected Price cannot be 0, got %v", err)
            }
        }
    })
}

func TestQuantityStock(t *testing.T) {
    t.Run("Quantity_is_invalid", func(t *testing.T) {
        stock := entity.Stock{
            ID:          123,
            Price:       100,
            Quantity:    0,
            Color:       "Red",
            ShapeSize:   "Medium",
            Image:       "/assets/sofa.jpg",
            ProductID:  1,
        }
        err := validate.Struct(stock)
        if err != nil {
            if err.Error() != "Quantity cannot be 0" {
                t.Errorf("Expected Quantity cannot be 0, got %v", err)
            }
        }
    })
}

func TestColor(t *testing.T) {
    t.Run("Color_is_invalid", func(t *testing.T) {
        stock := entity.Stock{
            ID:          123,
            Price:       100,
            Quantity:    10,
            Color:       "",
            ShapeSize:   "Medium",
            Image:       "/assets/sofa.jpg",
            ProductID:  1,
        }
        err := validate.Struct(stock)
        if err != nil {
            if err.Error() != "Color cannot be empty" {
                t.Errorf("Expected Color cannot be empty, got %v", err)
            }
        }
    })
}

func TestShapeSize(t *testing.T) {
    t.Run("ShapeSize_is_invalid", func(t *testing.T) {
        stock := entity.Stock{
            ID:          123,
            Price:       100,
            Quantity:    10,
            Color:       "Red",
            ShapeSize:   "",
            Image:       "/assets/sofa.jpg",
            ProductID:  1,
        }
        err := validate.Struct(stock)
        if err != nil {
            if err.Error() != "ShapeSize cannot be empty" {
                t.Errorf("Expected ShapeSize cannot be empty, got %v", err)
            }
        }
    })
}

func TestImage(t *testing.T) {
    t.Run("Image_is_invalid", func(t *testing.T) {
        stock := entity.Stock{
            ID:          123,
            Price:       100,
            Quantity:    10,
            Color:       "Red",
            ShapeSize:   "Medium",
            Image:       "",
            ProductID:  1,
        }
        err := validate.Struct(stock)
        if err != nil {
            if err.Error() != "Image cannot be empty" {
                t.Errorf("Expected Image cannot be empty, got %v", err)
            }
        }
    })
}

func TestProductID(t *testing.T) {
    t.Run("ProductID_is_invalid", func(t *testing.T) {
        stock := entity.Stock{
            ID:          123,
            Price:       0,
            Quantity:    10,
            Color:       "Red",
            ShapeSize:   "Medium",
            Image:       "/assets/sofa.jpg",
            ProductID:  0,
        }
        err := validate.Struct(stock)
        if err != nil {
            if err.Error() != "ProductID cannot be 0" {
                t.Errorf("Expected ProductID cannot be 0, got %v", err)
            }
        }
    })
}
