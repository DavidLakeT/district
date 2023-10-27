<template>
  <div class="container mt-5">
    <div class="row">
      <div class="col-md-10 offset-md-1">
        <SearchProduct @search-results="handleSearchResults" />
        <ProductList :products="products" />
      </div>
    </div>
  </div>
</template>

<script>
import ProductList from '@/components/product/ProductList.vue';
import SearchProduct from '@/components/product/SearchProducts.vue';
import { fetchProducts } from '@/js/products.js';

export default {
  name: "Products",
  components: {
    ProductList,
    SearchProduct,
  },
  data() {
    return {
      products: [],
    };
  },
  async created() {
    this.products = await fetchProducts();
  },
  methods: {
    handleSearchResults(products) {
      this.products = products;
    },
  },
};
</script>

  