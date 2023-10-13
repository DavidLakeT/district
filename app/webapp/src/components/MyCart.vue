<template>
    <div class="container mt-5">
      <h2>Your Cart</h2>
      <div class="row">
        <div class="col-md-8">
          <ul class="list-group">
            <li
              v-for="(item, index) in cart"
              :key="index"
              class="list-group-item d-flex justify-content-between align-items-center"
            >
              <span class="product-name">{{ item.product.name }}</span>
              <span class="badge badge-primary badge-pill quantity-badge">{{ item.quantity }}</span>
            </li>
          </ul>
        </div>
        <div class="col-md-4">
          <div class="card">
            <div class="card-body">
              <h5 class="card-title">Order Summary</h5>
              <p class="card-text">Total Items: {{ totalItems }}</p>
              <p class="card-text">Total Price: $ {{ totalPrice.toFixed(2) }}</p>
              <a href="#" class="btn btn-primary">Proceed to Checkout</a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: "MyCart",
    data() {
      return {
        cart: [
          { product: { name: "Product 1" }, quantity: 2 },
          { product: { name: "Product 2" }, quantity: 3 }
        ]
      };
    },
    computed: {
      totalItems() {
        return this.cart.reduce((total, item) => total + item.quantity, 0);
      },
      totalPrice() {
        return this.cart.reduce((total, item) => total + item.quantity * item.product.price, 0);
      }
    }
  };
  </script>
  
  <style scoped>
  .list-group-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 10px;
    background-color: #f9f9f9;
    border: none;
  }
  
  .product-name {
    font-size: 1.2rem;
  }
  
  .quantity-badge {
    font-size: 1.2rem;
    color: #000; /* Change to a color that contrasts with the card */
  }
  
  .card {
    border: none;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
    margin-bottom: 7%;
  }
  
  .card-title {
    font-size: 1.5rem;
    font-weight: bold;
  }
  
  .card-text {
    font-size: 1.25rem;
  }
  
  .btn-primary {
    background-color: #007BFF;
    border: none;
  }
  
  .btn-primary:hover {
    background-color: #0056b3;
  }
  </style>
  
  