<template>
  <transition name="slide-down">
    <div 
      v-if="!checkInStore.isConnected" 
       class="bg-warning-500 text-white font-medium py-3 px-6 text-center text-sm shadow flex items-center justify-center gap-2"
      role="alert"
      aria-live="assertive"
    >
      <svg class="h-5 w-5 animate-pulse" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
      </svg>
      <span>Offline - Network Connection Interrupted. Scans are being queued locally.</span>
    </div>
  </transition>
</template>

<script setup lang="ts">
import { useCheckInStore } from "../stores/checkIn";

const checkInStore = useCheckInStore();
</script>

<style scoped>
.slide-down-enter-active,
.slide-down-leave-active {
  transition: transform 0.3s ease, opacity 0.3s ease;
}

.slide-down-enter-from,
.slide-down-leave-to {
  transform: translateY(-100%);
  opacity: 0;
}
</style>
