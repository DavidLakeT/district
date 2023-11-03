<template>
  <div class="container mt-5">
    <h2>Your Cart</h2>
    <div class="row">
      <div class="col-md-8">
        <ul class="list-group">
          <cart-item v-for="(item, index) in cart.cart" :key="index" :item="item" @update-quantity="updateQuantity" @remove-item="removeItem" />
        </ul>
      </div>
      <div class="col-md-4">
        <order-summary :cart="cart.cart" :total-price="cart.total_price" />
        <button @click="placeOrder" class="btn btn-primary mt-2">Place Order</button>
      </div>
    </div>
  </div>
</template>

<script>
import CartItem from '@/components/cart/CartItem.vue';
import OrderSummary from '@/components/cart/OrderSummary.vue';
import { getCartInformation, placeOrder, updateProductQuantity, removeItemFromCart } from '@/js/cart.js';
import router from '@/router';

export default {
  name: "MyCart",
  components: {
    CartItem,
    OrderSummary,
  },
  data() {
    return {
      cart: {
        cart: [], 
        total_price: 0, 
      },
    };
  },
  async created() {
    this.cart = await getCartInformation();
  },
  methods: {
    async placeOrder() {
      try {
        const response = await placeOrder();
        if (response) {
          router.push('/products');
        }
      } catch (error) {
        console.error('Error placing order:', error);
      }
    },
    async updateQuantity(itemId, newQuantity) {
      try {
        const cartItemData = { product_id: itemId, quantity: newQuantity };
        const response = await updateProductQuantity(cartItemData);
        if (response) {
          router.push('/cart');
        }
      } catch (error) {
        console.error('Error updating product quantity:', error);
      }
    },
    async removeItem(itemId) {
      try {
        const productID = itemId;
        const response = await removeItemFromCart(productID);
        if (response) {
          router.push('/cart');
        }
      } catch (error) {
        console.error('Error removing item from the cart:', error);
      }
    },
  },
};
</script>

<style scoped>
@import '@/css/cartStyle.css';
</style>

