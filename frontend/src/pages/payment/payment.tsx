import React, { useEffect, useState } from 'react';
import '../../pages/payment/payment.css';
import {
  Layout, Card, List, Select, InputNumber, Button, Typography, Divider, message, Modal, Upload,
} from 'antd';
import { QRCode } from 'antd';
import generatePayload from 'promptpay-qr';
// controller
import { GetPaymentMethod, GetCollectedCodesToShow, GetAddressesByUserId, GetAllShipping, CreatePayment } from '../../services/https';
import { AddressInterface } from '../../interfaces/Address';
import { ShippingInterface } from "../../interfaces/Shipping"

const { Content } = Layout;
const { Title, Text } = Typography;
const { Option } = Select;

const PaymentPage: React.FC = () => {
  const [selectedLocation, setSelectedLocation] = useState<string | undefined>(undefined);
  const [selectedPayment, setSelectedPayment] = useState<string | undefined>(undefined);
  const [selectedDiscount, setSelectedDiscount] = useState<string | undefined>(undefined);
  const [selectedShippingCompany, setSelectedShippingCompany] = useState<string | undefined>(undefined);
  const [pointsToUse, setPointsToUse] = useState<number>(1000);
  const [qrCode, setQrCode] = useState<string>('');
  const [isModalVisible, setIsModalVisible] = useState<boolean>(false);
  const [paymentMethod, setPaymentMethod] = useState<any[]>([]);
  // const [adress, setAdress] = useState<AddressInterface | null>(null);
  const [codeCollectToshow, setcodeCollectToshow] = useState<any[]>([]);
  const [selectedCode, setSelectedCode] = useState<number>(1);
  const [shipping, setShipping] = useState<ShippingInterface[]>([]);

  const handleCodeChange = (value: number) => {
    console.log("Selected Code:", value);
    setSelectedCode(value);
  };

  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState("");


  const userId = localStorage.getItem("id") || "";
  console.log("-------------------userId---------------------", userId);

  const items = [
    { name: 'สินค้า1', quantity: 1, price: 2 },
    { name: 'สินค้า2', quantity: 2, price: 5000 },
    { name: 'สินค้า3', quantity: 3, price: 1 },
  ];

  const totalPoints = 10000;
  const pointsDiscount = pointsToUse / 100;
  const totalAmount = items.reduce((sum, item) => sum + item.price * item.quantity, 0);

  ////////// discountAmount คือ ส่วนลดที่ได้จากโค้ด
  const discountAmount = codeCollectToshow.find((code) => code.code === selectedCode)?.discount || 0;
  console.log("discountAmount:", discountAmount);


  const shippingFee = selectedDiscount === 'FREEDELIVERY' ? 0 : 45;
  const totalPayable = (totalAmount - discountAmount - pointsDiscount + shippingFee).toFixed(2);
  // const [address, setAddress] = useState<AddressInterface | null>(null);




  const [address, setAddress] = useState<AddressInterface[]>([]);
  

  

  const handleOrder = () => {
    if (!selectedPayment || !selectedShippingCompany || !selectedLocation) {
      message.error('ชำระเงินล้มเหลว');
      return;
    }

    const phoneNumber = '0934155151';
    const qrCodePayload = generatePayload(phoneNumber, { amount: parseFloat(totalPayable) });
    setQrCode(qrCodePayload);
    setIsModalVisible(true);

    message.success('ชำระเงินสำเร็จ');
  };

  const handleCancel = () => {
    setIsModalVisible(false);
  };

  const fetchPaymentMethod = async () => {
    try {
      const res = await GetPaymentMethod();
      console.log("-------------------fetchPaymentMethod----------------------", res);
      console.log("----------------------fetchPaymentMethod-------------------", res.data);
      if (res?.data && res.status === 200) {
        setPaymentMethod(res.data);
      } else {
        setError("Failed to load Payment Method");
      }
    } catch (error) {
      setError("Error fetching Payment Method. Please try again later.");
    } finally {
      setIsLoading(false);
    }
  };

  const getShipping = async () => {
    try {
      const res = await GetAllShipping();
      console.log("shipping response:", res);
      console.log("shipping data:", res.data);

      if (res?.data.shippings && Array.isArray(res.data.shippings)) {
        setShipping(res.data.shippings); // Only set if it's an array
      } else {
        setError("Invalid data format from API");
      }
    } catch (error) {
      setError("Error fetching shipping data");
    } finally {
      setIsLoading(false);
    }
  };



  // const fetchCollectedCodes = async () => {
  //   try {
  //     const res = await GetCollectedCodes(userId);
  //     if (res?.data && res.status === 200) {
  //       const formattedCodes = res.data.map((code: any) => ({
  //         code: code.ID,
  //         value: code.code_topic,
  //       }));
  //       setCodeCollect(formattedCodes);
  //     } else {
  //       setError("Failed to load collected discount codes.");
  //     }
  //   } catch (error) {
  //     setError("Error fetching collected discount codes. Please try again later.");
  //   }
  // };


  const fetchAddressData = async (userIdstr: string) => {
    try {
      const res = await GetAddressesByUserId(userIdstr);
      console.log("Address Response:", res.data.data);
      if (res.status === 200 && Array.isArray(res.data.data) && res.data.data.length > 0) {
        setAddress(res.data.data); // เก็บข้อมูลทั้งหมดเป็น array
        setSelectedLocation(res.data.data[0].ID);
      } else {
        setAddress([]);
        message.error("No Address Available!");
      }
    } catch (error) {
      console.error("Error fetching address:", error);
      setAddress([]);
      message.error("Your viewing address is not yet available.");
    }
  };


  const fetchCollectedCodesById = async (userId: string) => {
    try {
      const res = await GetCollectedCodesToShow(userId); // เรียก API
      console.log("-------------------fetchCollectedCodesById--------111111-------------", res.data);

      if (res?.data && res.status === 200) {
        // Total amount ของรายการทั้งหมด
        const totalAmount = items.reduce((sum, item) => sum + item.price * item.quantity, 0);

        // กรองเฉพาะโค้ดที่ totalAmount >= minimum
        const filteredCodes = res.data.filter((code: any) => totalAmount >= code.minimum);

        const formattedCodes = filteredCodes.map((code: any) => ({
          code: code.ID, // ใช้ ID เป็น key
          value: code.code_topic,
          minimum: code.minimum,
          discount: code.discount,
        }));

        console.log("formattedCodes (filtered):", formattedCodes);
        setcodeCollectToshow(formattedCodes); // อัปเดต state ด้วย array ที่กรองแล้ว
      } else {
        setError("Failed to load collected discount codes.");
      }
    } catch (error) {
      setError("Error fetching collected discount codes. Please try again later.");
    }
  };

  const handleSubmit= async () => {
    
    const dataPayment = {
      Date: new Date().toISOString(),
      UserID: userId,
      PaymentMethodID: Number(paymentMethod), // แปลงเป็น number
      PaymentStatusID: Number(),  // แปลงเป็น number
      PaymentMethod: selectedPayment,
      PaymentStatus: "Pending" // or any appropriate status
    };
  
    console.log("dataPayment:", dataPayment);
  
    // try {
    //     const response = await CreatePayment(userId, dataPayment); // เรียก API สร้าง Course
    //     console.log(dataPayment);
    //     console.log(response);
  
    //     if (response.status === 200 || response.status === 201) {
    //         console.log("ข้อมูล Course ถูกสร้างสำเร็จ");
    //     } else {
    //         console.error("เกิดข้อผิดพลาดในการสร้าง Course");
    //     }
    // } catch (error) {
    //     console.error("เกิดข้อผิดพลาดในการส่งข้อมูล Course:", error);
    // }
  };
  

  useEffect(() => {
    fetchPaymentMethod();
    getShipping();
    setIsLoading(true);
    // fetchCollectedCodes();
  }, []);

  useEffect(() => {
    if (userId) {
      fetchCollectedCodesById(userId); // ใช้ userId จาก localStorage
      fetchAddressData(userId)
    } else {
      message.error("The user ID was not found in localStorage.");
    }
  }, [userId]);

  return (
    <Layout>
      <Content className="layout-content">
        {/* Left Card */}
        <Card title="Payment Details" className="card-left">
          {/* Address Selection */}
          <div className="card-section">
            <Title level={5} className="card-section-title">ที่อยู่จัดส่ง</Title>
            <Select
              className="select-dropdown"
              value={selectedLocation}
              onChange={(value) => setSelectedLocation(value)}
            >
              {address?.map((item: AddressInterface) => (
                <Option key={item.ID} value={item.ID}>
                  {item.full_address}
                </Option>
              ))}

            </Select>
          </div>

          {/* Items List */}
          <List
            className="items-list"
            dataSource={items}
            renderItem={(item) => (
              <List.Item className="items-list-item">
                <div>
                  <Text className="items-list-item-text">{item.name}</Text>
                  <br />
                  <Text>จำนวน: {item.quantity}</Text>
                </div>
                <Text>{item.price * item.quantity} THB</Text>
              </List.Item>
            )}
            style={{ maxHeight: '200px', overflowY: 'auto', marginBottom: '10px' }}
          />

          {/* Payment Method */}
          <div>
            <div>
              <h5 className="card-section-title">เลือกวิธีการชำระเงิน</h5>
              <div className="radio-options">
                {paymentMethod.map((method) => (
                  <div className="radio-option" key={method.PaymentMethod}>
                    <input
                      id='paymentMethod'
                      type="radio"
                      value={method.PaymentMethod}
                      checked={selectedPayment === method.PaymentMethod}
                      onChange={(e) => setSelectedPayment(e.target.value)}
                    />
                    <label>{method.PaymentMethod}</label>
                  </div>
                ))}
              </div>
            </div>
          </div>

          <div className="card-section">
            <Title level={5} className="card-section-title">โค้ดส่วนลด</Title>
            <Select
              placeholder="เลือกโค้ดส่วนลด"
              className="select-dropdown"
              onChange={(value) => {
                handleCodeChange(value);
              }}
            >
              {codeCollectToshow.map((code) => (
                <Option key={code.code} value={code.code}>
                  {code.value} ส่วนลด {code.discount} บาท
                </Option>
              ))}
              <Option key="none" value={null}>
                -ไม่ใช้โค้ดส่วนลด-
              </Option>
            </Select>
          </div>

          <div className="card-section">
            <Title level={5} className="card-section-title">เลือกบริษัทขนส่ง</Title>
            <Select
              placeholder="เลือกบริษัทขนส่ง"
              className="select-dropdown"
              value={selectedShippingCompany}
              onChange={(value) => setSelectedShippingCompany(value)}
            >
              {shipping.map((item) => (
                <Option key={item.ID} value={item.ID}>
                  {item.Name}
                </Option>
              ))}
            </Select>
          </div>
        </Card>

        {/* Right Card */}
        <Card className="card-right">
          <div className="order-summary">
            <Text>ยอดรวมสินค้า: </Text>
            <Text strong className="order-summary-value">{totalAmount.toFixed(2)} THB</Text>
          </div>
          <div className="order-summary-2">
            <Text>คุณมี </Text>
            <Text strong className="order-summary-value-2">{totalPoints}</Text>
            <Text> คะแนน</Text>
          </div>
          <div className="order-summary-2">
            <Text>ใช้ </Text>
            <InputNumber
              min={0}
              max={totalPoints}
              value={pointsToUse}
              onChange={(value) => setPointsToUse(value || 0)}
            />
            <Text>&nbsp;&nbsp;คะแนน&nbsp;&nbsp;</Text>
            <Text className="spacing">ลด&nbsp;&nbsp;</Text>
            <Text strong>{pointsDiscount.toFixed(2)} THB</Text>
          </div>
          <div className="order-summary">
            <Text>ค่าจัดส่ง: </Text>
            <Text strong className="order-summary-value">{shippingFee.toFixed(2)} THB</Text>
          </div>
          {discountAmount > 0 && (
            <div className="order-summary">
              <Text>ส่วนลดพิเศษ: </Text>
              <Text strong style={{ color: 'red' }} className="order-summary-value">
                -{discountAmount.toFixed(2)} THB
              </Text>
            </div>
          )}
          {selectedDiscount === 'FREEDELIVERY' && (
            <div className="order-summary">
              <Text>ส่วนลดค่าจัดส่ง: </Text>
              <Text strong style={{ color: 'red' }}>
                -45.00 THB
              </Text>
            </div>
          )}
          <Divider className="order-summary-divider" />
          <div className="order-summary">
            <Text>ยอดชำระทั้งหมด: </Text>
            <Text strong>{totalPayable} THB</Text>
          </div>
          <Button
            type="primary"
            className="order-summary-button"
            onClick={handleOrder}
          >
            สั่งซื้อ
          </Button>
        </Card>

        {/* QR Code Modal */}
        {/* <Modal
          title="QR Code สำหรับการชำระเงิน"
          visible={isModalVisible}
          onCancel={handleCancel}
          footer={null}
        >
          {qrCode && (
            <QRCode
              type='svg'
              errorLevel="H"
              value={qrCode}
              style={{ width: '250px', height: '250px' }}
            />
          )}
        </Modal> */}
        <Modal
          title="QR Code สำหรับการชำระเงิน"
          visible={isModalVisible}
          onCancel={handleCancel}
          footer={null}
        >
          {qrCode ? (
            <>
              <QRCode
                type="svg"
                errorLevel="H"
                value={qrCode}
                style={{ width: '250px', height: '250px' }}
                id="qrCode"
              />
              <div style={{ textAlign: 'center', marginTop: '20px' }}>
                <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '20px' }}>
                  <Text strong >{totalPayable} THB</Text>
                  <button
                    onClick={() => {
                      const svgElement = document.getElementById('qrCode');
                      if (!svgElement) {
                        message.error('QR Code element not found');
                        return;
                      }
                      const svg = svgElement.outerHTML;
                      const blob = new Blob([svg], { type: 'image/svg+xml' });
                      const url = URL.createObjectURL(blob);
                      const a = document.createElement('a');
                      a.href = url;
                      a.download = 'qrcode.svg';
                      a.click();
                      URL.revokeObjectURL(url);
                    }}
                    style={{
                      padding: '10px 20px',
                      backgroundColor: '#1890ff',
                      color: '#fff',
                      border: 'none',
                      borderRadius: '5px',
                      cursor: 'pointer',
                    }}
                  >
                    ดาวน์โหลด QR Code
                  </button>
                  <Modal
  title="QR Code สำหรับการชำระเงิน"
  visible={isModalVisible}
  onCancel={handleCancel}
  footer={null}
>
  {qrCode ? (
    <>
      <QRCode
        type="svg"
        errorLevel="H"
        value={qrCode}
        style={{ width: '250px', height: '250px' }}
        id="qrCode"
      />
      <div style={{ textAlign: 'center', marginTop: '20px' }}>
        <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '20px' }}>
          <Text strong>{totalPayable} THB</Text>
          <button
            onClick={() => {
              const svgElement = document.getElementById('qrCode');
              if (!svgElement) {
                message.error('QR Code element not found');
                return;
              }
              const svg = svgElement.outerHTML;
              const blob = new Blob([svg], { type: 'image/svg+xml' });
              const url = URL.createObjectURL(blob);
              const a = document.createElement('a');
              a.href = url;
              a.download = 'qrcode.svg';
              a.click();
              URL.revokeObjectURL(url);
            }}
            style={{
              padding: '10px 20px',
              backgroundColor: '#1890ff',
              color: '#fff',
              border: 'none',
              borderRadius: '5px',
              cursor: 'pointer',
            }}
          >
            ดาวน์โหลด QR Code
          </button>
        </div>
      </div>

      {/* Upload Section */}
      <div style={{ textAlign: 'center', marginTop: '20px' }}>
        <Upload
          beforeUpload={() => false} // ป้องกันการอัปโหลดอัตโนมัติ
          listType="picture"
          maxCount={1} // อนุญาตให้อัปโหลดได้แค่ 1 รูป
          showUploadList={{
            showRemoveIcon: true,
          }}
          onChange={(info) => {
            if (info.file.status === 'removed') {
              message.info('File removed');
            } else {
              message.success('File selected');
            }
          }}
        >
          <Button>แสดงสลิป</Button>
        </Upload>
      </div>
            </>
          ) : (
            <div style={{ textAlign: 'center', margin: '20px 0' }}>
              กำลังโหลด QR Code...
            </div>
          )}
        </Modal>

                </div>
              </div>

            </>
          ) : (
            <div style={{ textAlign: 'center', margin: '20px 0' }}>
              กำลังโหลด QR Code...
            </div>
          )}
        </Modal>

      </Content>
    </Layout>
  );
};

export default PaymentPage;
