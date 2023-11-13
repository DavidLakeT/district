<template>
    <div class="search-product">
      <div class="input-group">
        <input v-model="productName" type="text" class="form-control" placeholder="Search for a product..." />
        <button @click="search" class="btn btn-primary">Search</button>
      </div>
    </div>
  </template>
  
  <script>
  import { searchProduct, fetchProducts  } from '@/js/products.js';
  
  export default {
    data() {
      return {
        productName: '',
      };
    },
    methods: {
      async search() {
        if (this.productName) {
          const products = await searchProduct(this.productName);
          this.$emit('search-results', products);
        } else {
            const allProducts = await fetchProducts();
            this.$emit('search-results', allProducts);
        }
      },
    },
  };
  </script>
  
<style scoped>
  @import '@/css/searchProduct.css';
</style>