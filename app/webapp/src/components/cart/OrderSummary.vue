<template>
  <div class="card">
    <div class="card-body">
      <h5 class="card-title">Order Summary</h5>
      <p class="card-text">Total Items: {{ totalItems }}</p>
      <p class="card-text">Total Price: $ {{ totalPrice }}</p>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    cart: {
      type: Array,
      required: true,
    },
    totalPrice: {
      type: Number,
      required: true,
    },
  },
  computed: {
    totalItems() {
      return this.cart.reduce((total, item) => total + item.quantity, 0);
    },
  },
};
</script>
