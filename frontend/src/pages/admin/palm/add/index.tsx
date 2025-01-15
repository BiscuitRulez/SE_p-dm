import { useState, useEffect } from "react";

import { Space, Table, Button, Col, Row, Divider, message } from "antd";

import { PlusOutlined, DeleteOutlined } from "@ant-design/icons";

import type { ColumnsType } from "antd/es/table";

import { DeleteUsersById, GetAdmin } from "../../../../services/https/index";

import { UserInterface } from "../../../../interfaces/User";

import { Link, useNavigate } from "react-router-dom";

import dayjs from "dayjs";


function AddAdmin() {

    const navigate = useNavigate();

    const [users, setUsers] = useState<UserInterface[]>([]);

    const [messageApi, contextHolder] = message.useMessage();

    const myId = localStorage.getItem("id");


    const columns: ColumnsType<UserInterface> = [
        {
            title: "ลำดับ",
            dataIndex: "ID",
            key: "id",
        },
        {
            title: "ชื่อ",
            dataIndex: "first_name",
            key: "first_name",
        },
        {
            title: "นามสกุล",
            dataIndex: "last_name",
            key: "last_name",
        },
        {
            title: "อีเมล",
            dataIndex: "email",
            key: "email",
        },
        {
            title: "วัน/เดือน/ปี เกิด",
            key: "birthday",
            render: (record) => <>{dayjs(record.birthday).format("DD/MM/YYYY")}</>,
        },
        {
            title: "จัดการ",
            render: (record) => (
                <Space>
                    {/* ปุ่มลบจะไม่แสดงเมื่อ ID เท่ากับ myId */}
                    {myId !== record?.ID && (
                        <Button
                            type="dashed"
                            danger
                            icon={<DeleteOutlined />}
                            onClick={() => deleteUserById(record.ID)}
                        >
                            ลบ
                        </Button>
                    )}
                    <Button
                        type="primary"
                        icon={<DeleteOutlined />}
                        onClick={() => navigate(`/add/edit/${record.ID}`)}
                    >
                        แก้ไขข้อมูล
                    </Button>
                </Space>
            ),
        },
    ];




    const deleteUserById = async (id: string) => {

        let res = await DeleteUsersById(id);


        if (res.status == 200) {

            messageApi.open({

                type: "success",

                content: res.data.message,

            });

            await getAdmins();

        } else {

            messageApi.open({

                type: "error",

                content: res.data.error,

            });

        }

    };


    const getAdmins = async () => {

        let res = await GetAdmin();



        if (res.status == 200) {

            setUsers(res.data);

        } else {

            setUsers([]);

            messageApi.open({

                type: "error",

                content: res.data.error,

            });

        }

    };


    useEffect(() => {

        getAdmins();

    }, []);


    return (

        <>

            {contextHolder}

            <Row>

                <Col span={12}>

                    <h2>จัดการข้อมูล ADMIN</h2>

                </Col>


                <Col span={12} style={{ textAlign: "end", alignSelf: "center" }}>

                    <Space>

                        <Link to="/add/create">

                            <Button type="primary" icon={<PlusOutlined />}>

                                สร้างข้อมูล

                            </Button>

                        </Link>

                    </Space>

                </Col>

            </Row>


            <Divider />


            <div style={{ marginTop: 20 }}>

                <Table

                    rowKey="ID"

                    columns={columns}

                    dataSource={users}

                    style={{ width: "100%", overflow: "scroll" }}

                />

            </div>

        </>

    );

}


export default AddAdmin;