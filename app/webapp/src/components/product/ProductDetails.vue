<template>
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
</template>

<script>
export default {
  props: {
    product: Object, // Pass the product data as a prop
  },
  data() {
    return {
      quantity: 1,
    };
  },
  methods: {
    addToCart() {
      // Implement cart functionality here, e.g., send the product and quantity to the cart
      console.log(`Added ${this.quantity} ${this.product.name} to cart`);
    },
  },
};
</script>

<style scoped>
@import '@/css/productStyle.css';
</style>

  