<template>
  <li class="list-group-item d-flex justify-content-between align-items-center">
    <span class="product-name">{{ productName }}</span>
    <div class="row align-items-center">
      <div class="col-2">
        <button @click="decrementQuantity" class="btn btn-sm btn-secondary">-</button>
      </div><div class="col-2"></div>
      <div class="col-4">
        <span class="quantity">{{ item.quantity }}</span>
      </div>
      <div class="col-2">
        <button @click="incrementQuantity" class="btn btn-sm btn-secondary">+</button>
      </div>
    </div>
    <button @click="removeItem" class="btn btn-danger">Remove</button>
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
  methods: {
    incrementQuantity() {
      this.$emit('update-quantity', this.item.product_id, this.item.quantity + 1);
    },
    decrementQuantity() {
      if (this.item.quantity > 1) {
        this.$emit('update-quantity', this.item.product_id, this.item.quantity - 1);
      }
    },
    removeItem() {
      this.$emit('remove-item', this.item.product_id);
    },
  },
};
</script>



