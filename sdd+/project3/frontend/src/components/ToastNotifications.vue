<template>
  <div class="fixed top-4 right-4 z-50 flex flex-col gap-sm">
    <transition-group name="list">
      <div 
        v-for="toast in toastStore.toasts" 
        :key="toast.id"
        :class="[
          'p-md rounded-md shadow-hard border-l-4 min-w-[300px] transition-all duration-300',
          toast.severity === 'success' ? 'bg-success-50 border-success-500 text-success-800' : 
          toast.severity === 'info' ? 'bg-info-50 border-info-500 text-info-800' :
          toast.severity === 'warn' ? 'bg-warning-50 border-warning-500 text-warning-800' : 
          'bg-danger-50 border-danger-500 text-danger-800'
        ]"
      >
        <div class="flex justify-between items-start">
          <div>
            <p class="font-bold text-sm">{{ toast.summary }}</p>
            <p class="text-xs opacity-90">{{ toast.detail }}</p>
          </div>
          <button @click="toastStore.removeToast(toast.id)" class="text-slate-400 hover:text-slate-600">
            &times;
          </button>
        </div>
      </div>
    </transition-group>
  </div>
</template>

<script setup lang="ts">
import { useToastStore } from "../stores/toast";
const toastStore = useToastStore();
</script>

<style scoped>
.list-enter-active, .list-leave-active {
  transition: all 0.3s ease;
}
.list-enter-from {
  opacity: 0;
  transform: translateX(30px);
}
.list-leave-to {
  opacity: 0;
  transform: translateX(30px);
}
</style>
