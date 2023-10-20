<template>
    <div class="card">
      <div class="card-body">
        <h5 class="card-title">Order Summary</h5>
        <p class="card-text">Total Items: {{ totalItems }}</p>
        <p class="card-text">Total Price: $ {{ totalPrice.toFixed(2) }}</p>
        <a href="#" class="btn btn-primary">Proceed to Checkout</a>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      cart: {
        type: Array,
        required: true,
      },
    },
    computed: {
      totalItems() {
        return this.cart.reduce((total, item) => total + item.quantity, 0);
      },
      totalPrice() {
        return this.cart.reduce((total, item) => total + item.quantity * item.product.price, 0);
      },
    },
  };
  </script>
  