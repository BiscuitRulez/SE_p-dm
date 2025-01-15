import { useState } from "react";
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
import { PlusOutlined } from "@ant-design/icons";
import { useNavigate } from "react-router-dom";
import ImgCrop from "antd-img-crop";
import { Upload, UploadFile, UploadProps } from "antd";
import dayjs from "dayjs"; // ใช้ Day.js
import { CreateCode } from "../../../../../services/https";
import { CodeInterface } from "../../../../../interfaces/Code";

function CodeCreate() {
  const navigate = useNavigate();
  const [messageApi, contextHolder] = message.useMessage();
  const [fileList, setFileList] = useState<UploadFile[]>([]);

  const handleCancel = () => {
    navigate(-1);
  };

  const onChange: UploadProps["onChange"] = ({ fileList: newFileList }) => {
    setFileList(newFileList);
  };

  const onFinish = async (values: CodeInterface) => {

    console.log("Date Start (Raw):", values.date_start);
    console.log("Date End (Raw):", values.date_end);

    const payload: CodeInterface = {
      ...values,
      date_start: values.date_start
        ? dayjs(values.date_start).format("YYYY-MM-DDTHH:mm:ss[Z]")
        : undefined, // ใช้ Day.js
      date_end: values.date_end
        ? dayjs(values.date_end).format("YYYY-MM-DDTHH:mm:ss[Z]")
        : undefined,
      discount: Number(values.discount),
      quantity: Number(values.quantity),
      minimum: Number(values.minimum),
      code_picture: fileList[0]?.thumbUrl || "",
    };

    console.log("Payload:", payload);

    const res = await CreateCode(payload);
    if (res) {
      messageApi.open({
        type: "success",
        content: "บันทึกข้อมูลโค้ดสำเร็จ",
      });
      setTimeout(() => {
        navigate("/code");
      }, 2000);
    } else {
      messageApi.open({
        type: "error",
        content: "เกิดข้อผิดพลาด !",
      });
    }
  };

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
        <h2>เพิ่มข้อมูลโค้ดส่วนลด</h2>
        <Divider />
        <Form name="basic" layout="vertical" onFinish={onFinish}>
          <Row gutter={[16, 16]}>
            <Col xs={24}>
              <Form.Item
                label="อัพโหลดรูปโค้ด"
                name="code_picture"
                valuePropName="fileList"
              >
                <ImgCrop rotationSlider>
                  <Upload
                    id={"code_picture"}
                    fileList={fileList}
                    onChange={onChange}
                    beforeUpload={(file) => {
                      setFileList([...fileList, file]);
                      return false;
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

export default CodeCreate;
