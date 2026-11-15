<template>
  <article 
    class="bg-white border border-gray-200 rounded-lg shadow-sm overflow-hidden flex flex-col hover:shadow-md transition-shadow focus-within:ring-2 focus-within:ring-indigo-500 focus-within:ring-offset-2 p-5"
    aria-labelledby="event-title"
  >
    <div class="flex-grow">
      <div class="flex justify-between items-start gap-2 mb-2">
        <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-indigo-100 text-indigo-800">
          Available: {{ event.availableTickets }} / {{ event.capacity }}
        </span>
        <span 
          v-if="event.availableTickets === 0" 
          class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-red-100 text-red-800"
          role="status"
        >
          Sold Out
        </span>
      </div>

      <h3 id="event-title" class="text-lg font-bold text-gray-900 mb-2 focus:outline-none">
        {{ event.title }}
      </h3>

      <p class="text-sm text-gray-600 mb-4 line-clamp-3">
        {{ event.description || 'No description provided.' }}
      </p>
    </div>

    <footer class="border-t border-gray-100 pt-4 mt-auto">
      <div class="flex flex-col gap-1.5 text-xs text-gray-500">
        <div class="flex items-center gap-2">
          <svg class="h-4 w-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
          </svg>
          <time :datetime="event.dateTime">
            {{ formatDateTime(event.dateTime) }}
          </time>
        </div>
        <div class="flex items-center gap-2">
          <svg class="h-4 w-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
          </svg>
          <span>{{ event.location || 'Location TBA' }}</span>
        </div>
      </div>
    </footer>
  </article>
</template>

<script setup lang="ts">
export interface Event {
  id: string;
  title: string;
  description: string;
  dateTime: string;
  location: string;
  capacity: number;
  availableTickets: number;
}

defineProps<{
  event: Event;
}>();

const formatDateTime = (isoString: string): string => {
  try {
    const date = new Date(isoString);
    return date.toLocaleDateString(undefined, {
      weekday: "short",
      year: "numeric",
      month: "short",
      day: "numeric",
      hour: "2-digit",
      minute: "2-digit",
    });
  } catch {
    return isoString;
  }
};
</script>
