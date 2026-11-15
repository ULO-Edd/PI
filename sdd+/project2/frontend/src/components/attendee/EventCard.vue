<script setup>
import { defineProps, defineEmits } from 'vue'

const props = defineProps({
  event: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['view'])

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString(undefined, {
    weekday: 'short',
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  })
}
</script>

<template>
  <div class="bg-white rounded-lg shadow overflow-hidden flex flex-col h-full border border-gray-200">
    <div class="p-6 flex-grow">
      <div class="flex justify-between items-start mb-2">
        <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
          {{ event.category_name || 'Event' }}
        </span>
        <span :class="[
          'text-xs font-semibold px-2 py-1 rounded capitalize',
          event.status === 'published' ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800'
        ]">
          {{ event.status }}
        </span>
      </div>
      <h3 class="text-xl font-bold text-gray-900 mb-2 truncate" :title="event.title">
        {{ event.title }}
      </h3>
      <p class="text-sm text-gray-600 mb-4 line-clamp-3 h-15">
        {{ event.description }}
      </p>
      
      <div class="space-y-2 text-sm text-gray-500">
        <div class="flex items-center">
          <svg class="h-4 w-4 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
          </svg>
          {{ formatDate(event.start_time) }}
        </div>
        <div class="flex items-center">
          <svg class="h-4 w-4 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
          </svg>
          {{ event.location }}
        </div>
      </div>
    </div>
    <div class="bg-gray-50 px-6 py-4 mt-auto">
      <button 
        @click="emit('view', event.id)"
        class="w-full bg-blue-600 text-white font-semibold py-2 px-4 rounded hover:bg-blue-700 transition duration-200 focus-visible:ring-2 focus-visible:ring-blue-500 focus-visible:ring-offset-2"
      >
        View Details
      </button>
    </div>
  </div>
</template>
