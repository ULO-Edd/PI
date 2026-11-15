import { defineStore } from "pinia";
import { api } from "../services/api";

export interface User {
  id: string;
  email: string;
  role: "Admin" | "Organiser" | "Attendee";
}

export const useAuthStore = defineStore("auth", {
  state: () => ({
    isAuthenticated: false,
    user: null as User | null,
    loading: false,
  }),

  getters: {
    isOrganiser: (state) => {
      return state.user ? ["Admin", "Organiser"].includes(state.user.role) : false;
    },
  },

  actions: {
    // Clear state locally (triggered by global HTTP 401 interceptor or logout)
    clearSessionState() {
      this.isAuthenticated = false;
      this.user = null;
    },

    // Authenticate user with credentials
    async login(email: string, password: string) {
      this.loading = true;
      try {
        const response = await api.post("/api/auth/login", { email, password });
        this.isAuthenticated = true;
        this.user = response.data.user;
        return response.data;
      } catch (error) {
        this.clearSessionState();
        throw error;
      } finally {
        this.loading = false;
      }
    },

    // Invalidate active session and logout
    async logout() {
      this.loading = true;
      try {
        await api.post("/api/auth/logout");
      } catch (error) {
        // Log error but proceed to clear local state regardless
        console.error("Logout failed on server:", error);
      } finally {
        this.clearSessionState();
        this.loading = false;
      }
    },

    // Verify session validity upon application boot
    async checkSession() {
      try {
        const response = await api.get("/api/auth/session");
        this.isAuthenticated = true;
        this.user = response.data.user;
      } catch (error) {
        this.clearSessionState();
      }
    },
  },
});
