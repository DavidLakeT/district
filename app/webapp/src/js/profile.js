import axios from 'axios';

export default {
    data() {
        return {
          user: {}, // store user details here
          editing: false,
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
                  this.user = response.data;
              } catch (error) {
                  console.error('Error fetching user profile:', error);
              }
          },
  
          async updateUserProfile(userId) {
              try {
                  const response = await axios.put(`http://localhost:5000/api/user/${userId}`, this.updatedUserInfo);
                  console.log('User profile updated:', response.data);
              } catch (error) {
                  console.error('Error updating user profile:', error);
              }
          },
          startEditing() {
              this.editing = true;
          },
      },
};