<template>
    <div class="container mt-5">
      <h2>Your Profile</h2>
      <form @submit.prevent="updateProfile" class="mt-4">
        <div class="form-group">
          <label for="name">Name:</label>
          <input type="text" v-model="user.name" class="form-control" id="name" required>
        </div>
        <div class="form-group">
          <label for="email">Email:</label>
          <input type="email" v-model="user.email" class="form-control" id="email" required>
        </div>
        <button @click="updateUserProfile" type="submit" class="btn btn-primary">Update Profile</button>
      </form>
    </div>
  </template>
  
<script>
import axios from 'axios';

export default {
  data() {
    return {
      user: {}, // store user details here
      updatedUserInfo: {} // store updated user info here
    };
  },
  created() {
    this.fetchUserProfile();
  },
  methods: {
    async fetchUserProfile(userId) {
      try {
        const response = await axios.get(`/api/user/${userId}`);
        this.user = response.data;
      } catch (error) {
        console.error('Error fetching user profile:', error);
      }
    },

    async updateUserProfile(userId) {
      try {
        const response = await axios.put(`/api/user/${userId}`, this.updatedUserInfo);
        console.log('User profile updated:', response.data);
      } catch (error) {
        console.error('Error updating user profile:', error);
      }
    }
  }
};
</script>

  
  <style scoped>
  .container {
    padding: 20px;
  }
  
  .form-group {
    margin-bottom: 1rem;
  }
  
  .btn-primary {
    background-color: #007BFF;
    border: none;
  }
  
  .btn-primary:hover {
    background-color: #0056b3;
  }
  </style>
  