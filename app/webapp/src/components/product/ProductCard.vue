<template>
    <div class="card">
      <img :src="imageUrl(product.id)" class="card-img-top" :alt="product.name">
      <div class="card-body">
        <h5 class="card-title">{{ product.name }}</h5>
        <p class="card-text">$ {{ product.price.toFixed(2) }}</p>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: "ProductCard",
    props: {
      product: Object,
    },
    computed: {
      imageUrl() {
        return (productId) => `/images/product${productId}.jpg`;
      },
    },
  };
  </script>
  