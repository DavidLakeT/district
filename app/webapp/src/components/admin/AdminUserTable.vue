<template>
    <table class="table table-striped">
      <thead>
        <tr>
          <th>Username</th>
          <th>Email</th>
          <th>Is Admin</th>
          <th>Balance</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="user in users" :key="user.id">
          <td>{{ user.username }}</td>
          <td>{{ user.email }}</td>
          <td>
            <input type="checkbox" v-model="user.is_admin" @change="updateUser(user)">
          </td>
          <td>
            <input type="number" v-model="user.balance" @change="updateUser(user)">
          </td>
          <td>
            <button @click="deleteUser(user.id)">Delete</button>
          </td>
        </tr>
      </tbody>
    </table>
  </template>
  
  <script>
  export default {
    props: {
      users: Array,
    },
    methods: {
      updateUser(user) {
        this.$emit('updateUser', user);
      },
      deleteUser(userId) {
        this.$emit('deleteUser', userId);
      },
    },
  };
  </script>
  
  
<style scoped>
@import '@/css/productList.css';
</style>