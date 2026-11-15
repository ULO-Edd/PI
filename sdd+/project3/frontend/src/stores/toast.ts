import { defineStore } from "pinia";
import { ref } from "vue";

export const useToastStore = defineStore("toast", () => {
  const toasts = ref([]);

  function addToast(severity: "success" | "info" | "warn" | "error", summary: string, detail: string) {
    const id = Date.now();
    toasts.value.push({ id, severity, summary, detail });
    
    // Auto-remove toast after 5 seconds
    setTimeout(() => {
      removeToast(id);
    }, 5000);
  }

  function removeToast(id: number) {
    toasts.value = toasts.value.filter(t => t.id !== id);
  }

  return {
    toasts,
    addToast,
    removeToast
  };
});

export const createToast = () => useToastStore();
