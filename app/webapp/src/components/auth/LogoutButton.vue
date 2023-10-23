<template>
    <button @click="logout" class="nav-link">Logout</button>
  </template>
  
  <script>
  export default {
    methods: {
      logout() {
        this.$store.commit('auth/logout');
        document.cookie = 'auth_token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/';
        this.$router.push('/login');
      },
    },
  };
  </script>
  