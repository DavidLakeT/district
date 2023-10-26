<template>
    <tr>
      <td>{{ product.name }}</td>
      <td>${{ product.price.toFixed(2) }}</td>
      <td>
        <button @click="redirectToProductUpdate" class="btn btn-info btn-sm">Update</button>
        <button @click="deleteProduct" class="btn btn-danger btn-sm">Delete</button>
      </td>
    </tr>
  </template>
  
  <script>
  export default {
    props: {
      product: Object,
    },
    methods:{
        async deleteProduct(){
            console.log("Product deleted")
        },
        redirectToProductUpdate() {
            const productId = this.product.id;
            this.$router.push('/admin/products/'+productId);
        },
    }

  };
  </script>
  