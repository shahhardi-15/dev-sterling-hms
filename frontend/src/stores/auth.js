import { defineStore } from "pinia";
import { authAPI } from "../api/auth.api";

export const useAuthStore = defineStore("auth", {
  state: () => ({
    user: JSON.parse(localStorage.getItem("user")) || null,
    token: localStorage.getItem("token") || null,
  }),

  getters: {
    isLoggedIn: (state) => !!state.token,
    userRole: (state) => state.user?.role || null,
    userName: (state) => state.user?.full_name || null,
  },

  actions: {
    async login(email, password) {
      const response = await authAPI.login(email, password);
      const { access_token, user } = response.data.data;

      this.token = access_token;
      this.user = user;

      localStorage.setItem("token", access_token);
      localStorage.setItem("user", JSON.stringify(user));

      return user;
    },

    logout() {
      this.token = null;
      this.user = null;
      localStorage.removeItem("token");
      localStorage.removeItem("user");
    },
  },
});
