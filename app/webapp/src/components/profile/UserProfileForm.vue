<template>
  <div class="container mt-5">
      <div class="user-info">
        <div class="info-label">Identification:</div>
        <div class="info-value">{{ editedUser.identification }}</div>
      </div>
      <div class="user-info">
        <div class="info-label">Balance:</div>
        <div class="info-value">{{ editedUser.balance }}</div>
      </div>
    <form @submit.prevent="submit" class="mt-4">
      <form-field label="Username:" field-type="text" :field-value="editedUser.username" input-class="form-control" required @update:fieldValue="editedUser.username = $event"></form-field>
      <form-field label="Email:" field-type="email" :field-value="editedUser.email" input-class="form-control" required @update:fieldValue="editedUser.email = $event"></form-field>
      <form-field label="Address:" field-type="address" :field-value="editedUser.address" input-class="form-control" required @update:fieldValue="editedUser.address = $event"></form-field>
      <button type="submit" class="btn btn-primary">Update Profile</button>
    </form>
  </div>
</template>

<script>
import FormField from '../common/FormField.vue';
import profileFetcher from '@/js/profile.js';
export default {
  components: {
    FormField,
  },
  mixins: [profileFetcher],
};
</script>