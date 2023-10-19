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
  
  export default {
    name: "ProductView",
    components: {
      ProductDetails,
      ProductReviews,
    },
    data() {
      return {
        product: {
          id: 1,
          name: "Product 1",
          price: 10.99,
          description: "This is a description of Product 1. Add more details as needed.",
          stock: 10,
        },
        reviews: [
          { author: "User 1", rating: 4, text: "This product is great!", date: "2023-08-27" },
          { author: "User 2", rating: 5, text: "Amazing quality!", date: "2023-08-28" },
          { author: "User 3", rating: 3, text: "Amazing quality!", date: "2023-08-28" },
          // Add more reviews as needed
        ],
      };
    },
    computed: {
      imageUrl() {
        return `/images/product${this.product.id}.jpg`;
      },
    },
  };
  </script>
  
  <style scoped>
  /* View-specific styles */
  </style>
  