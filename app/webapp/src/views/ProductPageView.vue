<template>
  <div class="container mt-5">
    <div class="row">
      <div class="col-md-6">
        <img :src="imageUrl" alt="Product Image" class="img-fluid" />
      </div>
      <div class="col-md-6">
        <product-details :product="product" />
      </div>
    </div>

    <product-reviews :reviews="reviews" />
  </div>
</template>

<script>
import ProductDetails from '@/components/product/ProductDetails.vue';
import ProductReviews from '@/components/product/ProductReviews.vue';
import { fetchProductDetails } from '@/js/products.js';

export default {
  name: "ProductView",
  components: {
    ProductDetails,
    ProductReviews,
  },
  data() {
    return {
      product: {},
      reviews: [],
    };
  },
  computed: {
    /*imageUrl() {
      return `/images/product${this.product.id}.jpg`;
    },*/
  },
  async created() {
    try {
      const productName = this.$route.params.id;
      const data = await fetchProductDetails(productName);
      this.product = data.product;
      this.reviews = data.reviews;
    } catch (error) {
      console.error('Error fetching product data:', error);
    }
  },
};
</script>

<style scoped>
/* View-specific styles */
</style>

