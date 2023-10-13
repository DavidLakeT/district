<template>
    <div class="container mt-5">
      <div class="row">
        <div class="col-md-8 offset-md-2">
          <!-- Display product information here -->
          <div v-for="product in products" :key="product.id" class="card">
            <img :src="product.image" class="card-img-top" :alt="product.name">
            <div class="card-body">
              <h5 class="card-title">{{ product.name }}</h5>
              <p class="card-text">$ {{ product.price.toFixed(2) }}</p>
              <a href="#" class="btn btn-primary">Add to Cart</a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name:"Products",
    data() {
      return {
        products: [
          { id: 1, name: "Product 1", price: 10.99, image: "images/product1.jpg" },
          { id: 2, name: "Product 2", price: 15.99, image: "images/product2.jpg" }
        ]
      };
    }
  }
  </script>
  
  <style scoped>
  /* Add component-specific styles here */
  </style>
  