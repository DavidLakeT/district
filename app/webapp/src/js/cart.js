import axios from 'axios';

const BASE_URL = 'http://localhost:5000/api/cart';

export const getCartInformation = async () => {
  try {
    const response = await axios.get(BASE_URL, {
      withCredentials: true,
    });
    return response.data;
  } catch (error) {
    console.error('Error fetching cart information:', error);
    return [];
  }
};

export const addCartItem = async (cartItemData) => {
  try {
    const response = await axios.post(BASE_URL, cartItemData, {
      withCredentials: true,
    });
    return response.data;
  } catch (error) {
    console.error('Error adding item to cart:', error);
    return [];
  }
};

export const placeOrder = async () => {
  try {
    const response = await axios.post(`${BASE_URL}/order`, null, {
      withCredentials: true,
    });
    return response.data;
  } catch (error) {
    console.error('Error placing order:', error);
    throw error;
    //return [];
  }
};

export const updateProductQuantity = async (cartItemData) => {
  try {
    const response = await axios.put(BASE_URL, cartItemData, {
      withCredentials: true,
    });
    return response.data;
  } catch (error) {
    console.error('Error updating product quantity in cart:', error);
    return [];
  }
};

export const removeItemFromCart = async (productID) => {
  try {
    const response = await axios.delete(BASE_URL, {
      withCredentials: true,
      data: { product_id: productID },
    });
    return response.data;
  } catch (error) {
    console.error('Error removing item from cart:', error);
    return [];
  }
};
