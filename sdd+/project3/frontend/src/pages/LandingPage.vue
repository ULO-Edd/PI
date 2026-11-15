<template>
  <div class="bg-gray-50 min-h-screen">
    <!-- Hero Banner section -->
    <section class="bg-gradient-to-r from-indigo-600 to-indigo-800 text-white py-16 px-6 text-center" aria-labelledby="hero-title">
      <div class="max-w-4xl mx-auto">
        <h1 id="hero-title" class="text-4xl md:text-5xl font-extrabold mb-4 tracking-tight">
          Seamless Event Discovery
        </h1>
        <p class="text-lg md:text-xl text-indigo-100 max-w-2xl mx-auto">
          Discover incredible gatherings, conferences, and workshops hosted globally. Join and book your tickets today.
        </p>
      </div>
    </section>

    <!-- Search and filtering section -->
    <section class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 pt-8" aria-label="Filters">
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 flex flex-col md:flex-row gap-4 items-center justify-between">
        <div class="w-full md:max-w-md">
          <label for="search-input" class="block text-sm font-semibold text-gray-700 mb-2">
            Search Events
          </label>
          <div class="relative rounded-md shadow-sm">
            <input
              id="search-input"
              type="text"
              v-model="searchQuery"
              placeholder="Filter by title or description..."
              class="block w-full rounded-md border-gray-300 pr-10 focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm p-3 border"
              aria-describedby="search-description"
            />
            <span id="search-description" class="sr-only">Type to filter the events list below automatically.</span>
            <div class="absolute inset-y-0 right-0 flex items-center pr-3 pointer-events-none">
              <svg class="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </div>
          </div>
        </div>

        <div class="text-sm text-gray-500" role="status" aria-live="polite">
          Showing {{ filteredEvents.length }} of {{ events.length }} events
        </div>
      </div>
    </section>

    <!-- Catalogue Container -->
    <main>
      <EventCatalogue 
        :events="filteredEvents" 
        :loading="loading" 
        :error="error" 
      />
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from "vue";
import { api } from "../services/api";
import EventCatalogue from "../components/EventCatalogue.vue";
import type { Event } from "../components/EventCard.vue";

const searchQuery = ref("");
const events = ref<Event[]>([]);
const loading = ref(false);
const error = ref<string | null>(null);

const fetchEvents = async () => {
  loading.value = true;
  error.value = null;
  try {
    const response = await api.get("/api/events");
    events.value = response.data;
  } catch (err: any) {
    console.error("Failed to load events:", err);
    error.value = "Unable to load events at this moment. Please check back later.";
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  fetchEvents();
});

const filteredEvents = computed(() => {
  const query = searchQuery.value.toLowerCase().trim();
  if (!query) return events.value;
  return events.value.filter(
    (event) =>
      event.title.toLowerCase().includes(query) ||
      (event.description && event.description.toLowerCase().includes(query))
  );
});
</script>
