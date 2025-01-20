import React, { useState } from "react";
import { List, Button, Typography, Layout } from "antd";
import ClaimPopup from "./ClaimPopUp";
import "./History.css";

const { Text, Title } = Typography;
const { Header, Content } = Layout;

interface Product {
  id: number;
  name: string;
  image: string;
  price: number;
  quantity: number;
  orderId: string;
}
// ยังไม่ได้ใช้งาน
const Claim: React.FC = () => {
  const [selectedProduct, setSelectedProduct] = useState<Product | null>(null);
  const [isPopupVisible, setPopupVisible] = useState(false);

  const products: Product[] = [
    { id: 1, name: "Chair", image: "https://media1.popsugar-assets.com/files/thumbor/9YlQweTY57BMmnyN0GYURZUn4Vg/fit-in/728xorig/filters:format_auto-!!-:strip_icc-!!-/2019/02/04/995/n/1922794/6d9f7dd1b09ad63d_netimgcdnDgq/i/Roundhill-Furniture-Tuchico-Contemporary-Fabric-Accent-Chair.jpg", price: 110, quantity: 1, orderId: "1234567890" },
    { id: 2, name: "สินค้า2", image: "https://th.bing.com/th/id/OIP.mDTFng7t7daf__y_l39KigHaHa?w=191&h=191&c=7&r=0&o=5&pid=1.7", price: 220, quantity: 2, orderId: "0987654321" },
  ];

  const handleProductClick = (product: Product) => {
    setSelectedProduct(product.id === selectedProduct?.id ? null : product); // Toggle selection
  };

  const handleClaimClick = () => {
    setPopupVisible(true);
  };

  const handlePopupCancel = () => {
    setPopupVisible(false);
    setSelectedProduct(null);
  };

  const handlePopupSubmit = () => {
    console.log("Claim submitted for:", selectedProduct);
    setPopupVisible(false);
    setSelectedProduct(null);
  };


  return (
    <Layout>
      <Content className="layout-content">
        <div className="history-claim">
          <Title
            level={3}
            style={{
              textAlign: "center",
              marginBottom: "30px",
            }}
          >
            Claim
          </Title>
          <List
            dataSource={products}
            renderItem={(product) => (
              <List.Item
                key={product.id}
                className={`product-item ${selectedProduct?.id === product.id ? "selected" : ""}`}
                onClick={() => handleProductClick(product)}
                style={{
                  cursor: "pointer",
                  padding: "10px",
                  border: "1px solid #f0f0f0",
                  backgroundColor: selectedProduct?.id === product.id ? "#f5f5f5" : "white",
                }}
              >
                <img
                  src={product.image}
                  alt={product.name}
                  className="product-image"
                  style={{ width: 50, marginRight: 10 }}
                />
                <div className="product-info">
                  <Text>{product.name}</Text>
                  <Text>จำนวน: {product.quantity}</Text>
                  <Text>{product.price.toFixed(2)} THB</Text>
                </div>
              </List.Item>
            )}
          />
          {isPopupVisible && (
            <ClaimPopup
              visible={isPopupVisible}
              product={selectedProduct}
              onCancel={handlePopupCancel}
              onSubmit={handlePopupSubmit}
            />
          )}
        </div>
      </Content>
      {selectedProduct && ( // ปุ่มจะแสดงเฉพาะเมื่อมีการเลือกสินค้า
        <div className="fixed-buttons">
          <Button
            type="primary"
            onClick={handleClaimClick}
            style={{ marginRight: "10px" }}
          >
            เคลม
          </Button>
        </div>
      )}
    </Layout>
  );
};
export default Claim;



  