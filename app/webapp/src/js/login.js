import axios from 'axios';
import store from '@/store';

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
          this.$root.authToken = response.data.auth_token;
          document.cookie = `auth_token=${response.data.auth_token}; path=/;`;
          
          const user = decodeAuthToken(response.data.auth_token);
          if (user) {
            store.commit('auth/setAuthToken', response.data.auth_token);
            store.commit('auth/setUserId', user.userId);
            store.commit('auth/setUserRole', !user.isAdmin);
            console.log(user.username)
            if (!user.isAdmin){
              this.$router.push('/admin/product/create');
            } else {
              this.$router.push('/products');
            }
          }
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
    },
  }
};

function decodeAuthToken(authToken) {
  try {
    const decoded = atob(authToken);
    const [userId, username, email, isAdmin] = decoded.split(':');
    return {
      userId,
      username,
      email,
      isAdmin: isAdmin === 'true',
    };
  } catch (error) {
    console.error('Error decoding auth token:', error);
    return null;
  }
}