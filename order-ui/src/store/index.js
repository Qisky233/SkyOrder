import { createStore } from 'vuex';

export default createStore({
  state: {
    token: null,
  },
  mutations: {
    setToken(state, token) {
      state.token = token;
      localStorage.setItem('token', token); // 可选：将 token 存储到 localStorage
    },
  },
});