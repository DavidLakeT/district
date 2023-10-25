<template>
    <div class="container">
      <h2>Product List</h2>
      <table class="table table-striped">
        <thead>
          <tr>
            <th>Name</th>
            <th>Price</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          <product-table-row
            v-for="product in products"
            :key="product.id"
            :product="product"
          ></product-table-row>
        </tbody>
      </table>
    </div>
  </template>
  
  <script>
  import ProductTableRow from "@/components/admin/ProductTableRow.vue";
  
  export default {
    props: {
      products: Array,
    },
    components: {
      ProductTableRow,
    },
  };
  </script>

<style scoped>
  @import '@/css/productList.css'
</style>
  