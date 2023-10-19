<template>
  <div>
    <h2 class="text-center">Login</h2>
    <form @submit.prevent="login" class="mt-4">
      <form-field
        label="Email:"
        field-type="email"
        :field-value="email"
        input-class="form-control"
        field-id="email"
        placeholder="Enter your email"
        required
        @update:fieldValue="email = $event"
      ></form-field>

      <form-field
        label="Password:"
        field-type="password"
        :field-value="password"
        input-class="form-control"
        field-id="password"
        placeholder="Enter your password"
        required
        @update:fieldValue="password = $event"
      ></form-field>

      <button @click="loginUser" type="submit" class="btn btn-primary btn-block">Login</button>
    </form>
    <div class="mt-3 text-center">
      Don't have an account? <router-link to="/register">Create one</router-link>
    </div>
    <div v-if="loginError" class="alert alert-danger mt-3 text-center">
      {{ loginError }}
    </div>
  </div>
</template>

<script>
import loginFetcher from '@/js/login.js';
import FormField from './FormField.vue';

export default {
  name: "LoginBox",
  components: {
    FormField,
  },
  mixins: [loginFetcher],
  data() {
    return {
      email: '',
      password: '',
      loginError: null,
    };
  },
};
</script>

<style scoped>
  @import '@/css/authStyle.css';
</style>
