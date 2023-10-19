<template>
    <form @submit.prevent="submit" class="mt-4">
      <div class="form-group">
        <label for="name">Name:</label>
        <input type="text" v-model="editedUser.name" class="form-control" id="name" required>
      </div>
      <div class="form-group">
        <label for="email">Email:</label>
        <input type="email" v-model="editedUser.email" class="form-control" id="email" required>
      </div>
      <button type="submit" class="btn btn-primary">Update Profile</button>
    </form>
  </template>
  
  <script>
  export default {
    props: {
      user: {
        type: Object,
        required: true,
      },
    },
    data() {
      return {
        editedUser: { ...this.user },
      };
    },
    methods: {
      submit() {
        this.$emit('submit', this.editedUser);
      },
    },
  };
  </script>