<template>
    <div class="container mt-5">
      <h2>Create a New Product</h2>
      <ProductCreateForm @create="createProduct" />
    </div>
  </template>
  
  <script>
  import ProductCreateForm from '@/components/admin/ProductCreateForm.vue';
  
  export default {
    components: {
      ProductCreateForm,
    }
  };
  </script>
  
<style scoped>
  @import '@/css/authStyle.css'
</style>