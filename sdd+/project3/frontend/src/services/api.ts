import axios from "axios";
import { useToastStore } from "../stores/toast";

// Create configured axios instance targeting the secure API Gateway
export const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || "http://localhost:8080",
  withCredentials: true, // Send secure HTTP-only cookies in cross-origin requests
  headers: {
    "Content-Type": "application/json",
  },
});

// Configure response interceptor to handle session expiration and network errors globally
api.interceptors.response.use(
  (response) => response,
  (error) => {
    // 1. Handle Connection Failures (ERR_CONNECTION_REFUSED / Network Error)
    if (!error.response) {
      const toast = useToastStore();
      toast.addToast(
        "error", 
        "Connection Error", 
        "Backend services are unavailable. Please check if Docker is running."
      );
      return Promise.reject(error);
    }

    // 2. Handle session expiration (HTTP 401)
    if (error.response && error.response.status === 401) {
      // Lazy load auth store to avoid circular dependency
      import("../stores/auth").then(({ useAuthStore }) => {
        const authStore = useAuthStore();
        authStore.clearSessionState();
      });
    }
    return Promise.reject(error);
  }
);

