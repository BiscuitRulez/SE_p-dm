import { useEffect, useState } from "react";
import { Form, Input, Button, message, DatePicker, Row, Col, Card, Divider, Space, Table } from "antd";
import { PlusOutlined } from "@ant-design/icons";
import { Link, useNavigate } from "react-router-dom";
import type { ColumnsType } from "antd/es/table";
import dayjs from "dayjs";
import { GetUsersById, UpdateUsersById, GetAddressesByUserId } from "../../../../services/https/index";
import { UserInterface } from "../../../../interfaces/User";
import { AddressInterface } from "../../../../interfaces/Address";

const ProfileEdit = () => {
    const navigate = useNavigate();
    const id = localStorage.getItem("id");
    const [messageApi, contextHolder] = message.useMessage();
    const [form] = Form.useForm();
    const [addresses, setAddresses] = useState<AddressInterface[]>([]);
    const [loading, setLoading] = useState(true);

    const columns: ColumnsType<AddressInterface> = [
        {
            title: "ลำดับ",
            key: "index",
            render: (_text, _record, index) => index + 1, // ใช้ index ของ Array + 1
        },
        {
            title: "ที่อยู่เต็ม",
            dataIndex: "full_address",
            key: "full_address",
           
        },
        {
            title: "เมือง",
            dataIndex: "city",
            key: "city",
           
        },
        {
            title: "จังหวัด",
            dataIndex: "province",
            key: "province",
            
        },
        {
            title: "รหัสไปรษณีย์",
            dataIndex: "postal_code",
            key: "postal_code",
            
        },
    ];

    const getUserById = async (id: string) => {
        try {
            const res = await GetUsersById(id);
            if (res.status === 200) {
                form.setFieldsValue({
                    first_name: res.data.first_name,
                    last_name: res.data.last_name,
                    birthday: dayjs(res.data.birthday),
                });
            } else {
                throw new Error("ไม่พบข้อมูลผู้ใช้");
            }
        } catch (error) {
            messageApi.error("ไม่พบข้อมูลผู้ใช้");
            setTimeout(() => navigate("/customer"), 2000);
        }
    };

    const getUserAddresses = async (id: string) => {
        try {
            setLoading(true);
            const res = await GetAddressesByUserId(id);
            console.log("API Response:", res); // ตรวจสอบการตอบกลับ
    
            // ปรับตามโครงสร้างจริง
            if (res.status === 200 && Array.isArray(res.data.data)) {
                setAddresses(res.data.data); // หรือ res.data.addresses ตามโครงสร้างจริง
            } else {
                console.error("Unexpected Response Format:", res);
                throw new Error("ไม่สามารถดึงข้อมูลที่อยู่ได้");
            }
        } catch (error) {
            console.error("Error fetching addresses:", error);
            messageApi.error("ไม่สามารถดึงข้อมูลที่อยู่ได้");
        } finally {
            setLoading(false);
        }
    };
    
    

    const onFinish = async (values: UserInterface) => {
        if (!id) {
            messageApi.error("User ID is missing.");
            return;
        }

        try {
            const res = await UpdateUsersById(id, values);
            if (res.status === 200) {
                messageApi.success(res.data.message);
                setTimeout(() => navigate("/customer"), 2000);
            } else {
                throw new Error(res.data.error);
            }
        } catch (error) {
            messageApi.error(error instanceof Error ? error.message : "Failed to update user");
        }
    };

    useEffect(() => {
        if (!id) {
            messageApi.error("User ID not found.");
            setTimeout(() => navigate("/customer"), 2000);
            return;
        }

        getUserById(id);
        getUserAddresses(id);
    }, [id]);

    return (
        <div>
            {contextHolder}
            
            <Card>
                <h2>แก้ไขข้อมูลส่วนตัว</h2>
                <Divider />

                <Form
                    name="basic"
                    form={form}
                    layout="vertical"
                    onFinish={onFinish}
                    autoComplete="off"
                >
                    <Row gutter={[16, 0]}>
                        <Col xs={24} sm={24} md={24} lg={24} xl={12}>
                            <Form.Item
                                label="ชื่อจริง"
                                name="first_name"
                                rules={[{ required: true, message: "กรุณากรอกชื่อ !" }]}
                            >
                                <Input />
                            </Form.Item>
                        </Col>

                        <Col xs={24} sm={24} md={24} lg={24} xl={12}>
                            <Form.Item
                                label="นามสกุล"
                                name="last_name"
                                rules={[{ required: true, message: "กรุณากรอกนามสกุล !" }]}
                            >
                                <Input />
                            </Form.Item>
                        </Col>

                        <Col xs={24} sm={24} md={24} lg={24} xl={12}>
                            <Form.Item
                                label="วัน/เดือน/ปี เกิด"
                                name="birthday"
                                rules={[{ required: true, message: "กรุณาเลือกวัน/เดือน/ปี เกิด !" }]}
                            >
                                <DatePicker style={{ width: "100%" }} />
                            </Form.Item>
                        </Col>
                    </Row>

                    <Row justify="end">
                        <Col style={{ marginTop: "40px" }}>
                            <Form.Item>
                                <Space>
                                    <Link to="/customer">
                                        <Button htmlType="button">ยกเลิก</Button>
                                    </Link>
                                    <Button type="primary" htmlType="submit" icon={<PlusOutlined />}>
                                        บันทึก
                                    </Button>
                                </Space>
                            </Form.Item>
                        </Col>
                    </Row>
                </Form>
            </Card>

            <Card style={{ marginTop: "20px" }}>
                <h2>ที่อยู่</h2>
                <h4 style={{ fontSize: "14px", color: "#888" }}>
                    เพิ่มที่อยู่ของคุณเพื่อการชำระเงินที่ง่ายและเร็วขึ้น
                </h4>
                <Divider />

                <div style={{
                    marginTop: 20,
                    background: "#ffffff",
                    padding: "20px",
                    borderRadius: "15px",
                    boxShadow: "0 2px 8px rgba(0,0,0,0.1)",
                }}>
                    <Table
                        rowKey="ID"
                        columns={columns}
                        dataSource={addresses}
                        loading={loading}
                        style={{ background: "#ffffff" }}
                        size="middle"
                        locale={{ emptyText: "ไม่พบข้อมูลที่อยู่" }}
                    />
                </div>

                <Row justify="end">
                    <Col>
                        <Link to="/profile/address">
                            <Button type="default" style={{ marginTop: "20px" }}>
                                เพิ่มที่อยู่ใหม่
                            </Button>
                        </Link>
                    </Col>
                </Row>
            </Card>
        </div>
    );
};

export default ProfileEdit;
