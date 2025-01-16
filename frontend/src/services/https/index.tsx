import { AddressInterface } from "../../interfaces/Address";
import { CodeInterface } from "../../interfaces/Code";
import { SignInInterface } from "../../interfaces/SignIn";
import { UserInterface } from "../../interfaces/User";
import { Payment } from "../../interfaces/Payment";
import { Order } from "../../interfaces/Order";
import { CartItem } from "../../interfaces/Cart";
import { Product } from "../../interfaces/Product";
import { Review, ReviewAnalytics, ReviewsPagination, ReviewInput } from "../../interfaces/Review";
import { Stock } from "../../interfaces/Stock";

import axios from "axios";

const apiUrl = "http://localhost:8000";

const Authorization = localStorage.getItem("token");

const Bearer = localStorage.getItem("token_type");

const requestOptions = {

  headers: {

    "Content-Type": "application/json",

    Authorization: `${Bearer} ${Authorization}`,

  },

};

async function SignIn(data: SignInInterface) {

  return await axios

    .post(`${apiUrl}/signin`, data, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}

async function CreateCode(data: CodeInterface) {

  return await axios

    .post(`${apiUrl}/codes`, data, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}

async function GetCodes() {

  return await axios

    .get(`${apiUrl}/codes`, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}

async function GetCodesById(id: Number | undefined) {

  return await axios

    .get(`${apiUrl}/codes/${id}`, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}


async function UpdateCode(data: CodeInterface) {

  return await axios

    .put(`${apiUrl}/codes/${data.ID}`, data, requestOptions)  // เพิ่ม ${data.ID} ใน URL

    .then((res) => res)

    .catch((e) => e.response);
}



async function DeleteCodeById(id: Number | undefined) {

  return await axios

    .delete(`${apiUrl}/codes/${id}`, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}

async function CreateUser(data: UserInterface) {

  return await axios

    .post(`${apiUrl}/signup`, data, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}

async function CreateAdmin(data: UserInterface) {

  return await axios

    .post(`${apiUrl}/signupadmin`, data, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}

async function GetUsers() {

  return await axios

    .get(`${apiUrl}/users`, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}

async function GetAdmin() {

  return await axios

    .get(`${apiUrl}/admins`, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}

async function GetPaymentMethod() {

  return await axios

    .get(`${apiUrl}/paymentMethod`, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}


async function GetUsersById(id: string) {

  return await axios

    .get(`${apiUrl}/user/${id}`, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}


async function UpdateUsersById(id: string, data: UserInterface) {

  return await axios

    .put(`${apiUrl}/user/${id}`, data, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}


async function DeleteUsersById(id: string) {

  return await axios

    .delete(`${apiUrl}/user/${id}`, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}

async function UpdateCodeAfterCollect(codeId: string) {

  return await axios

    .put(`${apiUrl}/code-collect/${codeId}`, {}, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);
}

async function AddCodeToCollect(userId: string, codeId: string) {

  return await axios

    .post(`${apiUrl}/code-collect/${userId}/${codeId}`, {}, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);
}

async function GetCollectedCodes(userId: string) {  ///use this

  return await axios

    .get(`${apiUrl}/code-collect/${userId}`, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);
}

async function GetCollectedCodesToShow(Id: string) {

  return await axios

    .get(`${apiUrl}/show-collect/${Id}`, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}

async function CreateAddress(data: AddressInterface) {

  return await axios

    .post(`${apiUrl}/address`, data, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}

async function GetAddressesByUserId(id: string) {

  return await axios

    .get(`${apiUrl}/address/${id}`, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}

//------------------------------------------------------------------

async function GetPayments() {
  return await axios
    .get(`${apiUrl}/payments`, requestOptions)
    .then((res) => res)
    .catch((e) => e.response);
}

async function GetPaymentByID(id: string) {
  return await axios
    .get(`${apiUrl}/payments/${id}`, requestOptions)
    .then((res) => res)
    .catch((e) => e.response);
}

async function CreatePayment(id: string, data: Payment) {
  return await axios
    .post(`${apiUrl}/payments/${id}`, data, requestOptions)
    .then((res) => res)
    .catch((e) => e.response);
}

async function UpdatePaymentByUserID(id: string, data: Payment) {
  return await axios
    .put(`${apiUrl}/payments/${id}`, data, requestOptions)
    .then((res) => res)
    .catch((e) => e.response);
}

async function DeletePayment(id: string) {
  return await axios
    .delete(`${apiUrl}/payments/${id}`, requestOptions)
    .then((res) => res)
    .catch((e) => e.response);
}

async function GetAllShipping() {
  return await axios
    .get(`${apiUrl}/Shipping`, requestOptions)
    .then((res) => res)
    .catch((e) => e.response);
}

async function getCart(userId: string): Promise<CartItem[]> {
  try {
    const response = await fetch(`${apiUrl}/cart?user_id=${userId}`);
    if (!response.ok) {
      throw new Error('Failed to fetch cart');
    }
    return response.json();
  } catch (error) {
    console.error('Cart fetch error:', error);
    throw error;
  }
}

async function addToCart(userId: string, productId: number): Promise<CartItem> {
  try {
    const response = await fetch(`${apiUrl}/cart`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        user_id: userId,
        product_id: productId,
        quantity: 1,
      }),
    });

    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.error || 'Failed to add to cart');
    }

    return response.json();
  } catch (error) {
    console.error('Add to cart error:', error);
    throw error;
  }
}

async function updateCartItem(itemId: number, quantity: number): Promise<CartItem> {
  const response = await fetch(`${apiUrl}/cart/${itemId}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ quantity }),
  });
  if (!response.ok) throw new Error('Failed to update cart item');
  return response.json();
}

async function createOrder(userId: string): Promise<Order> {
  try {
    console.log('Making create order request...');
    const response = await fetch(`${apiUrl}/orders?user_id=${userId}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
    });

    if (!response.ok) {
      const errorData = await response.json();
      console.error('Order creation failed:', errorData);
      throw new Error(errorData.error || 'Failed to create order');
    }

    const data = await response.json();
    console.log('Order creation successful:', data);
    return data;
  } catch (error) {
    console.error('Order creation error:', error);
    throw error;
  }
}

async function getOrders(userId: string): Promise<Order[]> {
  try {
    const response = await fetch(`${apiUrl}/orders?user_id=${userId}`);
    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.error || 'Failed to fetch orders');
    }
    return response.json();
  } catch (error) {
    console.error('Order fetch error:', error);
    throw error;
  }
}

async function getProducts(): Promise<Product[]> {
  const response = await fetch(`${apiUrl}/products`);
  if (!response.ok) throw new Error('Failed to fetch products');
  return response.json();
}

async function getProductStock(productId: number): Promise<Stock> {
  const response = await fetch(`${apiUrl}/products/${productId}/stock`);
  if (!response.ok) throw new Error('Failed to fetch stock');
  return response.json();
}

async function getProductDetails(id: number): Promise<Product> {
  try {
    const response = await fetch(`${apiUrl}/products/${id}`);
    if (!response.ok) {
      if (response.status === 404) {
        throw new Error('Product not found');
      }
      throw new Error('Failed to fetch product details');
    }
    const data = await response.json();
    console.log('Product details response:', data);
    return data;
  } catch (error) {
    console.error('Error fetching product:', error);
    throw error;
  }
}

// async function createReview(review: ReviewInput,getProductDetails: (productId: number) => Promise<void>): Promise<Review> {
//   const response = await fetch(`${apiUrl}/reviews`, {
//     method: 'POST',
//     headers: { 'Content-Type': 'application/json' },
//     body: JSON.stringify({
//       product_id: review.ProductID,
//       user_id: review.UserID,
//       rating: review.Rating,
//       comment: review.Comment,
//     }),
//   });

//   if (!response.ok) throw new Error('Failed to submit review');

//   await getProductDetails(review.ProductID); // Call the function directly
//   return response.json();
// }


async function getProductReviews(productId: number): Promise<Review[]> {
  const response = await fetch(`${apiUrl}/products/${productId}/reviews`);
  if (!response.ok) throw new Error('Failed to fetch reviews');
  return response.json();
}

async function getReviews(productId: number, pageParam: any): Promise<ReviewsPagination> {
  try {
    const response = await fetch(`${apiUrl}/products/${productId}/reviews`);
    if (!response.ok) {
      if (response.status === 404) {
        return {
          items: [],
          total: 0,
          page: 1,
          totalPages: 1,
          hasNextPage: false
        };
      }
      throw new Error('Failed to fetch reviews');
    }
    const data = await response.json();
    return {
      items: data.items || [],
      total: data.total || 0,
      page: data.page || 1,
      totalPages: data.totalPages || 1,
      hasNextPage: data.hasNextPage || false
    };
  } catch (error) {
    console.error('Error fetching reviews:', error);
    throw error;
  }
}

async function createReview(review: ReviewInput): Promise<Review> {
  try {
    // Log the request data
    console.log('Sending review data:', review);

    const response = await fetch(`${apiUrl}/reviews`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json'
      },
      body: JSON.stringify({
        product_id: review.productId,
        user_id: review.userId,
        rating: review.rating,
        comment: review.comment,
        images: review.images || []
      }),
    });

    // Log the response status
    console.log('Response status:', response.status);

    if (!response.ok) {
      const errorText = await response.text();
      console.error('Error response:', errorText);

      try {
        const errorData = JSON.parse(errorText);
        throw new Error(errorData.error || 'Failed to create review');
      } catch (e) {
        throw new Error(errorText || 'Failed to create review');
      }
    }

    const data = await response.json();
    console.log('Response data:', data);

    return {
      id: data.data.ID,
      productId: data.data.ProductID,
      userId: data.data.UserID,
      rating: data.data.Rating,
      comment: data.data.Comment,
      images: data.data.Images || [],
      helpfulVotes: data.data.HelpfulVotes || 0,
      verifiedPurchase: data.data.VerifiedPurchase || false,
      createdAt: data.data.CreatedAt,
      updatedAt: data.data.UpdatedAt
    };
  } catch (error) {
    console.error('Error creating review:', error);
    throw error;
  }
}

async function uploadImage(file: File): Promise<string> {
  const formData = new FormData();
  formData.append('file', file);

  const response = await fetch(`${apiUrl}/reviews/upload`, {
    method: 'POST',
    body: formData,
  });

  if (!response.ok) {
    const error = await response.json();
    throw new Error(error.error || 'Failed to upload image');
  }

  const data = await response.json();
  return data.imageUrl;
}

async function voteHelpful(reviewId: number): Promise<void> {
  try {
    const response = await fetch(`${apiUrl}/reviews/${reviewId}/vote`, {
      method: 'POST',
    });

    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.error || 'Failed to vote');
    }
  } catch (error) {
    console.error('Error voting:', error);
    throw error;
  }
}

async function getAnalytics(productId: number): Promise<ReviewAnalytics> {
  try {
    const response = await fetch(
      `${apiUrl}/products/${productId}/reviews/analytics`
    );

    if (!response.ok) {
      const error = await response.json();
      throw new Error(error.error || 'Failed to fetch analytics');
    }

    const data = await response.json();
    return {
      productId: productId,
      totalReviews: data.total_reviews || 0,
      averageRating: data.average_rating || 0,
      ratingDistribution: data.rating_distribution || {},
      helpfulVotes: data.helpful_votes || 0,
      responseRate: data.response_rate || 0,
      verifiedPurchaseRate: data.verified_purchase_rate || 0
    };
  } catch (error) {
    console.error('Error fetching analytics:', error);
    throw error;
  }
}

async function GetProducts() {

  return await axios

    .get(`${apiUrl}/products`, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);

}

async function GetCatagory() {
  try {
    const res = await axios.get(`${apiUrl}/catagory`, requestOptions);
    return res;
  } catch (error: any) {
    return error.response || { status: 500, message: "Unknown Error" };
  }
}

async function GetTags() {
  try {
    const res = await axios.get(`${apiUrl}/tags`, requestOptions);
    return res;
  } catch (error: any) {
    return error.response || { status: 500, message: "Unknown Error" };
  }
}

async function GetAddress() {
  try {
    const res = await axios.get(`${apiUrl}/tags`, requestOptions);
    return res;
  } catch (error: any) {
    return error.response || { status: 500, message: "Unknown Error" };
  }
}

async function UploadProductImages(data: Product) {
  return await axios
    .post(`${apiUrl}/product`, data, requestOptions)
    .then((res) => res)
    .catch((e) => e.response);
}

async function CreateProduct(data: Product) {
  return await axios
    .post(`${apiUrl}/product`, data, requestOptions)
    .then((res) => res)
    .catch((e) => e.response);
}

async function GetProductByID(id: string) {
  return await axios

    .get(`${apiUrl}/product/${id}`, requestOptions)

    .then((res) => res)

    .catch((e) => e.response);
}

async function DeleteProduct(id: Number) {
  return await axios
    .delete(`${apiUrl}/product/${id}`, requestOptions)
    .then((res) => res)
    .catch((e) => e.response);
}

async function UpdateProductByID(id: string, data: Product) {
  return await axios
    .put(`${apiUrl}/product/${id}`, data, requestOptions)
    .then((res) => res)
    .catch((e) => e.response);
}

async function GetProduct() {
  try {
    const res = await axios.get(`${apiUrl}/product`, requestOptions);
    return res;
  } catch (error: any) {
    return error.response || { status: 500, message: "Unknown Error" };
  }
}


export {

  CreateCode,
  GetCodes,
  GetCodesById,
  UpdateCode,
  DeleteCodeById,
  SignIn,
  CreateUser,
  CreateAdmin,
  GetUsers,
  GetAdmin,
  GetUsersById,
  UpdateUsersById,
  DeleteUsersById,
  UpdateCodeAfterCollect,
  AddCodeToCollect,
  GetCollectedCodes,
  CreateAddress,
  GetAddressesByUserId,
  GetPayments,
  GetPaymentByID,
  CreatePayment,
  UpdatePaymentByUserID,
  DeletePayment,
  GetPaymentMethod,
  GetCollectedCodesToShow,
  GetAllShipping,
  getCart,
  addToCart,
  updateCartItem,
  createOrder,
  getOrders,
  getProducts,
  getProductStock,
  getProductDetails,
  createReview,
  getProductReviews,
  uploadImage,
  getReviews,
  voteHelpful,
  getAnalytics,
  GetProducts,
  
  GetCatagory,
  GetTags,
  GetAddress,
  UploadProductImages,
  CreateProduct,
  GetProductByID,
  DeleteProduct,
  UpdateProductByID,
  GetProduct


};