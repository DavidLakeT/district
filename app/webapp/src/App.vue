<template>
  <div>
    <AdminNavbar v-if="isAuthenticated && userRole" />
    <Navbar v-else />
    <div class="content-area" v-if="isAuthenticated && userRole">
      <router-view />
    </div>
    <router-view v-else />
    <Footer />
  </div>
</template>

<script>
import Navbar from '@/components/common/Navbar.vue';
import Footer from '@/components/common/Footer.vue';
import AdminNavbar from '@/components/common/AdminNavbar.vue';


export default {
  components: {
    Navbar,
    AdminNavbar,
    Footer
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
}
</script>

<style scoped>
.content-area {
  margin-left: 20%;
}
</style>
