<template>
  <div class="container mt-5">
      <h2>Admin User Management</h2>
      <admin-user-table :users="users" @updateUser="updateUser" @deleteUser="deleteUser" />
    </div>
  </template>
  
  <script>
  import { getUsers, updateUser, deleteUser } from '@/js/userManagement.js';
  import AdminUserTable from '@/components/admin/AdminUserTable.vue';

  
  export default {
    data() {
      return {
        users: [],
      };
    },
    components: {
      AdminUserTable,
    },
    created() {
      this.loadUsers();
    },
    methods: {
      async loadUsers() {
        this.users = await getUsers(); 
      },
      async updateUser(user) {
        const response = await updateUser(user, user.identification); 
        if (response) {
          //  success
        } else {
          //  error
        }
      },
      async deleteUser(userId) {
        const response = await deleteUser(userId); 
        if (response) {
          //  success
          this.loadUsers();
        } else {
          //  error
        }
      },
    },
  };
  </script>
  