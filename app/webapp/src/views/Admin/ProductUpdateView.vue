<template>
    <div class="container mt-5">
      <div class="row">
        <div class="col-md-6">
          <ImageLoader :productId="this.$route.params.id" class="img-fluid" />
        </div>
        <div class="col-md-6">
          <ProductUpdateForm :product="product" />
        </div>
      </div>
      <product-reviews :reviews="reviews" :product="product" />
    </div>
</template>
  
  <script>
  import ProductUpdateForm from '@/components/admin/ProductUpdateForm.vue';
  import ProductReviews from '@/components/product/ProductReviews.vue';
  import ImageLoader from '@/components/common/ImageLoader.vue';
  import { fetchProductDetails } from '@/js/products.js';
  
  export default {
    name: "ProductUpdateView",
    components: {
      ProductUpdateForm,
      ProductReviews,
      ImageLoader,
    },
    data() {
      return {
        product: {},
        reviews: [],
      };
    },
    async created() {
      try {
        const product = await fetchProductDetails(this.$route.params.id);
        this.product = product;
        this.reviews = product.reviews;
      } catch (error) {
        //console.error('Error fetching product data:', error);
      }
    },
  };
  </script>

