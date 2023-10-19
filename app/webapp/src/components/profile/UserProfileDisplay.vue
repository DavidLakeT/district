<template>
    <div>
      <div class="form-group">
        <label>Name:</label>
        <p>{{ user.name }}</p>
      </div>
      <div class="form-group">
        <label>Email:</label>
        <p>{{ user.email }}</p>
      </div>
      <button @click="edit" class="btn btn-primary">Edit Profile</button>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      user: {
        type: Object,
        required: true,
      },
    },
    methods: {
      edit() {
        this.$emit('edit');
      },
    },
  };
  </script>
  
  <style scoped>
  /* Your styles for this component */
  </style>
  