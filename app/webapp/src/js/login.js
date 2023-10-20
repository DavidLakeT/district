import axios from 'axios';

export default {
  name: "LoginFetcher",
  data() {
    return {
      email: '',
      password: '',
      loginError: null,
    };
  },
  methods: {
    async loginUser() {
      this.loginError = null;

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
          document.cookie = `auth_token=${response.data.auth_token}; path=/`;
          this.$router.push('/products');
        } else {
          this.loginError = 'Incorrect credentials. Please try again.';
        }
      } catch (error) {
        if (error.response) {
          this.loginError = 'Incorrect credentials. Please try again.';
        } else if (error.request) {
          this.loginError = 'Network error. Please try again.';
        } else {
          this.loginError = 'An error occurred. Please try again.';
        }
      }
    }
  }
};