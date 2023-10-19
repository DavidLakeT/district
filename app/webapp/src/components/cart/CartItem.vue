<template>
    <li class="list-group-item d-flex justify-content-between align-items-center">
      <span class="product-name">{{ item.product.name }}</span>
      <span class="badge badge-primary badge-pill quantity-badge">{{ item.quantity }}</span>
    </li>
  </template>
  
<script>
  export default {
    props: {
      item: {
        type: Object,
        required: true,
      },
    },
  };
</script>
  