<template>
    <nav id="admin-sidebar">
      <div class="admin-links" :class="{ 'show': adminLinksVisible }" id="adminNavbar">
        <ul class="nav flex-column">
          <li class="nav-item">
            <button class="nav-link admin-button" @click="navigateTo('/')">
              Home
            </button>
          </li>
          <li class="nav-item">
            <button class="nav-link admin-button" @click="navigateTo('/profile/' + userId)">
              My Profile
            </button>
          </li>
          <li class="nav-item">
            <button class="nav-link admin-button" @click="navigateTo('/admin/products/create')">
              Create Products
            </button>
          </li>
          <li class="nav-item">
            <button class="nav-link admin-button" @click="navigateTo('/admin/products/')">
              Manage Products
            </button>
          </li>
          <li class="nav-item">
            <button class="nav-link admin-button">
              User Management
            </button>
          </li>
          <li class="nav-item" v-if="isAuthenticated">
            <button class="nav-link admin-button" @click="logout">Logout</button>
          </li>
        </ul>
      </div>
    </nav>
  </template>
  
  <script>
  import { mapMutations } from 'vuex';
  
  export default {
    name: "AdminNavbar",
    data() {
      return {
        adminLinksVisible: true,
      };
    },
    methods: {
      navigateTo(route) {
        this.$router.push(route);
      },
      ...mapMutations('auth', ['logout']),
        logout() {
            this.$store.commit('auth/logout');
            document.cookie = 'auth_token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/';
            this.$router.push('/login');
        }
    },
    computed: {
      isAuthenticated() {
        return this.$store.getters['auth/isAuthenticated'];
      },
      userId() {
        return this.$store.getters['auth/userId'];
      }
    },
  };
  </script>
  
  
  <style scoped>
  #admin-sidebar {
    width: 250px;
    position: fixed;
    top: 0;
    left: 0;
    height: 100%;
    background-color: #333;
    transition: left 0.3s ease;
  }
  
  .admin-navbar {
    display: flex;
    justify-content: space-between;
    padding: 10px;
    background-color: #222;
  }
  
  .admin-button {
    text-align: left;
    background: none;
    border: none;
    color: #fff;
    font-size: 16px;
    padding: 10px 20px; /* Increase the clickable area */
    cursor: pointer;
    width: 100%;
  }
  
  .admin-button:hover {
    background-color: #444; /* Darker background on hover */
  }
  
  .admin-links {
    padding: 10px;
  }
  
  .admin-links ul {
    list-style: none;
    padding: 0;
    margin: 0;
  }
  
  .admin-links li {
    margin-bottom: 10px;
  }
  .nav-link.admin-button {
    color: #fff; /* Set the text color to white */
    text-decoration: none; /* Remove underline on hover */
  }

  .nav-link.admin-button:hover {
    color: #fff; /* Set the text color to white on hover */
  }
  </style>
  