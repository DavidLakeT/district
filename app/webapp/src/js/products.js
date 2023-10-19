import axios from 'axios';

const BASE_URL = 'your_api_base_url'; // Replace with your API's base URL

export const fetchProducts = async () => {
  try {
    const response = await axios.get(`${BASE_URL}/products`);
    return response.data;
  } catch (error) {
    console.error('Error fetching products:', error);
    return [];
  }
};
