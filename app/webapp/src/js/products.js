import axios from 'axios';

const BASE_URL = 'http://localhost:5000';

export const fetchProducts = async () => {
  try {
    const response = await axios.get(`${BASE_URL}/api/product`);
    return response.data;
  } catch (error) {
    console.error('Error fetching products:', error);
    return [];
  }
};

export const createProduct = async (productData) => {
  try {
    const response = await axios.post(`${BASE_URL}/api/product`, productData);
    return response.data;
  } catch (error) {
    console.error('Error creating product:', error);
    return [];
  }
};

export const fetchProductDetails = async (productId) => {
  try {
    const response = await axios.get(`${BASE_URL}/api/product/id/${productId}`);
    return response.data;
  } catch (error) {
    console.error('Error fetching product details:', error);
    throw error;
  }
};

export const searchProduct = async (productName) => {
  try {
    const response = await axios.get(`${BASE_URL}/api/product/name/${productName}`);
    return response.data;
  } catch (error) {
    console.error('Error fetching product details:', error);
    throw error;
  }
};

export const updateProduct = async (productData, productId) => {
  try {
    delete productData.reviews;
    const response = await axios.put(`${BASE_URL}/api/product/id/${productId}`,
      productData, {
      withCredentials:true
    });
    return response.data;
  } catch (error) {
    console.error('Error creating product:', error);
    return [];
  }
};

export const deleteProduct = async (productId) => {
  try {
    const response = await axios.delete(`${BASE_URL}/api/product/id/${productId}`,{
      withCredentials:true
    });
    return response.data;
  } catch (error) {
    console.error('Error creating product:', error);
    return [];
  }
};
