<template>
  <div class="container mt-5">
    <h2>Your Profile</h2>
    <UserProfile />
  </div>
</template>

<script>
import UserProfile from '@/components/profile/UserProfileForm.vue';

export default {
  components: {
    UserProfile,
  },
};
</script>

<style scoped>
@import '@/css/authStyle.css';
</style>
