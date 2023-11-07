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
    <div class="form-group">
      <label for="image">Image (JPG only):</label>
      <div class="custom-file">
        <input type="file" class="form-control-file" accept=".jpg" @change="handleImageUpload" id="image">
      </div>
    </div>
    
    <button type="submit" class="btn btn-primary">Create Product</button>
  </form>
  <div v-if="productCreated" class="alert alert-success mt-2">Product created successfully!</div>
  <div v-if="productError" class="alert alert-danger mt-2">Error creating the product. Please try again.</div>
</template>

<script>
import { createProduct } from '@/js/products.js';
import { uploadProductImage } from '@/js/picture.js';

export default {
  name: "ProductCreateForm",
  data() {
    return {
      name: "",
      description: "",
      price: "",
      imageFile: null,
      productCreated: false,
      productError: false,
    };
  },
  methods: {
    async create() {
      try {
        const productData = {
          name: this.name,
          description: this.description,
          price: this.price,
        };
        const productResponse = await createProduct(productData);
        const productId = productResponse.id;

        if (this.imageFile) {
          const imageName = productId + this.imageFile.name.substring(this.imageFile.name.lastIndexOf('.'));
          const newImageFile = new File([this.imageFile], imageName, { type: this.imageFile.type });
          await uploadProductImage(newImageFile);
        }

        this.productCreated = true;
        this.productError = false;
        this.name = "";
        this.description = "";
        this.price = "";
        this.imageFile = null;
      } catch (error) {
        this.productError = true;
        this.productCreated = false;
      }
    },
    handleImageUpload(event) {
      this.imageFile = event.target.files[0];
    },
  },
};
</script>