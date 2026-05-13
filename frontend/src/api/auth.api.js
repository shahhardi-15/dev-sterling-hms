import api from "./axios";

export const authAPI = {
  login(email, password) {
    return api.post("/auth/login", { email, password });
  },

  changePassword(oldPassword, newPassword) {
    return api.post("/auth/change-password", {
      old_password: oldPassword,
      new_password: newPassword,
    });
  },
};
