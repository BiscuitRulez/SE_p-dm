//AddProduct.tsx 
import React, { useState, useEffect } from "react";
import "./AddProduct.css";
import back from "../../../assets/back.png";
import "../../../components/card4.css";
import "../../../components/t.css";
// import { GetProductByID } from "../../services/https";
// import axios from "axios";

import { Typography, Modal, Form, Upload, message } from "antd";
import { CatagoryInterface } from "../../../interfaces/Catagory";
import { TagsInterface } from "../../../interfaces/Tags";
import { Product } from "../../../interfaces/Product";
import { CreateProduct, GetCatagory, GetTags, UploadProductImages } from "../../../services/https";

import { useNavigate, useLocation } from "react-router-dom";


// interface EditProductProps {
//   closePopup: () => void;
//   productId: string | undefined;
// }

const AddProduct: React.FC = () => {
  const [messageApi] = message.useMessage();
  const [form] = Form.useForm();

  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const navigate = useNavigate();
  const [catagory, setCatagory] = useState<CatagoryInterface[]>([]);
  const [selectedCategory, setSelectedCategory] = useState("");

  // const [tags, setTags] = useState<TagsInterface[]>([]);
  // const [selectedTags, setSelectedTags] = useState("");

  const [image, setImage] = useState<string | null>(null);
  const userIdstr = localStorage.getItem("id");
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState("");

  const location = useLocation();
  const { product } = location.state || {}; // รับข้อมูลจาก state
  // if (!productId || !product) {
  //   return <div>ไม่พบข้อมูลสินค้า</div>;
  // }

  const fetchCatagoryData = async () => {
    try {
      const res = await GetCatagory();
      if (res?.data) {
        setCatagory(res.data.catagorys);
      } else {
        setError("Failed to load Categories");
      }
    } catch (err) {
      setError("Error fetching Categories. Please try again later.");
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    console.log(catagory); // ตรวจสอบค่า categories
  }, [catagory]);

  // const GetProductByID = async (productId: any) => {
  //   if (productId) {
  //     try {
  //       const res = await GetProductByID(productId);
  //       console.log(res.data);
  //       form.setFieldsValue({
  //         ID: res.data.ID,
  //         Name: res.data.Name,
  //         Description: res.data.Description,
  //         Image: res.data.Image,
  //         CatagoryID: res.data.CatagoryID,
  //         UserID: res.data.UserID,
  //       });
  //       setCatagory(res.data.CatagoryID); // Ensure the category state is set
  //     } catch (error) {
  //       messageApi.error("Error fetching Album Information");
  //     }
  //   } else {
  //     messageApi.error("Album ID is required.");
  //   }
  // };

  // const fetchTagsData = async () => {
  //   try {
  //     const res = await GetTags();
  //     if (res?.data) {
  //       console.log(res.data)
  //       setTags(res.data.tags);
  //     } else {
  //       setError("Failed to load Tags");
  //     }
  //   } catch (err) {
  //     setError("Error fetching Tags. Please try again later.");
  //   } finally {
  //     setIsLoading(false);
  //   }
  // };

  // useEffect(() => {
  //   console.log(tags); // ตรวจสอบค่า categories
  // }, [tags]);

  useEffect(() => {
    if (userIdstr) {
      // fetchTagsData();
      fetchCatagoryData();
    } else {
      message.error("The user ID was not found in localStorage.");
    }
  }, [userIdstr]);

  const [formData, setFormData] = useState({
    name: "",
    selectedCategory: "",
    description: "",
    image: null,
  });

  useEffect(() => {
    console.log("Received product in AddProduct: ", product);
    if (product) {
      setName(product.Name);
      setDescription(product.Description);
      setSelectedCategory(product.CatagoryID);
      setImage(product.Image || null); // กำหนดค่าภาพ
    }
  }, [product]);

  const handleImageUpload = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files[0]) {
      const file = e.target.files[0];
      const reader = new FileReader();

      reader.onload = () => {
        setImage(reader.result as string);
      };
      reader.readAsDataURL(file);
    }
  };

  const handleSave = async () => {
    if (!name || !description || !selectedCategory || !image) {
      alert("กรุณากรอกข้อมูลให้ครบถ้วน!");
      return;
    }
  
    const productData = {
      name,
      description,
      selectedCategory,
      image,
    };
  
    // ตรวจสอบว่า `product.id` มีค่าหรือไม่
    const url = product?.ID
      ? `http://localhost:8000/product/${product.ID}`  // ใช้ PUT หากมี ID (อัปเดต)
      : "http://localhost:8000/product";  // ใช้ POST หากไม่มี ID (เพิ่มใหม่)
  
    const method = product?.ID ? "PUT" : "POST";  // ใช้ PUT หากมี ID (อัปเดต) หรือ POST (เพิ่มใหม่)
  
    try {
      const response = await fetch(url, {
        method,
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(productData),
      });
  
      if (response.ok) {
        message.success(product?.ID ? "อัปเดตข้อมูลเรียบร้อย!" : "เพิ่มข้อมูลเรียบร้อย!");
        setTimeout(() => {
          // นำทางกลับไปหน้า EditProduct และส่งข้อมูล product.id (หากต้องการ)
          navigate(`/`);
        }, 500);
      } else {
        message.error("เกิดข้อผิดพลาดในการบันทึกข้อมูล!");
      }
    } catch (error) {
      console.error("Error saving stock:", error);
      message.error("ไม่สามารถเชื่อมต่อกับเซิร์ฟเวอร์ได้!");
    }
  };

  return (
    <div>
      <a href="/" style={{ position: "absolute", top: "100px", right: "1400px" }}>
        <img style={{ width: "50px", height: "auto" }} src={back} alt="Back" />
      </a>
      <div className="body">
        <div className="head">
          <h1>เพิ่มรายการสินค้า</h1>
        </div>
        <div className="add-product-container">
          <form className="add-product-form" >
            {/* <form className="add-product-form" onSubmit={handleSubmit}> */}
            <label>ชื่อ</label>
            <input
              className="name"
              type="text"
              placeholder="ชื่อ"
              value={name}
              onChange={(e) => setName(e.target.value)}
            />
            <label>คำอธิบาย</label>
            <textarea
              className="describe"
              placeholder="คำอธิบาย"
              value={description}
              onChange={(e) => setDescription(e.target.value)}
            ></textarea>
            <div className="hr-p-c">
              {/* <label>ราคา</label> */}
              <label>หมวดหมู่</label>
            </div>
            <div className="price-category">
              <select
                className="cate"
                value={selectedCategory}
                onChange={(e) => setSelectedCategory(e.target.value)}
              >
                <option value="">เลือกหมวดหมู่</option>
                {isLoading ? (
                  <option>Loading...</option>
                ) : catagory.length > 0 ? (
                  catagory.map((item) => (
                    <option key={item.ID} value={item.ID}>  {/* ใช้ ID แทน Name */}
                      {item.Name}
                    </option>
                  ))
                ) : (
                  <option>No categories available</option>
                )}
              </select>
            </div>
            {/* <label>Tags</label>
            <select
              className="tag"
              value={selectedTags}
              onChange={(e) => setSelectedTags(e.target.value)}
            >
              <option value="">เลือก tags ที่ใช้ในการค้นหา</option>
              {isLoading ? (
                <option>Loading...</option>
              ) : tags.length > 0 ? (
                tags.map((items) => (
                  <option key={items.ID} value={items.Tag_Name}>
                    {items.Tag_Name}
                  </option>
                ))
              ) : (
                <option>No Tags available</option>
              )}
            </select> */}

            {/* <input
              className="tag"
              type="text"
              placeholder="เลือก tags ที่ใช้ในการค้นหา"
              value={tags}
              onChange={(e) => setTags(e.target.value)}
            /> */}
          </form>

          <div className="image-upload">
            <label htmlFor="image-upload" className="upload-box" style={{ cursor: "pointer" }}>
              {image ? (  // ตรวจสอบค่า image
                <img src={image} alt="Uploaded" style={{ width: "100%", height: "auto" }} />
              ) : (
                <span className="upload-icon">+</span>
              )}
            </label>
            <input
              type="file"
              id="image-upload"
              style={{ display: "none" }}
              accept="image/*"
              onChange={handleImageUpload}
            />
          </div>

        </div>
        <div className="button">
          <div className="save">
            <button onClick={handleSave}>บันทึก</button>
          </div>
          <div className="cancel">
            <button onClick={() => alert("ยกเลิก")}>ยกเลิก</button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default AddProduct;


