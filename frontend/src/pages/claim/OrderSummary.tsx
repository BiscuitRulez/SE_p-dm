import React from "react";
import "./OrderSummary.css"; // สร้างไฟล์ CSS แยกต่างหากเพื่อจัดการสไตล์
import { Modal } from "antd"; // ใช้ Modal จาก Ant Design

const OrderSummary: React.FC<{ visible: boolean; onClose: () => void }> = ({ visible, onClose }) => {
  const orderDetails = {
    orderNumber: "65416508916",
    orderDate: "ธ.ค. 20, 2024",
    deliveryDate: "ธ.ค. 25, 2024",
    total: 891.15,
    shippingFee: 30.0,
    discount: -30.0,
    vat: -8.15,
    grandTotal: 890.0,
  };

  const status: string = "3"; // ตัวอย่างสถานะปัจจุบัน

  return (
    <Modal visible={visible} onCancel={onClose} footer={null} width={700}>
      <div className="order-summary-container">
        <header className="order-header">
          <h1>iGotSofa</h1>
        </header>

        <div className="order-progress">
          <div className={`step ${status >= "1" ? (status === "1" ? "active" : "completed") : ""}`}>
            1 เตรียมจัดส่ง
          </div>
          <div className={`step ${status >= "2" ? (status === "2" ? "active" : "completed") : ""}`}>
            2 กำลังส่ง
          </div>
          <div className={`step ${status >= "3" ? (status === "3" ? "active" : "completed") : ""}`}>
            3 เสร็จสมบูรณ์
          </div>
        </div>

        <div className="order-details">
          <h2>รายละเอียดการสั่งซื้อ</h2>

          <div className="details-section">
            <h3>ข้อมูลคำสั่งซื้อ</h3>
            <div className="detail-row">
              <span>ยอดรวม</span>
              <span>{orderDetails.total.toFixed(2)}</span>
            </div>
            <div className="detail-row">
              <span>ค่าจัดส่ง</span>
              <span>{orderDetails.shippingFee.toFixed(2)}</span>
            </div>
            <div className="detail-row">
              <span>ส่วนลดค่าจัดส่ง</span>
              <span>{orderDetails.discount.toFixed(2)}</span>
            </div>
            <div className="detail-row">
              <span>เเ\u0007ดมูล</span>
              <span>{orderDetails.vat.toFixed(2)}</span>
            </div>
            <div className="detail-row total">
              <span>รวมทั้งสิ้น</span>
              <span>{orderDetails.grandTotal.toFixed(2)}</span>
            </div>
          </div>

          <div className="order-info">
            <div className="info-row">
              <span>หมายเลขคำสั่งซื้อ</span>
              <span className="highlight">{orderDetails.orderNumber}</span>
            </div>
            <div className="info-row">
              <span>วันที่สั่งซื้อ</span>
              <span>{orderDetails.orderDate}</span>
            </div>
            <div className="info-row">
              <span>วันที่จัดส่ง</span>
              <span>{orderDetails.deliveryDate}</span>
            </div>
          </div>
        </div>
      </div>
    </Modal>
  );
};

export default OrderSummary;
