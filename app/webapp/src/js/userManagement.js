import axios from 'axios';

const BASE_URL = 'http://localhost:5000';

export const getUsers = async () => {
  try {
    const response = await axios.get(`${BASE_URL}/api/user`,{
      withCredentials: true
    });
    return response.data;
  } catch (error) {
    console.error('Error fetching products:', error);
    return [];
  }
};

export const updateUser = async (userData, userId) => {
  try {
    const response = await axios.put(`${BASE_URL}/api/user/${userId}`,
      userData, {
      withCredentials:true
    });
    return response.data;
  } catch (error) {
    console.error('Error updating user:', error);
    return [];
  }
};

export const deleteUser = async (userId) => {
  try {
    const response = await axios.delete(`${BASE_URL}/api/user/${userId}`,{
      withCredentials:true
    });
    return response.data;
  } catch (error) {
    console.error('Error deleting user:', error);
    return [];
  }
};
