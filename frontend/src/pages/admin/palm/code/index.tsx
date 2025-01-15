import { useState, useEffect } from "react";
import { Space, Table, Button, Col, Row, Divider, Modal, message } from "antd";
import { PlusOutlined, EditOutlined, DeleteOutlined } from "@ant-design/icons";
import type { ColumnsType } from "antd/es/table";
import { GetCodes, DeleteCodeById } from "../../../../services/https";
import { CodeInterface } from "../../../../interfaces/Code";
import { Link, useNavigate } from "react-router-dom";
import dayjs from "dayjs";

function Codes() {
  const columns: ColumnsType<CodeInterface> = [
    {
      title: "ลำดับ",
      dataIndex: "ID",
      key: "id",
    },
    {
      title: "รูปโค้ด",
      dataIndex: "code_picture",
      key: "code_picture",
      width: "15%",
      render: (text, record, index) => (
        <img src={record.code_picture} className="w3-left w3-circle w3-margin-right" width="100%" />
      )
    },
    {
      title: "ชื่อโค้ด",
      dataIndex: "code_topic",
      key: "code_topic",
      width: 50,
    },
    {
      title: "คำอธิบายโค้ด",
      dataIndex: "code_description",
      key: "code_description",
      width: 300,
    },
    {
      title: "ส่วนลด(บาท)",
      dataIndex: "discount",
      key: "discount",
    },
    {
      title: "จำนวน",
      dataIndex: "quantity",
      key: "quantity",
    },
    {
      title: "ขั้นต่ำ",
      dataIndex: "minimum",
      key: "minimum",
    },
    {
      title: "วันเริ่มต้นโค้ด",
      dataIndex: "date_start",
      key: "date_start",
      width: 250,
      render: (record) => <p>{dayjs(record).format("dddd DD MMM YYYY")}</p>,
    },
    {
      title: "วันสิ้นสุดโค้ด",
      dataIndex: "date_end",
      key: "date_end",
      width: 250,
      render: (record) => <p>{dayjs(record).format("dddd DD MMM YYYY")}</p>,
    },
    {
      title: "จัดการ",
      dataIndex: "Manage",
      key: "manage",
      width: 200,
      render: (text, record, index) => (
        <>
          <Button
            onClick={() => navigate(`/code/edit/${record.ID}`)}
            shape="circle"
            icon={<EditOutlined />}
            size={"large"}
          />
          <Button
            onClick={() => showModal(record)}
            style={{ marginLeft: 10 }}
            shape="circle"
            icon={<DeleteOutlined />}
            size={"large"}
            danger
          />
        </>
      ),
    },
  ];

  const navigate = useNavigate();

  const [codes, setCodes] = useState<CodeInterface[]>([]);

  const [messageApi, contextHolder] = message.useMessage();

  // Model
  const [open, setOpen] = useState(false);
  const [confirmLoading, setConfirmLoading] = useState(false);
  const [modalText, setModalText] = useState<String>();
  const [deleteId, setDeleteId] = useState<Number>();

  const getCodes = async () => {
    try {
      let res = await GetCodes();
      console.log("GetCodes Response: ", res.data); // ตรวจสอบข้อมูลที่ได้จาก API
      if (res) {
        setCodes(res.data);
      }
    } catch (error) {
      console.error("Error fetching codes: ", error); // ดูว่ามี error อะไร
    }
  };
  
  const showModal = (val: CodeInterface) => {
    setModalText(
      `คุณต้องการลบข้อมูลโค้ด "${val.code_topic} " หรือไม่ ?`
    );
    setDeleteId(val.ID);
    setOpen(true);
  };

  const handleOk = async () => {

    setConfirmLoading(true);
    let res = await DeleteCodeById(deleteId);
    if (res) {
      setOpen(false);
      messageApi.open({
        type: "success",
        content: "ลบข้อมูลสำเร็จ",
      });
      getCodes();
    } else {
      setOpen(false);
      messageApi.open({
        type: "error",
        content: "เกิดข้อผิดพลาด !",
      });
    }
    setConfirmLoading(false);
  };

  const handleCancel = () => {
    setOpen(false);
  };

  useEffect(() => {
    getCodes();
  }, []);
  
  useEffect(() => {
    console.log("Current codes:", codes);
  }, [codes]);
  

  return (
  <>
    {contextHolder}
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        justifyContent: "center",
        alignItems: "center",
        minHeight: "100vh", // ทำให้ container ครอบคลุมเต็มจอ
        background: "#f0f2f5", // เพิ่มสีพื้นหลังเพื่อแยกชั้น
        padding: "20px", // เพิ่มระยะห่าง
        boxSizing: "border-box",
      }}
    >
      <div style={{ width: "100%", maxWidth: "1200px" }}>
        {/* Header Section */}
        <Row
          style={{
            background: "#ffffff",
            padding: "15px 20px",
            borderRadius: "15px",
            boxShadow: "0 2px 8px rgba(0,0,0,0.1)",
          }}
        >
          <Col span={12}>
            <h2 style={{ margin: 0, color: "#333" }}>จัดการข้อมูลโค้ดส่วนลด</h2>
          </Col>
          <Col span={12} style={{ textAlign: "end", alignSelf: "center" }}>
            <Space>
              <Link to="/code/create">
                <Button
                  type="primary"
                  icon={<PlusOutlined />}
                  style={{
                    backgroundColor: "#1890ff",
                    borderColor: "#1890ff",
                    color: "#fff",
                  }}
                >
                  เพิ่มโค้ด
                </Button>
              </Link>
            </Space>
          </Col>
        </Row>

        {/* Divider */}
        <Divider style={{ margin: "20px 0" }} />

        {/* Table Section */}
        <div
          style={{
            marginTop: 20,
            background: "#ffffff",
            padding: "20px",
            borderRadius: "15px", // ลดความโค้งของขอบเล็กน้อย
            boxShadow: "0 2px 8px rgba(0,0,0,0.1)",
          }}
        >
          <Table
            rowKey="ID"
            columns={columns}
            dataSource={Array.isArray(codes) ? codes : []}
            style={{ background: "#ffffff" }}
            size="middle"
          />
        </div>

        {/* Modal Section */}
        <Modal
          title={
            <span style={{ fontWeight: "bold", color: "#ff4d4f" }}>
              ลบข้อมูล ?
            </span>
          }
          open={open}
          onOk={handleOk}
          confirmLoading={confirmLoading}
          onCancel={handleCancel}
          okText="ยืนยัน"
          cancelText="ยกเลิก"
          okButtonProps={{ danger: true }}
        >
          <p style={{ fontSize: "16px", color: "#555" }}>{modalText}</p>
        </Modal>
      </div>
    </div>
  </>
);

  
}

export default Codes;

