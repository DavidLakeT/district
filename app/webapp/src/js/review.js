import axios from 'axios';

const BASE_URL = 'http://localhost:5000';

export const submitReviewForm = async (reviewData, authToken) => {
  try {
    const response = await axios.post(`${BASE_URL}/api/review`, reviewData, {
      headers: {
        Authorization: `${authToken}`,
      },
    });
    return response.data;
  } catch (error) {
    console.error('Error creating review:', error);
    return [];
  }
};
