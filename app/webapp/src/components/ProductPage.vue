<template>
    <div class="container mt-5">
      <div class="row">
        <div class="col-md-6">
          <img :src="imageUrl" alt="Product Image" class="img-fluid" />
        </div>
        <div class="col-md-6">
          <div class="product-details">
            <h2 class="product-title">{{ product.name }}</h2>
            <p class="product-price">$ {{ product.price.toFixed(2) }}</p>
            <p class="product-description">{{ product.description }}</p>
            <p class="product-stock" v-if="product.stock > 0">In Stock: {{ product.stock }}</p>
            <p class="product-stock" v-else>Out of Stock</p>
            <div class="quantity">
              <label for="quantity">Quantity:</label>
              <input type="number" id="quantity" v-model="quantity" min="1" :max="product.stock">
            </div>
            <button class="btn btn-primary" @click="addToCart">Add to Cart</button>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: "Product",
    data() {
    return {
      product: {
        id: 1,
        name: "Product 1",
        price: 10.99,
        description: "This is a description of Product 1. Add more details as needed.",
        stock: 10  // Example stock quantity
      },
      quantity: 1
    };
  },
  computed: {
    imageUrl() {
      return `/images/product${this.product.id}.jpg`;
    }
  },
    methods: {
      addToCart() {
        // Implement cart functionality here, e.g., send the product and quantity to the cart
        console.log(`Added ${this.quantity} ${this.product.name} to cart`);
      }
    }
  };
  </script>
  
  <style scoped>
  .product-details {
    padding: 20px;
  }
  
  .product-title {
    font-size: 1.5rem;
    font-weight: bold;
    margin-bottom: 10px;
  }
  
  .product-price {
    font-size: 1.25rem;
    color: #007BFF;
    margin-bottom: 15px;
  }
  
  .product-description {
    font-size: 1rem;
    margin-bottom: 15px;
  }
  
  .product-stock {
    font-size: 1rem;
    margin-bottom: 15px;
  }
  
  .quantity {
    margin-bottom: 15px;
  }
  
  /* Additional styles as needed */
  </style>
  