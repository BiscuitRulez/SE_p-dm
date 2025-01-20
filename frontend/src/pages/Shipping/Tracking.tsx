import React from 'react';
import './Tracking.css';
import back from "../../assets/back.png";
import axios from 'axios';

const Tracking: React.FC = () => {
    return (
        <div>
          
            <a href="/" style={{ position: "absolute", top: "100px", right: "1400px" }}>
                <img style={{ width: "50px", height: "auto" }} src={back} alt="Back" />
            </a>
            <div className="body">
                <div className="tracking-page">
                    <h1 className="title">ติดตามสถานะการจัดส่ง</h1>

                    <div className="order-section">
                        <div className="order-item">
                            <img
                                src="https://via.placeholder.com/150"
                                alt="เก้าอี้"
                                className="item-image"
                            />
                            <div className="item-details">
                                <p>เก้าอี้เชือก</p>
                                <p>สีน้ำตาล ขนาดกลาง จำนวน 1 ตัว</p>
                                <p>ยอดรวมทั้งหมด 500 บาท</p>
                            </div>
                        </div>

                        <div className="delivery-info">
                            <h2>ที่อยู่ในการจัดส่ง</h2>
                            <p>นางสาว ใจกล้า โทร: 000-000-0000</p>
                            <p>บ้านเลขที่ 123 ถนนสุขุมวิท ซอย 11 แขวงคลองตันเหนือ เขตวัฒนา</p>
                            <p>กรุงเทพมหานคร 10110</p>
                            <p>จะได้รับในวันที่ 00/00/0000</p>
                            <p>ให้บริการจัดส่งโดย iGotSofa.com</p>
                        </div>
                    </div>

                    <div className="status-section">
                        <h2>สถานะการจัดส่ง</h2>
                        <div className="status-item">
                            <span className="status-date">15/10/24</span>
                            <span className="status-description">ขนส่งเข้ารับพัสดุแล้ว</span>
                        </div>
                        <div className="status-item">
                            <span className="status-date">16/10/24</span>
                            <span className="status-description">พัสดุอยู่ระหว่างการจัดส่ง</span>
                        </div>
                        <div className="status-item">
                            <span className="status-date">19/10/24</span>
                            <span className="status-description">พัสดุถูกจัดส่งเรียบร้อยแล้ว</span>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default Tracking;
