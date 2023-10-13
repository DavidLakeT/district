<template>
    <div class="container mt-5">
      <h2 class="text-center">Register</h2>
      <form @submit.prevent="register" class="mt-4">
        <div class="form-group">
          <label for="name">Full Name:</label>
          <input type="text" v-model="username" class="form-control" id="name" placeholder="Enter your full name" required>
        </div>
        <div class="form-group">
          <label for="email">Email:</label>
          <input type="email" v-model="email" class="form-control" id="email" placeholder="Enter your email" required>
        </div>
        <div class="form-group">
          <label for="password">Password:</label>
          <input type="password" v-model="password" class="form-control" id="password" placeholder="Enter your password" required>
        </div>
        <button  @click="registerUser" type="submit" class="btn btn-primary btn-block">Register</button>
      </form>
      <div class="mt-3 text-center">
        Already have an account? <router-link to="/login">Login</router-link>
      </div>
    </div>
  </template>
  
<script>
import axios from 'axios';

export default {

  data() {
    return {
      username: '',
      email: '',
      password: ''
    };
  },

  methods: {
    async registerUser() {
      try {
        const response = await axios.post('/api/user', {
          username: this.username,
          email: this.email,
          password: this.password
        });
        console.log('User registered:', response.data);
      } catch (error) {
        console.error('Error registering user:', error);
      }
    }
  }
};
</script>



  
  <style scoped>
  .container {
    margin-bottom: 17%;
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
  
  .btn-block {
    width: 100%;
  }
  
  .text-center {
    text-align: center;
  }
  </style>
  