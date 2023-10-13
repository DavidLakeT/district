<template>
    <div class="container mt-5">
      <h2 class="text-center">Register</h2>
      <form @submit.prevent="register" class="mt-4">
        <div class="form-group">
          <label for="name">Full Name:</label>
          <input type="text" v-model="name" class="form-control" id="name" placeholder="Enter your full name" required>
        </div>
        <div class="form-group">
          <label for="email">Email:</label>
          <input type="email" v-model="email" class="form-control" id="email" placeholder="Enter your email" required>
        </div>
        <div class="form-group">
          <label for="password">Password:</label>
          <input type="password" v-model="password" class="form-control" id="password" placeholder="Enter your password" required>
        </div>
        <button type="submit" class="btn btn-primary btn-block">Register</button>
      </form>
      <div class="mt-3 text-center">
        Already have an account? <router-link to="/login">Login</router-link>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: "Register",
    data() {
      return {
        name: "",
        email: "",
        password: ""
      };
    },
    methods: {
      register() {
        // Implement registration functionality here
        console.log("Registered with name:", this.name, "email:", this.email);
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
  