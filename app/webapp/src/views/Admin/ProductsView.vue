<template>
    <div class="container mt-5">
      <div class="row">
        <div class="col-md-10 offset-md-1">
          <ProductList :products="products" />
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import ProductList from '@/components/admin/ProductList.vue';
  import { fetchProducts } from '@/js/products.js'; 
  
  export default {
    name: "Products",
    components: {
      ProductList,
    },
    data() {
      return {
        products: [],
      };
    },
    async created() {
      this.products = await fetchProducts();
    },
  };
  </script>