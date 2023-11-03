<template>
  <div class="product-details">
    <h2 class="product-title">{{ product.name }}</h2>
    <p class="product-price">$ {{ product.price }}</p>
    <p class="product-description">{{ product.description }}</p>
    <p class="product-stock" v-if="product.stock > 0">In Stock: {{ product.stock }}</p>
    <p class="product-stock" v-else>Out of Stock</p>
    <div class="quantity">
      <label for="quantity">Quantity:</label>
      <input type="number" id="quantity" v-model="quantity" min="1" :max="product.stock">
    </div>
    <button class="btn btn-primary" @click="addToCart">Add to Cart</button>
  </div>
</template>

<script>
import { addCartItem } from '@/js/cart.js'
export default {
  props: {
    product: Object,
  },
  methods: {
    async addToCart(){
        const productData = { product_id: this.product.id, quantity: 1, price: this.product.price };
        const response = await addCartItem(productData);
        if (response) {
          this.$router.push('/mycart');
        }
    }
  },
};
</script>

<style scoped>
@import '@/css/productStyle.css';
</style>

  