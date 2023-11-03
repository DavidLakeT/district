<template>
  <li class="list-group-item d-flex justify-content-between align-items-center">
    <span class="product-name">{{ productName }}</span>
    <span class="text-dark">{{ item.quantity }}</span>
  </li>
</template>

<script>
import { fetchProductDetails } from '@/js/products';

export default {
  props: {
    item: {
      type: Object,
      required: true,
    },
  },
  data() {
    return {
      productName: null,
    };
  },
  async created() {
    if (this.item.product_id) {
      const productDetails = await fetchProductDetails(this.item.product_id);
      this.productName = productDetails.name;
    }
  },
};
</script>



