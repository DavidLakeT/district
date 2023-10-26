import { createStore } from 'vuex';

const authModule = {
  namespaced: true,
  state: {
    authToken: null,
    userId: null,
    userRole: false
  },
  mutations: {
    setAuthToken(state, token) {
      state.authToken = token;
    },
    setUserId(state, id) {
      state.userId = id;
    },
    setUserRole(state, role) {
      state.userRole = role;
    },
    logout(state) {
        state.authToken = null;
        state.userId = null;
        state.userRole = null;
    },
  },
  getters: {
    isAuthenticated: (state) => state.authToken !== null,
    isAdmin: (state) => state.userRole === true,
    userId: (state) => state.userId,
    authToken: (state) => state.authToken
  },
};

export default createStore({
  modules: {
    auth: authModule,
  },
});
