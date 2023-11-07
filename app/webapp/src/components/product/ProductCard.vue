<template>
  <div class="card">
    <img :src="imageSrc" class="card-img-top" :alt="product.name">
    <div class="card-body">
      <h5 class="card-title">{{ product.name }}</h5>
      <p class="card-text">$ {{ product.price.toFixed(2) }}</p>
    </div>
  </div>
</template>

<script>
import { getProductImage } from '@/js/picture.js';

export default {
  name: "ProductCard",
  props: {
    product: Object,
  },
  data() {
    return {
      imageSrc: null,
    };
  },
  async created() {
    try {
      const response = await getProductImage(this.product.id);
      if (response) {
        const blob = new Blob([response], { type: 'image/jpeg' });
        this.imageSrc = URL.createObjectURL(blob);
      }
    } catch (error) {
      console.error('Error loading product image:', error);
    }
  },
};
</script>
