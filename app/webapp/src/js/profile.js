import axios from 'axios';

export default {
    data() {
        return {
          editedUser: {},
        };
      },
      created() {
        const userId = parseInt(this.$route.params.userId, 10);
        this.fetchUserProfile(userId);
      },
      methods: {
        async fetchUserProfile(userId) {
          try {
            const response = await axios.get(`http://localhost:5000/api/user/${userId}`);
            this.editedUser = { ...response.data };
          } catch (error) {
            console.error('Error fetching user profile:', error);
          }
        },
        async submit() {
          try {
            const userId = parseInt(this.$route.params.userId, 10);
            const response = await axios.put(`http://localhost:5000/api/user/${userId}`, this.editedUser);
            /*console.log('Request Configuration:', {
                method: 'PUT',
                url: `http://localhost:5000/api/user/${userId}`,
                data: this.editedUser,
              });*/
              
            console.log('User profile updated:', response.data);
          } catch (error) {
            console.error('Error updating user profile:', error);
          }
        },
      }
};