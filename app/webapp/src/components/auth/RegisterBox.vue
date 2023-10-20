<template>
  <div>
    <h2 class="text-center">Register</h2>
    <form @submit.prevent="register" class="mt-4">
      <form-field
        label="Identification:"
        field-type="text"
        :field-value="identification"
        input-class="form-control"
        field-id="identification"
        placeholder="Enter your identification"
        @update:fieldValue="identification = $event"
      ></form-field>

      <form-field
        label="Full Name:"
        field-type="text"
        :field-value="username"
        input-class="form-control"
        field-id="username"
        placeholder="Enter your full name"
        @update:fieldValue="username = $event"
      ></form-field>

      <form-field
        label="Email:"
        field-type="email"
        :field-value="email"
        input-class="form-control"
        field-id="email"
        placeholder="Enter your email"
        @update:fieldValue="email = $event"
      ></form-field>

      <form-field
        label="Password:"
        field-type="password"
        :field-value="password"
        input-class="form-control"
        field-id="password"
        placeholder="Enter your password"
        @update:fieldValue="password = $event"
      ></form-field>

      <form-field
        label="Address:"
        field-type="text"
        :field-value="address"
        input-class="form-control"
        field-id="address"
        placeholder="Enter your address"
        @update:fieldValue="address = $event"
      ></form-field>

      <button @click="registerUser" type="submit" class="btn btn-primary btn-block">Register</button>
    </form>
    <div class="mt-3 text-center">
      Already have an account? <router-link to="/login">Login</router-link>
    </div>
  </div>
</template>

<script>
import registerFetcher from '@/js/register.js';
import FormField from '../common/FormField.vue';

export default {
  name: "RegisterBox",
  components: {
    FormField,
  },
  mixins: [registerFetcher],
  data() {
    return {
      identification: '',
      username: '',
      email: '',
      password: '',
      address: '',
    };
  },
};
</script>

<style scoped>
  @import '@/css/authStyle.css';
</style>
