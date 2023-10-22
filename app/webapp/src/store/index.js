import { createStore } from 'vuex';

const authModule = {
  namespaced: true,
  state: {
    authToken: null,
    user: null,
  },
  mutations: {
    setAuthToken(state, token) {
      state.authToken = token;
    },
    setUser(state, user) {
      state.user = user;
    },
    logout(state) {
        state.authToken = null;
        state.user = null;
    },
  },
  getters: {
    isAuthenticated: (state) => state.authToken !== null,
  },
};

export default createStore({
  modules: {
    auth: authModule,
  },
});
