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
        const identificationInt = parseInt(this.identification, 10);
        const response = await axios.post('http://localhost:5000/api/user', {
          identification: identificationInt,
          email: this.email,
          username: this.username,
          password: this.password,
          address: this.address,
          isadmin: false
        });
        console.log('User registered:', response.data);
        this.$router.push('/login');
      } catch (error) {
        console.error('Error registering user:', error);
      }
    }
  }
};