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
            const response = await axios.get(`http://localhost:5000/api/user/${userId}`,{
              withCredentials: true
            });
            this.editedUser = response.data.user;
          } catch (error) {
            console.error('Error fetching user profile:', error);
          }
        },
        async submit() {
          try {
            const userId = parseInt(this.$route.params.userId, 10);
            await axios.put(`http://localhost:5000/api/user/${userId}`, this.editedUser, {
              withCredentials: true
            });
          } catch (error) {
            console.error('Error updating user profile:', error);
          }
        },
      }
};