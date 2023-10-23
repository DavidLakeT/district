import axios from 'axios';

const BASE_URL = 'http://localhost:5000';

export const submitReviewForm = async (reviewData) => {
  try {
    const response = await axios.post(`${BASE_URL}/api/review`, reviewData, {
      withCredentials: true
    });
    return response.data;
  } catch (error) {
    console.error('Error creating review:', error);
    return [];
  }
};
