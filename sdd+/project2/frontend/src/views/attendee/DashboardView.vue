<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../../store/auth'

const registrations = ref([])
const loading = ref(true)
const error = ref('')

const authStore = useAuthStore()
const router = useRouter()

const fetchRegistrations = async () => {
  loading.value = true
  try {
    // In a real implementation, this would call a specific attendee dashboard endpoint.
    // For now, we simulate fetching the user's registrations from the events list or a mock endpoint.
    const res = await fetch(`http://localhost:8080/events`, {
        headers: { 'Authorization': `Bearer ${authStore.token}` }
    })
    if (!res.ok) throw new Error('Failed to fetch registrations')
    const allEvents = await res.json()
    // Mock filtering logic for the attendee's own registrations
    registrations.value = allEvents.slice(0, 3) 
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
}

onMounted(fetchRegistrations)

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString(undefined, {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  })
}
</script>

<template>
  <div class="min-h-screen bg-gray-50">
    <nav class="bg-white shadow-sm h-16 flex items-center mb-8">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 w-full flex justify-between items-center">
        <h1 class="text-xl font-bold text-blue-600">Event Portal</h1>
        <router-link to="/" class="text-sm font-medium text-gray-600 hover:text-blue-600 px-3 py-2 rounded">
           &larr; Back to Discovery
        </router-link>
      </div>
    </nav>

    <main class="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
      <header class="mb-8">
        <h2 class="text-3xl font-extrabold text-gray-900">My Registrations</h2>
        <p class="mt-2 text-gray-600">Track your upcoming and past events.</p>
      </header>

      <div v-if="loading" class="flex justify-center py-12">
        <div class="animate-spin rounded-full h-10 w-10 border-b-2 border-blue-600"></div>
      </div>

      <div v-else-if="error" class="bg-red-50 text-red-700 p-4 rounded-lg border border-red-200 mb-6">
        {{ error }}
      </div>

      <div v-else-if="registrations.length === 0" class="bg-white p-12 text-center rounded-xl shadow-sm border border-gray-200">
        <div class="text-gray-400 mb-4">
           <svg class="mx-auto h-12 w-12" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253" />
           </svg>
        </div>
        <h3 class="text-lg font-bold text-gray-900">No registrations found</h3>
        <p class="text-gray-500 mt-1">You haven't registered for any events yet.</p>
        <router-link to="/" class="mt-6 inline-block bg-blue-600 text-white px-6 py-2 rounded-lg font-medium hover:bg-blue-700 transition">
          Browse Events
        </router-link>
      </div>

      <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div v-for="reg in registrations" :key="reg.id" class="bg-white p-6 rounded-xl shadow-sm border border-gray-200 flex flex-col">
          <div class="flex justify-between items-start mb-4">
             <h4 class="text-lg font-bold text-gray-900 leading-tight line-clamp-2">{{ reg.title }}</h4>
             <span class="px-2 py-1 bg-green-100 text-green-800 text-xs font-bold rounded uppercase">Registered</span>
          </div>
          
          <div class="space-y-2 mb-6 flex-grow">
            <div class="flex items-center text-sm text-gray-500">
               <svg class="h-4 w-4 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
               </svg>
               {{ formatDate(reg.start_time) }}
            </div>
            <div class="flex items-center text-sm text-gray-500">
               <svg class="h-4 w-4 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
               </svg>
               {{ reg.location }}
            </div>
          </div>

          <button 
            @click="router.push(`/events/${reg.id}`)"
            class="w-full text-center py-2 px-4 border border-blue-600 text-blue-600 font-medium rounded-lg hover:bg-blue-50 transition focus-visible:ring-2 focus-visible:ring-blue-500"
          >
            Manage Registration
          </button>
        </div>
      </div>
    </main>
  </div>
</template>
