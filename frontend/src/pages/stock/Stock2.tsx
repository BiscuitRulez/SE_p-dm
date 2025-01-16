import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import "../../components/card2.css";
import { GetProducts } from '../../services/https';

interface Product {
  id: number;
  name: string;
  description: string;
  image: string;
}

const Stock2: React.FC = () => {
  const [products, setProducts] = useState<Product[]>([]); // ใช้ interface Product
  const navigate = useNavigate();

  const getProducts = async () => {
    try {
      let res = await GetProducts();
      console.log("GetProduct Response: ", res.data.products); // ตรวจสอบข้อมูลที่ได้จาก API
      if (res) {
        setProducts(res.data.products || []);
      }
    } catch (error) {
      console.error("Error fetching product: ", error); // ดูว่ามี error อะไร
    }
  };

  // ฟังก์ชันดึงข้อมูลจาก API
  useEffect(() => {
    getProducts();
  }, []);

  const handleNext3Click = (productId: number) => {
    navigate(`/Stock2/Stock3`, { state: { productId } });
  };

  return (
    <div>
      <a href="/" className="back-button">
        <img src="/assets/back.png" alt="Back" className="back-img" />
      </a>

      <h1 className="title">กรุณาเลือกสินค้า</h1>


      <div className="product-grid2">
        {products.map((product) => (
          <button
            key={product.id}
            onClick={() => handleNext3Click(product.id)}
            className="product-card2"
          >
            <img className="product-card-img2" src={product.image} alt={product.name} />
            <div className="product-content2">
              <h2 className="product-title2">{product.name}</h2>
              <p className="product-description2">{product.description}</p>
            </div>
          </button>
        ))}
      </div>
    </div>
  );
};

export default Stock2;
