import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import styled from 'styled-components';
import back from "../../../assets/back.png";
import axios from 'axios';
import '../../../components/card2.css';
import { GetProductByID, DeleteProduct, UpdateProductByID, GetProduct, GetCatagory } from "../../../services/https";
import { Product } from '../../../interfaces/Product';
import { CatagoryInterface } from '../../../interfaces/Catagory';
import { Typography, Modal, Form, Upload, message, Button, Input, Select, UploadFile, UploadProps } from "antd";

const EditProduct: React.FC = () => {
    const [ProductDetails, setProductDetails] = useState<Product | null>(null);
    const [products, setProducts] = useState<any[]>([]);
    const navigate = useNavigate();
    const [isEditModalVisible, setIsEditModalVisible] = useState(false);
    const [selectedProduct, setSelectedProduct] = useState<Product | null>(null);
    const [catagory, setCatagory] = useState<CatagoryInterface[]>([]);
    const [isLoading, setIsLoading] = useState(true);
    const [selectedCatagoryID, setSelectedCatagoryID] = useState<number | null>(null);
    const [fileList, setFileList] = useState<UploadFile[]>([]);
    const [form] = Form.useForm();

    // Fetch all products
    useEffect(() => {
        axios
            .get("http://localhost:8000/product")
            .then((response) => {
                setProducts(response.data.products || []);
            })
            .catch((error) => {
                console.error("Error fetching products:", error);
            });
    }, []);

    // Fetch categories
    useEffect(() => {
        const fetchCatagoryData = async () => {
            try {
                const res = await GetCatagory();
                if (res?.data) {
                    setCatagory(res.data.catagorys);
                } else {
                    message.error("Failed to load categories.");
                }
            } catch (err) {
                message.error("Error fetching categories.");
            } finally {
                setIsLoading(false);
            }
        };
        fetchCatagoryData();
    }, []);

    // Open edit modal (แก้ไขฟังก์ชันนี้)
    const handleEditClick = (product: Product) => {
        console.log("Sending product to AddProduct:", product); // ตรวจสอบค่าก่อนส่ง
        navigate("/AddProduct", { state: { product } });
      };
    // Close edit modal
    const closeEditModal = () => {
        setIsEditModalVisible(false);
        setFileList([]);
    };

    // Handle product update
    const handleEditSubmit = async (values: Product) => {
        try {
            // Upload image if a new file is selected
            let avatarUrl = values.image;
            if (fileList[0]?.originFileObj) {
                avatarUrl = await uploadImage(fileList[0]);
            }

            const updatedProduct = {
                ...values,
                Image: avatarUrl,
                CatagoryID: selectedCatagoryID || values.CatagoryID,
            };

            const res = await UpdateProductByID(String(selectedProduct?.ID), updatedProduct);
            if (res && res.status === 200) {
                message.success("Product updated successfully.");
                setIsEditModalVisible(false);
                window.location.reload();
            } else {
                message.error("Failed to update product.");
            }
        } catch (err) {
            console.error("Error updating product:", err);
            message.error("Error updating product. Please try again.");
        }
    };

    // Upload image to server
    const uploadImage = async (file: UploadFile): Promise<string> => {
        const formData = new FormData();
        formData.append("file", file.originFileObj as File);

        const res = await axios.post("http://localhost:8000/upload", formData, {
            headers: { "Content-Type": "multipart/form-data" },
        });

        return res.data.url; // Assuming the server responds with { url: "uploaded-image-url" }
    };

    return (
        <div>
          
            <a href="/" style={{ position: "absolute", top: "100px", right: "1400px" }}>
                <img style={{ width: "50px", height: "auto" }} src={back} />
            </a>

            <h1 style={{ marginTop: "290px", textAlign: "center" }}>กรุณาเลือกสินค้า</h1>
            <div className="product-grid2">
                {products.map((product) => (
                    <button key={product.ID} className="product-card2">
                        <img className="product-card-img2" src={product.Image} alt={product.Name} />
                        <div className="product-content2">
                            <h2 className="product-title2">{product.Name}</h2>
                            <p className="product-description2">{product.Description}</p>
                            <div style={{ display: "flex", gap:"30px", marginLeft:"120px", marginTop:"30px"}}>
                                <button
                                    style={{ backgroundColor: "#f4f4f4", color: "#001529", padding: "5px 10px", borderRadius: "4px" }}
                                    onClick={() => handleEditClick(product)}
                                >
                                    แก้ไข
                                </button>
                                <button
                                    style={{ backgroundColor: "#f44336", color: "white", padding: "5px 10px", borderRadius: "4px" }}
                                    onClick={() => DeleteProduct(product.ID)}
                                >
                                    ลบ
                                </button>
                            </div>
                        </div>
                    </button>
                ))}
                <div
                    style={{
                        marginLeft: "170px",
                        width: "300px",
                        height: "50px",
                        backgroundColor: "#001529",
                        color: "#f7f3f3",
                        padding: "15px",
                        borderRadius: "10px",
                        textAlign: "center",
                    }}
                    onClick={() => navigate("/AddProduct")}
                >
                    + เพิ่มรายการสินค้า
                </div>
            </div>
            <Modal
                title="แก้ไขสินค้า"
                visible={isEditModalVisible}
                onCancel={closeEditModal}
                footer={null}
            >
                <Form form={form} layout="vertical" onFinish={handleEditSubmit}>
                    <Form.Item
                        label="Name"
                        name="Name"
                        rules={[{ required: true, message: "กรุณาใส่ชื่อสินค้า" }]}
                    >
                        <Input placeholder="ชื่อสินค้า" />
                    </Form.Item>
                    <Form.Item
                        label="Description"
                        name="Description"
                        rules={[{ required: true, message: "กรุณาใส่คำอธิบาย" }]}
                    >
                        <Input placeholder="คำอธิบาย" />
                    </Form.Item>
                    <Form.Item
                        label="Category"
                        name="CatagoryID"
                        rules={[{ required: true, message: "กรุณาเลือกหมวดหมู่" }]}
                    >
                        <Select
                            placeholder="เลือกหมวดหมู่"
                            onChange={(value) => setSelectedCatagoryID(value)}
                        >
                            {catagory.map((cat) => (
                                <Select.Option key={cat.ID} value={cat.ID}>
                                    {cat.Name}
                                </Select.Option>
                            ))}
                        </Select>
                    </Form.Item>
                    <Form.Item label="Image">
                        <Upload
                            fileList={fileList}
                            onChange={({ fileList: newFileList }) => setFileList(newFileList)}
                            beforeUpload={() => false}
                            listType="picture-card"
                        >
                            {fileList.length < 1 && "+ Upload"}
                        </Upload>
                    </Form.Item>
                    <div style={{ textAlign: "right" }}>
                        <Button onClick={closeEditModal} style={{ marginRight: "10px" }}>ยกเลิก</Button>
                        <Button type="primary" htmlType="submit">บันทึก</Button>
                    </div>
                </Form>
            </Modal>
        </div>
    );
};

export default EditProduct;
