<template>
    <div class="container mt-5">
      <div class="row">
        <div class="col-md-10 offset-md-1">
          <ProductList :products="products" />
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import ProductList from '@/components/product/ProductList.vue';
  
  export default {
    name: "Products",
    components: {
      ProductList,
    },
    data() {
        return {
            products: [
            { id: 1, name: "Product 1", price: 10.99 },
            { id: 2, name: "Product 2", price: 15.99 },
            ]
        };
    }
  };
  </script>
  