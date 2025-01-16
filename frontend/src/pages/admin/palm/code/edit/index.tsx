import { useState, useEffect } from "react";
import {
  Space,
  Button,
  Col,
  Row,
  Divider,
  Form,
  Input,
  Card,
  message,
  DatePicker,
} from "antd";
import dayjs from "dayjs";
import ImgCrop from "antd-img-crop";
import { Upload, UploadFile, UploadProps } from "antd";
import { PlusOutlined } from "@ant-design/icons";
import { CodeInterface } from "../../../../../interfaces/Code";
import { GetCodesById, UpdateCode } from "../../../../../services/https";
import { useNavigate, useParams } from "react-router-dom";


function CodeEdit() {

  const navigate = useNavigate();
  const [messageApi, contextHolder] = message.useMessage();
  const [fileList, setFileList] = useState<UploadFile[]>([]);
  const [code, setCode] = useState<CodeInterface>();

  // รับข้อมูลจาก params
  let { id } = useParams();
  // อ้างอิง form กรอกข้อมูล
  const [form] = Form.useForm();

  const handleCancel = () => {
    navigate(-1);
  };
  const onChange: UploadProps["onChange"] = ({ fileList: newFileList }) => {
    setFileList(newFileList); // อัปเดตสถานะ fileList

    // หากต้องการเก็บ Base64 ของรูปภาพใหม่
    if (newFileList.length > 0 && newFileList[0].originFileObj) {
      const reader = new FileReader();
      reader.onload = () => {
        form.setFieldValue("code_picture", reader.result); // เก็บ Base64 ในฟอร์ม
      };
      reader.readAsDataURL(newFileList[0].originFileObj);
    }
  };


  const onFinish = async (values: CodeInterface) => {
    values.ID = code?.ID;
    values.quantity = Number(values.quantity || 0);
    values.discount = Number(values.discount || 0);
    values.minimum = Number(values.minimum || 0);

    if (fileList.length > 0 && fileList[0].originFileObj) {
      const reader = new FileReader();
      reader.onload = async () => {
        values.code_picture = reader.result as string; // เก็บ Base64
        let res = await UpdateCode(values);
        if (res.status) {
          messageApi.open({
            type: "success",
            content: res.message,
          });
          setTimeout(() => navigate("/code"), 2000);
        } else {
          messageApi.open({
            type: "error",
            content: res.message,
          });
        }
      };
      reader.readAsDataURL(fileList[0].originFileObj);
    } else {
      // กรณีไม่มีการเปลี่ยนรูป
      let res = await UpdateCode(values);
      if (res.status) {
        messageApi.open({
          type: "success",
          content: res.message,
        });
        setTimeout(() => navigate("/code"), 2000);
      } else {
        messageApi.open({
          type: "error",
          content: res.message,
        });
      }
    }
  };


  const getCodeById = async () => {
    let res = await GetCodesById(Number(id));
    console.log("Response from GetCodesById:", res); // ตรวจสอบโครงสร้างข้อมูล

    if (res.status && Array.isArray(res.data) && res.data.length > 0) {
      const codeData = res.data[0]; // เข้าถึงข้อมูลที่ index 0
      setCode(codeData); // เก็บข้อมูลใน state

      console.log("Fetched Code Data:", codeData); // ตรวจสอบค่าข้อมูล

      // ตั้งค่าให้กับฟอร์ม
      form.setFieldsValue({
        code_topic: codeData.code_topic,
        code_description: codeData.code_description,
        discount: codeData.discount,
        quantity: codeData.quantity,
        minimum: codeData.minimum,
        date_start: dayjs(codeData.date_start),
        date_end: dayjs(codeData.date_end),
      });

      // ตั้งค่า fileList สำหรับการแสดงรูปภาพเดิม
      setFileList([
        {
          uid: "-1",
          name: "Existing Image",
          status: "done",
          url: codeData.code_picture, // URL หรือ Base64
        },
      ]);

    } else {
      console.error("Failed to fetch code or no data available.");
    }
  };


  useEffect(() => {
    getCodeById(); // เรียกใช้งานฟังก์ชันเพื่อดึงข้อมูล
  }, []);

  return (
    <div
      style={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        minHeight: "100vh", // already correct for full viewport height
        width: "100vw",     // add full viewport width
        padding: "0",       // remove padding to maximize space
        background: "white",
        margin: 0,          // remove any default margins
        boxSizing: "border-box" // ensure padding and border are included in width/height
      }}
    >
      {contextHolder}
      <Card>
        <h2> แก้ไขโค้ดส่วนลด</h2>
        <Divider />
        {/* Make sure form is connected */}
        <Form
          name="basic"
          form={form}
          layout="vertical"
          onFinish={onFinish}
          autoComplete="off"
        >
          <Row gutter={[16, 16]}>
            <Col xs={24}>
              <Form.Item
                label="อัพโหลดรูปโค้ด"
                name="code_picture"
                valuePropName="fileList"
              >
                <ImgCrop rotationSlider>
                  <Upload
                    fileList={fileList}
                    onChange={onChange}
                    beforeUpload={(file) => {
                      // อ่านไฟล์เป็น Base64 และอัปเดตฟอร์ม
                      const reader = new FileReader();
                      reader.onload = () => {
                        form.setFieldValue("code_picture", reader.result); // เก็บ Base64 ในฟอร์ม
                      };
                      reader.readAsDataURL(file);
                      return false; // ป้องกันการอัปโหลดโดยตรง
                    }}
                    maxCount={1}
                    multiple={false}
                    listType="picture-card"
                  >
                    <div>
                      <PlusOutlined />
                      <div style={{ marginTop: 8 }}>อัพโหลด</div>
                    </div>
                  </Upload>

                </ImgCrop>
              </Form.Item>
            </Col>
            <Col xs={24}>
              <Form.Item
                label="TOPIC"
                name="code_topic"
                rules={[{ required: true, message: "กรุณากรอกหัวข้อโค้ด !" }]}
              >
                <Input />
              </Form.Item>
            </Col>
            <Col xs={24}>
              <Form.Item
                label="DESCRIPTION"
                name="code_description"
                rules={[
                  { required: true, message: "กรุณากรอกคำอธิบาย !" },
                ]}
              >
                <Input.TextArea rows={4} />
              </Form.Item>
            </Col>
            <Col xs={24}>
              <Form.Item
                label="DISCOUNT"
                name="discount"
                rules={[{ required: true, message: "กรุณากรอกส่วนลด !" }]}
              >
                <Input />
              </Form.Item>
            </Col>
            <Col xs={24}>
              <Form.Item
                label="QUANTITY"
                name="quantity"
                rules={[{ required: true, message: "กรุณากรอกจำนวน !" }]}
              >
                <Input />
              </Form.Item>
            </Col>
            <Col xs={24}>
              <Form.Item
                label="MINIMUM"
                name="minimum"
                rules={[{ required: true, message: "กรุณากรอกราคาขั้นต่ำ !" }]}
              >
                <Input />
              </Form.Item>
            </Col>
            <Col xs={24} lg={12}>
              <Form.Item label="วันเริ่มต้นโค้ด" name="date_start">
                <DatePicker style={{ width: "100%" }} />
              </Form.Item>
            </Col>
            <Col xs={24} lg={12}>
              <Form.Item label="วันสิ้นสุดโค้ด" name="date_end">
                <DatePicker style={{ width: "100%" }} />
              </Form.Item>
            </Col>
          </Row>
          <Row justify="end">
            <Col>
              <Form.Item>
                <Space>
                  <Button onClick={handleCancel}>ยกเลิก</Button>
                  <Button type="primary" htmlType="submit">
                    ยืนยัน
                  </Button>
                </Space>
              </Form.Item>
            </Col>
          </Row>
        </Form>
      </Card>
    </div>
  );
}

export default CodeEdit;