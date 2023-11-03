<template>
    <div class="card">
      <img :src="imageUrl(product.id)" class="card-img-top" :alt="product.name">
      <div class="card-body">
        <h5 class="card-title">{{ product.name }}</h5>
        <p class="card-text">$ {{ product.price.toFixed(2) }}</p>
        <button class="btn btn-primary" @click="addToCart">Add to Cart</button>
      </div>
    </div>
  </template>
  
  <script>
  import { addCartItem } from '@/js/cart.js'
  import router from '@/router';

  export default {
    name: "ProductCard",
    props: {
      product: Object,
    },
    computed: {
      imageUrl() {
        return (productId) => `/images/product${productId}.jpg`;
      },
      methods:{
        async addToCart(){
          const productData = { product_id: this.product.id, quantity: 1, price: this.product.price };
          const response = await addCartItem(productData);
          if (response) {
            router.push('/products')
          }
        }
      }
    },
  };
  </script>
  