<template>
  <div class="card">
    <ImageLoader :productId="product.id" :altText="product.name" class="card-img-top" />
    <div class="card-body">
      <h5 class="card-title">{{ product.name }}</h5>
      <p class="card-text">$ {{ product.price.toFixed(2) }}</p>
    </div>
  </div>
</template>

<script>
import ImageLoader from '@/components/common/ImageLoader.vue';
export default {
  name: "ProductCard",
  components:{
    ImageLoader
  },
  props: {
    product: Object,
  },
};
</script>
