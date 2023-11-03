<template>
    <tr>
      <td>{{ product.name }}</td>
      <td>${{ product.price.toFixed(2) }}</td>
      <td>
        <button @click="redirectToProductUpdate" class="btn btn-info btn-sm">Update</button>
        <button @click="deleteP" class="btn btn-danger btn-sm">Delete</button>
      </td>
    </tr>
  </template>
  
  <script>
  import { deleteProduct } from '@/js/products';
  export default {
    props: {
      product: Object,
    },
    methods:{
        async deleteP(){
          const response = await deleteProduct(this.product.id)
          if(response){
            window.location.reload();
          }
        },
        redirectToProductUpdate() {
            const productId = this.product.id;
            this.$router.push('/admin/products/'+productId);
        },
    }

  };
  </script>
  