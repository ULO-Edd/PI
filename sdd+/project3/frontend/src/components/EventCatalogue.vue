<template>
  <section 
    class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8" 
    aria-label="Event Catalogue"
  >
    <div v-if="loading" class="flex justify-center items-center py-12" role="status">
      <div class="animate-spin rounded-full h-10 w-10 border-b-2 border-indigo-600"></div>
      <span class="sr-only">Loading events...</span>
    </div>

    <div v-else-if="error" class="bg-red-50 border border-red-200 rounded-lg p-4 text-center text-red-800" role="alert">
      {{ error }}
    </div>

    <div v-else-if="events.length === 0" class="text-center py-16 bg-white border border-gray-100 rounded-lg shadow-sm" role="status">
      <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
      </svg>
      <h3 class="mt-2 text-sm font-medium text-gray-900">No events found</h3>
      <p class="mt-1 text-sm text-gray-500">Try adjusting your search filters.</p>
    </div>

    <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <EventCard 
        v-for="event in events" 
        :key="event.id" 
        :event="event" 
      />
    </div>
  </section>
</template>

<script setup lang="ts">
import EventCard from "./EventCard.vue";
import type { Event } from "./EventCard.vue";

defineProps<{
  events: Event[];
  loading: boolean;
  error: string | null;
}>();
</script>
