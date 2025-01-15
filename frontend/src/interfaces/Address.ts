import { UserInterface } from "./User";

export interface AddressInterface {

    ID?: number;

    full_address?: 	string;    	

   	city?:  	    string;    	

  	province?:     	string;    

   	postal_code?:     string;		

	user_id?: number;   // Add user_id for linking with User
	
	user?: UserInterface;  
	
}