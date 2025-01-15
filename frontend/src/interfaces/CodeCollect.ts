import { CodeInterface } from "./Code";
import { UserInterface } from "./User";

export interface CodeCollectorInterface {

    ID?: number;
    
    user_id: number;

    user?: UserInterface; // ใช้ `Users` Interface สำหรับข้อมูลผู้ใช้

    code_id: number;

    code?: CodeInterface; // ใช้ `Codes` Interface สำหรับข้อมูลโค้ด

  }
  