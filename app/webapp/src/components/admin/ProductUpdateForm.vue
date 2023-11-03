<template>
  <div class="product-details">
    <h2 class="product-title">Update Product: {{ product.name }}</h2>

    <form @submit.prevent="updateProduct">
      <form-field label="Name:" field-type="text" :field-value="product.name" input-class="form-control" required @update:fieldValue="updatedProduct.name = $event"></form-field>
      <form-field label="Price:" field-type="number" :field-value="product.price" input-class="form-control" required @update:fieldValue="updatedProduct.price = $event"></form-field>
      <form-field label="Description:" field-type="textarea" :field-value="product.description" input-class="form-control" required @update:fieldValue="updatedProduct.description = $event"></form-field>
      <form-field label="Stock:" field-type="number" :field-value="product.stock" input-class="form-control" required @update:fieldValue="updatedProduct.stock = $event"></form-field>
      <button type="submit" class="btn btn-primary">Update Product</button>
    </form>
  </div>
</template>

<script>
import FormField from '@/components/common/FormField.vue';
import { updateProduct } from '@/js/products.js';

export default {
  props: {
    product: Object,
  },
  components: {
    FormField,
  },
  data() {
    return {
      updatedProduct: this.product,
    };
  },
  methods: {
    async updateProduct() {
      this.updatedProduct.stock =  parseInt(this.updatedProduct.stock, 10);
      this.updatedProduct.price =  parseInt(this.updatedProduct.price, 10);
      const response = await updateProduct(this.updatedProduct, this.product.id);
      if (response) {
        window.location.reload();
      }
    },
  },
};
</script>
