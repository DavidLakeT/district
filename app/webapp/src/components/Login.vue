<template>
    <div class="container mt-5">
      <h2 class="text-center">Login</h2>
      <form @submit.prevent="login" class="mt-4">
        <div class="form-group">
          <label for="email">Email:</label>
          <input type="email" v-model="email" class="form-control" id="email" placeholder="Enter your email" required>
        </div>
        <div class="form-group">
          <label for="password">Password:</label>
          <input type="password" v-model="password" class="form-control" id="password" placeholder="Enter your password" required>
        </div>
        <button type="submit" class="btn btn-primary btn-block">Login</button>
      </form>
      <div class="mt-3 text-center">
        Don't have an account? <router-link to="/register">Create one</router-link>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: "Login",
    data() {
      return {
        email: "",
        password: ""
      };
    },
    methods: {
      login() {
        // Implement login functionality here
        console.log("Logged in with email:", this.email);
      }
    }
  };
  </script>
  
  <style scoped>
  /* Add component-specific styles here */
  
  /* Optional: Add custom styles */
  /* Adjust margins and padding as needed */
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
  