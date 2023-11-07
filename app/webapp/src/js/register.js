import axios from 'axios';

export default {
  name: "RegisterFetcher",
  data() {
    return {
      identification: '',
      email: '',
      username: '',
      password: '',
      address: ''
    };
  },

  methods: {
    async registerUser() {
      try {
        const isAdmin = false;
        const identificationInt = parseInt(this.identification, 10);
        await axios.post('http://localhost:5000/api/user', {
          identification: identificationInt,
          email: this.email,
          username: this.username,
          password: this.password,
          address: this.address,
          is_admin: isAdmin
        });
        this.$router.push('/login');
      } catch (error) {
        console.error('Error registering user:', error);
      }
    }
  }
};