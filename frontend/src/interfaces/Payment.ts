import { UserInterface } from './User'

export interface Payment {

  Date: string; // วันที่ (ISO 8601 format, required)

  UserID: number; // User ID
  User?: UserInterface; // Optional, ข้อมูลผู้ใช้ที่เกี่ยวข้อง

  PaymentMethodID: number;
  PaymentMethod: PaymentMethod;

  PaymentStatusID: number;
  PaymentStatus: PaymentStatus;

}

export interface PaymentMethod {

  PaymentMethod: string;

}

export interface PaymentStatus {

  PaymentStatus: string;

}


