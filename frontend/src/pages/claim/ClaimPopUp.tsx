import * as React from 'react';
import { Modal, Button, Typography, Input, Select, Upload } from "antd";
import { PlusOutlined } from "@ant-design/icons";
import "./ClaimPopup.css";

const { TextArea } = Input;
const { Option } = Select;
const { Title, Text } = Typography;

interface ClaimPopupProps {
  visible: boolean;
  product: {
    name: string;
    image: string;
    quantity: number;
    price: number;
    orderId: string;
  } | null;
  onCancel: () => void;
  onSubmit: () => void;
}

const ClaimPopup: React.FC<ClaimPopupProps> = ({ visible, product, onCancel, onSubmit }) => {
  return (
    <Modal
      visible={visible}
      onCancel={onCancel}
      footer={null}
      className="claim-popup"
    >
      <div className="popup-content">
        <Title level={3}>Claim</Title>
        {product && (
          <div className="product-details">
            <img src={product.image} alt={product.name} className="product-image" />
            <div className="product-info">
              <Text>{product.name}</Text>
              <Text>จำนวน: {product.quantity}</Text>
              <Text>{product.price.toFixed(2)} THB</Text>
            </div>
          </div>
        )}
        <Text>หมายเลขคำสั่งซื้อ: {product?.orderId}</Text>
        <div className="claim-form">
          <Select placeholder="เลือกเหตุผล" className="claim-select">
            <Option value="reason1">เหตุผลที่ 1</Option>
            <Option value="reason2">เหตุผลที่ 2</Option>
            <Option value="reason3">เหตุผลที่ 3</Option>
          </Select>
          <TextArea rows={4} placeholder="กรอกรายละเอียด" className="claim-textarea" />
          <Upload className="claim-upload">
            <Button icon={<PlusOutlined />}>เพิ่มรูปภาพ</Button>
          </Upload>
        </div>
        <Button type="primary" block onClick={onSubmit} className="claim-submit">
          ส่ง
        </Button>
      </div>
    </Modal>
  );
};

export default ClaimPopup;