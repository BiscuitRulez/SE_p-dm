import { UserInterface } from './User';


export interface ClaimInterface {
  Date: string; // วันที่ (ISO 8601 format, required)
  Photo: string; // URL หรือ Path ของรูปภาพที่เกี่ยวข้อง
  Reason: string;

  ProblemID: number; // ID ของปัญหา
  Problem?: Problem; // Optional, ข้อมูลปัญหาที่เกี่ยวข้อง

  ClaimStatusID: number; // ID ของสถานะการเคลม
  ClaimStatus?: ClaimStatus; // Optional, ข้อมูลสถานะการเคลม

  UserID: number; // ID ของผู้ใช้
  User?: UserInterface; // Optional, ข้อมูลผู้ใช้ที่เกี่ยวข้อง

  OrderID: number; // ID ของคำสั่งซื้อ
  Order?: Order; // Optional, ข้อมูลคำสั่งซื้อที่เกี่ยวข้อง
}

export interface ClaimStatus {

    ClaimStatus: string;
  
}

export interface Problem {

    Problem: string;
  
}

export interface Order {

    Order: string;
  
}