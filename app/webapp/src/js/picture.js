import axios from 'axios';

const BASE_URL = 'http://localhost:5000/api/product';

export const getProductImage = async (productId) => {
    try {
      const response = await axios.get(`${BASE_URL}/picture?filename=${productId}.jpg`, {
        withCredentials: true,
        responseType: 'blob', 
      });
      return response.data;
    } catch (error) {
      return null;
    }
  };
  

  export const uploadProductImage = async (imageFile) => {
    try {
      const formData = new FormData();
      formData.append('file', imageFile);
  
      const response = await axios.post(`${BASE_URL}/upload`, formData, {
        withCredentials: true,
      });
      return response.data;
    } catch (error) {
      console.error('Error uploading product image:', error);
      return null;
    }
  };
  
  