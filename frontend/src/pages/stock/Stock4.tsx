import React, { useEffect, useState } from "react";

import back from "../../assets/back.png";
import "../../components/card4.css";
import { useNavigate, useLocation } from "react-router-dom";
import { message } from "antd";
import "../../components/t.css";

const Stock4: React.FC = () => {
  const [quantity_in_stock, setQuantity] = useState(1);
  const [price, setPrice] = useState(1);
  const [color, setColor] = useState("");
  const [shape_size, setSize] = useState("");
  const [image, setImage] = useState<string | null>(null);

  const navigate = useNavigate();
  const location = useLocation();
  const { productId, product, stock } = location.state || {}; // รับ product, productId, และ stock

  if (!productId || !product) {
    return <div>ไม่พบข้อมูลสินค้า</div>;
  }

  useEffect(() => {
    if (stock) {
      setColor(stock.color);
      setSize(stock.shape_size);
      setPrice(stock.price);
      setQuantity(stock.quantity_in_stock);
      setImage(stock.image || null);
    }
  }, [stock]);

  const handleQuantityChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = Math.max(1, Number(e.target.value));
    setQuantity(value);
  };

  const handlePriceChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = Math.max(1, Number(e.target.value));
    setPrice(value);
  };

  const handleImageUpload = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files[0]) {
      const file = e.target.files[0];
      const reader = new FileReader();

      reader.onload = () => {
        setImage(reader.result as string);
      };
      reader.readAsDataURL(file);
    }
  };

  const handleSave = async () => {
    if (!color || !shape_size || !price || !quantity_in_stock || !image) {
      message.error("กรุณากรอกข้อมูลให้ครบถ้วน!");
      return;
    }

    const stockData = {
      product_id: productId,
      color,
      shape_size,
      price,
      quantity_in_stock,
      image,
    };

    try {
      const url = stock?.id
        ? `http://localhost:8000/stock/${stock.id}` // อัปเดตข้อมูล
        : "http://localhost:8000/stock"; // เพิ่มใหม่
      const method = stock?.id ? "PUT" : "POST";

      const response = await fetch(url, {
        method,
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(stockData),
      });

      if (response.ok) {
        message.success(stock?.id ? "อัปเดตข้อมูลเรียบร้อย!" : "เพิ่มข้อมูลเรียบร้อย!");
        navigate("/Stock3", { state: { productId } });
      } else {
        message.error("เกิดข้อผิดพลาดในการบันทึกข้อมูล!");
      }
    } catch (error) {
      console.error("Error saving stock:", error);
      message.error("ไม่สามารถเชื่อมต่อกับเซิร์ฟเวอร์ได้!");
    }
  };

  const handleCancel = () => {
    navigate("/Stock3", { state: { productId } });
  };

  return (
    <div>
      
    
      
      <a
        onClick={handleCancel}
        style={{ position: "absolute", top: "100px", right: "1400px", cursor: "pointer" }}
      >
        <img style={{ width: "50px", height: "auto" }} src={back} alt="back" />
      </a>

      <h1>เพิ่ม/แก้ไขรายการสินค้า</h1>

      <div className="product-card5">
        <img className="product-card-img5" src={product.image} alt={product.name} />
        <div className="product-content5">
          <h3 className="product-title5">{product.name}</h3>
          <p className="product-description5">{product.description}</p>
        </div>
      </div>

      <div className="form-container">
        <div className="form-section">
          <label htmlFor="color">สี</label>
          <input
            type="text"
            id="color"
            className="input-field"
            value={color}
            onChange={(e) => setColor(e.target.value)}
          />

          <label htmlFor="size">รูปร่าง/ขนาด</label>
          <input
            type="text"
            id="size"
            className="input-field"
            value={shape_size}
            onChange={(e) => setSize(e.target.value)}
          />

          <div className="quantity-container">
            <label htmlFor="price">ราคา</label>
            <label htmlFor="quantity">จำนวนสินค้าในคลัง</label>
          </div>
          <div className="quantity-container">
            <div className="quantity-box">
              <input
                type="number"
                min="1"
                value={price}
                onChange={handlePriceChange}
                className="input-field"
              />
            </div>

            <div className="quantity-box">
              <input
                type="number"
                min="1"
                value={quantity_in_stock}
                onChange={handleQuantityChange}
                className="input-field"
              />
            </div>
          </div>
        </div>

        <div className="image-upload">
          <label htmlFor="image-upload" className="upload-box" style={{ cursor: "pointer" }}>
            {image ? (
              <img src={image} alt="Uploaded" style={{ width: "100%", height: "auto" }} />
            ) : (
              <span className="upload-icon">+</span>
            )}
          </label>
          <input
            type="file"
            id="image-upload"
            style={{ display: "none" }}
            accept="image/*"
            onChange={handleImageUpload}
          />
        </div>

        <div className="button-section">
          <button className="cancel-btn" onClick={handleCancel}>
            ยกเลิก
          </button>
          <button className="save-btn" onClick={handleSave}>
            บันทึก
          </button>
        </div>
      </div>
    </div>
  );
};

export default Stock4;
