import { Stock } from "./Stock";
import { Review } from "./Review";

export interface Product {
    ID: number;
    name: string;
    description: string;
    image: string;
    stock: Stock;   // lowercase for consistency
    Stock?: Stock;  // uppercase for legacy support
    reviews: Review[];
    Reviews?: Review[];  // uppercase for legacy support
    avg_rating: number;
    AvgRating?: number;  // uppercase for legacy support
    CatagoryID?: number; 
	UserID?: number;
   
}