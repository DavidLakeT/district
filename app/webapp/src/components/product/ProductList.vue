<template>
  <div class="row">
    <div v-for="product in products" :key="product.id" class="col-md-4 mb-4">
      <router-link :to="'/products/' + product.id" class="card-link">
        <ProductCard :product="product" />
      </router-link>
    </div>
  </div>
</template>

<script>
import ProductCard from '@/components/product/ProductCard.vue';

export default {
  name: "ProductList",
  props: {
    products: Array,
  },
  components: {
    ProductCard,
  },
};
</script>

<style scoped>
  @import '@/css/productsStyle.css';
</style>
  
  