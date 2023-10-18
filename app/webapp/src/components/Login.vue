<template>
  <div class="container mt-5">
    <h2 class="text-center">Login</h2>
    <form @submit.prevent="login" class="mt-4">
      <div class="form-group">
        <label for="email">Email:</label>
        <input type="email" v-model="email" class="form-control" id="email" placeholder="Enter your email" required>
      </div>
      <div class="form-group">
        <label for="password">Password:</label>
        <input type="password" v-model="password" class="form-control" id="password" placeholder="Enter your password" required>
      </div>
      <button @click="loginUser" type="submit" class="btn btn-primary btn-block">Login</button>
    </form>
    <div class="mt-3 text-center">
      Don't have an account? <router-link to="/register">Create one</router-link>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      email: '',
      password: ''
    };
  },
  methods: {
    async loginUser() {
      try {
        const response = await axios.post('http://localhost:5000/api/auth/login', {
          email: this.email,
          password: this.password
        }, {
          maxRedirects: 0,
          validateStatus: function (status) {
            return (status >= 200 && status < 400) || (status >= 300 && status < 400);
          }
        });

        if (response.status === 302 || response.data.auth_token) {
          console.log('User logged in successfully');

          const token = response.data.auth_token;

          console.log('Token:', token);

          localStorage.setItem('auth_token', token);

          this.$router.push('/products');
        } else {
          console.error('Error logging in dggdgdg:', response.data);
        }
      } catch (error) {
        if (error.response) {
          console.error('Error logging in dddd:', error.response.data);
        } else if (error.request) {
          console.error('Error logging in xxxx:', 'No response received');
        } else {
          console.error('Error :', error.message);
        }
      }
    }
  }
};
</script>

<style scoped>
.container {
  margin-bottom: 17%;
}

.form-group {
  margin-bottom: 1rem;
}

.btn-primary {
  background-color: #007BFF;
  border: none;
}

.btn-primary:hover {
  background-color: #0056b3;
}

.btn-block {
  width: 100%;
}

.text-center {
  text-align: center;
}
</style>
