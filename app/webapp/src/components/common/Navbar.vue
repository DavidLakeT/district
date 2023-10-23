<template>
  <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
    <a class="navbar-brand" href="#">District VulnWeb</a>
    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse justify-content-end" id="navbarNav">
      <ul class="navbar-nav">
        <li class="nav-item">
          <router-link to="/" class="nav-link">Home</router-link>
        </li>
        <li class="nav-item">
          <router-link to="/products" class="nav-link">Products</router-link>
        </li>
        <li class="nav-item">
          <router-link to="/mycart" class="nav-link" v-if="isAuthenticated && userRole != true">Cart</router-link>
        </li>
        <li class="nav-item">
          <router-link to="/admin/product/create" class="nav-link" v-if="isAuthenticated && userRole == true">Create</router-link>
        </li>
        <li class="nav-item">
          <router-link v-if="isAuthenticated" :to="'/profile/' + userId" class="nav-link">My Profile</router-link>
        </li>
        <li class="nav-item">
          <router-link to="/login" class="nav-link" v-if="!isAuthenticated">Sign in</router-link>
        </li>
        <li class="nav-item" v-if="isAuthenticated" >
          <LogoutButton />
        </li>
      </ul>
    </div>
  </nav>
</template>

<script>
import LogoutButton from '@/components/auth/LogoutButton.vue';

export default {
  name: "Navbar",
  components: {
    LogoutButton,
  },
  computed: {
    isAuthenticated() {
      return this.$store.getters['auth/isAuthenticated'];
    },
    userId() {
      return this.$store.getters['auth/userId'];
    },
    userRole() {
      return this.$store.getters['auth/isAdmin'];
    },
  },
};
</script>

<style scoped>
  @import '@/css/nav.css'
</style>
  