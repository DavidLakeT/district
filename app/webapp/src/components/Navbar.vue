<template>
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
      <a class="navbar-brand" href="#">District VulnWeb</a>
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse justify-content-end" id="navbarNav">
        <ul class="navbar-nav">
          <li class="nav-item">
            <router-link to="/" class="nav-link">Home</router-link>
          </li>
          <li class="nav-item">
            <router-link to="/products" class="nav-link">Products</router-link>
          </li>
          <li class="nav-item">
            <router-link to="/mycart" class="nav-link">Cart</router-link>
          </li>
          <li class="nav-item">
            <router-link to="/login" class="nav-link">Sign in</router-link>
          </li>
          <!-- Add more links as needed -->
        </ul>
      </div>
    </nav>
  </template>
  
  <script>
  export default {
    name: "Navbar"
  }
  </script>
  
  <style scoped>
  /* Add component-specific styles here */
  /* If needed */
  </style>
  