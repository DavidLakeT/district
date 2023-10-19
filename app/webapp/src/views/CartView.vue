<template>
    <div class="container mt-5">
      <h2>Your Cart</h2>
      <div class="row">
        <div class="col-md-8">
          <ul class="list-group">
            <cart-item v-for="(item, index) in cart" :key="index" :item="item" />
          </ul>
        </div>
        <div class="col-md-4">
          <order-summary :cart="cart" />
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import CartItem from '@/components/cart/CartItem.vue';
  import OrderSummary from '@/components/cart/OrderSummary.vue';
  
  export default {
    name: "MyCart",
    components: {
      CartItem,
      OrderSummary,
    },
    data() {
      return {
        cart: [
          { product: { name: "Product 1" }, quantity: 2 },
          { product: { name: "Product 2" }, quantity: 3 },
          { product: { name: "Product 2" }, quantity: 3 }
        ],
      };
    },
  };
  </script>
  
  <style scoped>
  @import '@/css/cartStyle.css';
  </style>
  