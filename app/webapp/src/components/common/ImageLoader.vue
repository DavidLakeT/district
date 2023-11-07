<template>
    <img :src="imageSrc" :alt="altText" />
  </template>
  
  <script>
  import { getProductImage } from '@/js/picture.js';
  
  export default {
    props: {
      src: String,
      productId: Number,
      altText: String, 
    },
    data() {
      return {
        imageSrc: null,
      };
    },
    async created() {
      if (this.src) {
        this.imageSrc = this.src;
      } else if (this.productId) {
        try {
          const response = await getProductImage(this.productId);
          if (response) {
            const blob = new Blob([response], { type: 'image/jpeg' });
            this.imageSrc = URL.createObjectURL(blob);
          }
        } catch (error) {
          console.error('Error loading product image:', error);
        }
      }
    },
  };
  </script>
  