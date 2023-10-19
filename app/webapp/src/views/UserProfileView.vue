<template>
    <div class="container mt-5">
      <h2>Your Profile</h2>
      <UserProfileForm v-if="editing" @submit="updateUserProfile" :user="user" />
      <UserProfileDisplay v-else :user="user" @edit="startEditing" />
    </div>
  </template>
  
<script>
  import UserProfileForm from '@/components/profile/UserProfileForm.vue';
  import UserProfileDisplay from '@/components/profile/UserProfileDisplay.vue';
  import profileFetcher from '@/js/profile.js';
  export default {
    components: {
        UserProfileForm,
        UserProfileDisplay
    },
    mixins: [profileFetcher]
 }
</script>
  
  <style scoped>
    @import '@/css/profileStyle.css';
  </style>
  