import { Form, Input, Button, Card, Divider, Row, Col, message } from "antd";
import { useNavigate } from "react-router-dom";
import { AddressInterface } from "../../../../../interfaces/Address";
import { CreateAddress } from "../../../../../services/https"; // Import the AddAddress API function

const AddAddressPage = () => {
    const [form] = Form.useForm();
    const navigate = useNavigate();
    const [messageApi, contextHolder] = message.useMessage();

    const onFinish = async (values: AddressInterface) => {
        try {
            const userId = localStorage.getItem("id"); // Assume the user ID is stored in localStorage
            if (!userId) {
                throw new Error("User ID not found.");
            }

            const payload = {
                user_id: parseInt(userId, 10), // Convert to number
                ...values,
            };

            const res = await CreateAddress(payload);
            if (res.status === 201) {
                messageApi.success("เพิ่มที่อยู่สำเร็จ!");
                setTimeout(() => {
                    navigate("/profile");
                }, 1000);
            } else {
                throw new Error(res.data.error || "การเพิ่มที่อยู่ล้มเหลว");
            }
        } catch (error: any) {
            messageApi.error(error.message || "เกิดข้อผิดพลาด");
        }
    };

    return (
        <div>
            {contextHolder}

            <Card>
                <h2>เพิ่มที่อยู่ใหม่</h2>
                <Divider />

                <Form
                    form={form}
                    layout="vertical"
                    onFinish={onFinish}
                    autoComplete="off"
                >
                    <Row gutter={[16, 16]}>
                        <Col span={24}>
                            <Form.Item
                                label="ที่อยู่เต็ม"
                                name="full_address"
                                rules={[{ required: true, message: "กรุณากรอกที่อยู่เต็ม!" }]}
                            >
                                <Input.TextArea rows={3} />
                            </Form.Item>
                        </Col>

                        <Col span={12}>
                            <Form.Item
                                label="เมือง"
                                name="city"
                                rules={[{ required: true, message: "กรุณากรอกชื่อเมือง!" }]}
                            >
                                <Input />
                            </Form.Item>
                        </Col>

                        <Col span={12}>
                            <Form.Item
                                label="จังหวัด"
                                name="province"
                                rules={[{ required: true, message: "กรุณากรอกชื่อจังหวัด!" }]}
                            >
                                <Input />
                            </Form.Item>
                        </Col>

                        <Col span={12}>
                            <Form.Item
                                label="รหัสไปรษณีย์"
                                name="postal_code"
                                rules={[{ required: true, message: "กรุณากรอกรหัสไปรษณีย์!" }]}
                            >
                                <Input />
                            </Form.Item>
                        </Col>
                    </Row>

                    <Row justify="end" style={{ marginTop: "16px" }}>
                        <Col style={{ textAlign: "right" }}>
                            <Button
                                type="default"
                                onClick={() => navigate("/profile")}
                                style={{ marginRight: "10px" }} // Add space between buttons
                            >
                                ย้อนกลับ
                            </Button>
                            <Button
                                type="primary"
                                htmlType="submit"
                            >
                                บันทึกที่อยู่
                            </Button>
                        </Col>
                    </Row>


                </Form>
            </Card>
        </div>
    );
};

export default AddAddressPage;
