<template>
    <form @submit.prevent="create" class="mt-4">
      <div class="form-group">
        <label for="name">Name:</label>
        <input type="text" v-model="name" class="form-control" id="name" required>
      </div>
      <div class="form-group">
        <label for="description">Description:</label>
        <textarea v-model="description" class="form-control" id="description" rows="4" required></textarea>
      </div>
      <div class="form-group">
        <label for="price">Price:</label>
        <input type="number" v-model="price" class="form-control" id="price" required>
      </div>
      <button type="submit" class="btn btn-primary">Create Product</button>
    </form>
    <div v-if="productCreated" class="alert alert-success mt-2">Product created successfully!</div>
    <div v-if="productError" class="alert alert-danger mt-2">Error creating the product. Please try again.</div>
  </template>
  
  <script>
  import { createProduct } from '@/js/products.js';
  
  export default {
    name: "ProductCreateForm",
    mixins: [createProduct],
    data() {
      return {
        name: "",
        description: "",
        price: "",
        productCreated: false,
        productError: false,
      };
    },
    methods: {
      async create() {
        const productData = {
          name: this.name,
          description: this.description,
          price: this.price,
        };
        try {
            await createProduct(productData);
            this.productCreated = true;
            this.productError = false;
            this.name = "";
            this.description = "";
            this.price = "";
        } catch (error) {
            console.error('Error creating product:', error);
            this.productError = true;
            this.productCreated = false;
        }
      },
    },
  };
  </script>
  
<style scoped>
body {
  margin-left: 0; 
  padding-left: 550px; 
}
</style>