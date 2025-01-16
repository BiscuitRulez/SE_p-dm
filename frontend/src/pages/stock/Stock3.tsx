import React, { useEffect, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import axios from "axios";
import back from "../../assets/back.png";
import "../../components/card3.css";
import { message } from "antd";
import "../../components/t.css";
import { GetProductByID } from "../../services/https";

const Stock3: React.FC = () => {
  const location = useLocation();
  const navigate = useNavigate();
  const { productId } = location.state || {}; // รับ productId จาก state
  const [product, setProduct] = useState<any>(null); // ข้อมูลสินค้า
  const [stocks, setStocks] = useState<any[]>([]); // ข้อมูล stock
  const [loading, setLoading] = useState(true); // สถานะโหลดข้อมูล

  const getProductByID = async () => {
   
    try {
      const res = await GetProductByID(productId);
      setProduct(res.data.data);
      console.log("kuy rai nere :",res.data.data)

      
    } catch (error) {
      message.error("Error fetching product:");
    }
  };

  useEffect(() => {
    if (productId) {
      // ดึงข้อมูลสินค้า
      getProductByID();
      // ดึงข้อมูล stock
      // axios
      //   .get(`http://localhost:8000/stocks?product_id=${productId}`)เหลืออีก services นึง
      //   .then((response) => {
      //     setStocks(response.data.stocks || []);
      //   })
      //   .catch((error) => {
      //     console.error("Error fetching stocks:", error);
      //   })
      //   .finally(() => {
      //     setLoading(false);
      //   });
    } else {
      setLoading(false);
    }
  }, [productId]);

  // ฟังก์ชันลบ stock
  const handleDeleteClick = async (stockId: number) => {
    if (isNaN(stockId)) {
      message.error("Invalid stock ID");
      return;
    }
  
    const confirmDelete = window.confirm("คุณแน่ใจหรือไม่ว่าต้องการลบรายการนี้?");
    if (confirmDelete) {
      // Optimistic UI: ลบ stock ทันทีจาก state
      setStocks((prevStocks) => prevStocks.filter((stock) => stock.id !== stockId));
  
      try {
        const response = await axios.delete(`http://localhost:8000/stock/${stockId}`); //สร้าง services delete
        if (response.status !== 200) {
          throw new Error("ไม่สามารถลบ stock ได้");
        }
        alert(response.data.message);
      } catch (error) {
        console.error("เกิดข้อผิดพลาดระหว่างการลบ stock:", error);
        message.error("เกิดข้อผิดพลาดระหว่างการลบ");
        // ถ้ามีข้อผิดพลาด ให้ย้อนกลับ UI
        setStocks((prevStocks) => [...prevStocks, { id: stockId }]);
      }
    }
  };
  
   
  

  // ฟังก์ชันไปที่หน้าแก้ไข stock
  const handleEditClick = (stock: any) => {
    navigate("/Stock4", { state: { productId, product, stock } });
  };

  // ฟังก์ชันไปที่หน้าเพิ่ม stock
  const handleNext4Click = () => {
    navigate("/Stock4", { state: { productId, product } }); // ส่ง product ไปยังหน้า Stock4
  };

  if (loading) {
    return <div>Loading...</div>; // แสดงข้อความโหลด
  }

  if (!productId || !product) {
    return <div>ไม่พบข้อมูลสินค้า</div>; // กรณีไม่มี productId หรือข้อมูลสินค้า
  }

  return (
    <div>
    
      <a href="/Stock2" style={{ position: "absolute", top: "100px", right: "1450px" }}>
        <img style={{ width: "50px", height: "auto" }} src={back} alt="Back" />
      </a>

      {/* ข้อมูลสินค้า */}
      <div className="product-card3">
        <img className="product-card-img3" src={product.image} alt={product.name} />
        <div className="product-content3">
          <h3 className="product-title3">{product.name}</h3>
          <p className="product-description3">{product.description}</p>
        </div>
      </div>

      {/* ข้อมูล stock */}
      <div className="card-container4">
  {stocks.length > 0 ? (
    stocks.map((stock) => (
      <div className="card4" key={stock.id}>
        <img src={stock.image || "/assets/placeholder.png"} alt={stock.color} className="card-image4" />
        <h3 className="card-title4">{stock.color}</h3>
        <div className="card-actions4">
          <button className="edit-btn" onClick={() => handleEditClick(stock)}>แก้ไข</button>
          <button className="delete-btn" onClick={() => handleDeleteClick(stock.id)}>ลบ</button>
        </div>
      </div>
    ))
  ) : (
    <div className="textno">ไม่มี stock สำหรับสินค้านี้</div>
  )}

  {/* การ์ดเพิ่ม stock */}
  <div className="card4" onClick={handleNext4Click}>
    <div className="add-card4">
      <div className="add-image4">
        <span>+</span>
      </div>
    </div>
    <h3 className="card-title4"></h3>
  </div>
</div>

    </div>
  );
};

export default Stock3;
