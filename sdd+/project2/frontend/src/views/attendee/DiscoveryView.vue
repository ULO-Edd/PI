<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../../store/auth'
import EventCard from '../../components/attendee/EventCard.vue'

const events = ref([])
const categories = ref([])
const loading = ref(true)
const error = ref('')

const searchQuery = ref('')
const selectedCategory = ref('')

const authStore = useAuthStore()
const router = useRouter()

const fetchEvents = async () => {
  loading.value = true
  try {
    const res = await fetch('http://localhost:8080/events?status=published')
    if (!res.ok) throw new Error('Failed to fetch events')
    events.value = await res.json()
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
}

const fetchCategories = async () => {
  try {
    const res = await fetch('http://localhost:8080/categories')
    if (!res.ok) throw new Error('Failed to fetch categories')
    categories.value = await res.json()
  } catch (err) {
    console.error(err)
  }
}

onMounted(() => {
  fetchEvents()
  fetchCategories()
})

const filteredEvents = computed(() => {
  if (!events.value) return []
  return events.value.filter(e => {
    const matchesSearch = e.title.toLowerCase().includes(searchQuery.value.toLowerCase()) || 
                         e.description.toLowerCase().includes(searchQuery.value.toLowerCase())
    const matchesCategory = !selectedCategory.value || e.category_id === selectedCategory.value
    return matchesSearch && matchesCategory
  })
})

const handleViewEvent = (id) => {
  router.push(`/events/${id}`)
}

const handleLogout = () => {
  authStore.logout()
  router.push('/login')
}
</script>

<template>
  <div class="min-h-screen bg-gray-50">
    <nav class="bg-white shadow-sm sticky top-0 z-10">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <div class="flex items-center">
            <h1 class="text-xl font-bold text-blue-600">Event Portal</h1>
            <div class="hidden sm:ml-6 sm:flex sm:space-x-8">
              <router-link to="/" class="border-blue-500 text-gray-900 inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium">
                Discovery
              </router-link>
              <router-link to="/dashboard" class="border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700 inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium">
                My Registrations
              </router-link>
            </div>
          </div>
          <div class="flex items-center space-x-4">
            <span class="text-sm text-gray-700 hidden md:inline-block">Welcome, <strong>{{ authStore.user?.username }}</strong></span>
            <button @click="handleLogout" class="text-sm text-red-600 font-medium hover:text-red-500 focus-visible:ring-2 focus-visible:ring-red-500 rounded px-2 py-1">
              Logout
            </button>
          </div>
        </div>
      </div>
    </nav>

    <header class="bg-white border-b">
      <div class="max-w-7xl mx-auto py-8 px-4 sm:px-6 lg:px-8">
        <h2 class="text-3xl font-extrabold text-gray-900 tracking-tight">Discover Events</h2>
        <p class="mt-2 text-lg text-gray-500">Find and register for upcoming events at ESTG / IPLeiria.</p>
        
        <div class="mt-8 flex flex-col md:flex-row space-y-4 md:space-y-0 md:space-x-4">
          <div class="flex-grow">
            <label for="search" class="sr-only">Search events</label>
            <div class="relative">
              <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                <svg class="h-5 w-5 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                </svg>
              </div>
              <input 
                id="search"
                v-model="searchQuery"
                class="block w-full pl-10 pr-3 py-2 border border-gray-300 rounded-md leading-5 bg-white placeholder-gray-500 focus:outline-none focus:placeholder-gray-400 focus:ring-1 focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                placeholder="Search by title or description..."
                type="search"
              >
            </div>
          </div>
          <div class="w-full md:w-64">
            <label for="category" class="sr-only">Filter by category</label>
            <select 
              id="category"
              v-model="selectedCategory"
              class="block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm rounded-md"
            >
              <option value="">All Categories</option>
              <option v-for="cat in categories" :key="cat.id" :value="cat.id">{{ cat.name }}</option>
            </select>
          </div>
        </div>
      </div>
    </header>

    <main class="max-w-7xl mx-auto py-12 px-4 sm:px-6 lg:px-8">
      <div v-if="loading" class="flex justify-center py-12">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600 text-center"></div>
      </div>

      <div v-else-if="error" class="bg-red-50 border border-red-200 text-red-700 p-4 rounded-md text-center">
        {{ error }}
        <button @click="fetchEvents" class="ml-4 underline font-bold">Retry</button>
      </div>

      <div v-else-if="filteredEvents.length === 0" class="text-center py-12 bg-white rounded-lg shadow-sm border border-dashed border-gray-300">
        <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
        <h3 class="mt-2 text-sm font-medium text-gray-900">No events found</h3>
        <p class="mt-1 text-sm text-gray-500">Try adjusting your search or category filter.</p>
      </div>

      <div v-else class="grid grid-cols-1 gap-y-10 sm:grid-cols-2 gap-x-6 lg:grid-cols-3 xl:gap-x-8">
        <EventCard 
          v-for="event in filteredEvents" 
          :key="event.id" 
          :event="event"
          @view="handleViewEvent"
        />
      </div>
    </main>
  </div>
</template>
