<template>
    <div class="product-details">
      <h2 class="product-title">Update Product: {{ product.name }}</h2>
  
      <form @submit.prevent="updateProduct">
        <form-field :field-value="product.name" label="Name" name="name" :required="true" />
        <form-field :field-value="product.price" label="Price" name="price" type="number" :required="true" />
        <form-field :field-value="product.description" label="Description" name="description" type="textarea" :required="true" />
        <form-field :field-value="product.stock" label="Stock" name="stock" type="number" :required="true" />
  
        <button type="submit" class="btn btn-primary">Update Product</button>
      </form>
    </div>
  </template>
  
  <script>
  import FormField from '@/components/common/FormField.vue';
  
  export default {
    props: {
      product: Object,
    },
    components: {
      FormField,
    },
    data() {
      return {
        updatedProduct: { ...this.product }, // Make a copy of the prop
      };
    },
    methods: {
      updateProduct() {
        // Work with the updatedProduct data to perform the update
        // ...
      },
    },
  };
  </script>