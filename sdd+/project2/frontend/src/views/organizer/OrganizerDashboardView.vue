<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../../store/auth'
import StatsView from './StatsView.vue'

const router = useRouter()
const authStore = useAuthStore()

const categories = ref([])
const events = ref([])
const loading = ref(true)
const error = ref('')

const eventForm = ref({
  title: '',
  description: '',
  location: '',
  startTime: '',
  endTime: '',
  capacity: 1,
  categoryId: '',
  contactEmail: '',
})

const isEditMode = ref(false)
const currentEventId = ref(null)

const fetchCategories = async () => {
  try {
    const res = await fetch('http://localhost:8080/categories')
    if (!res.ok) throw new Error('Failed to fetch categories')
    categories.value = await res.json()
  } catch (err) {
    console.error(err)
  }
}

const fetchOrganizerEvents = async () => {
  loading.value = true
  try {
    const res = await fetch(`http://localhost:8080/events?organiser_id=${authStore.user.id}`, {
      headers: { 'Authorization': `Bearer ${authStore.token}` }
    })
    if (!res.ok) throw new Error('Failed to fetch organizer events')
    events.value = await res.json()
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchCategories()
  fetchOrganizerEvents()
})

const handleEventSubmit = async () => {
  error.value = ''
  try {
    const method = isEditMode.value ? 'PUT' : 'POST'
    const url = isEditMode.value 
      ? `http://localhost:8080/events/${currentEventId.value}` 
      : 'http://localhost:8080/events'

    const payload = {
      ...eventForm.value,
      startTime: new Date(eventForm.value.startTime).toISOString(),
      endTime: new Date(eventForm.value.endTime).toISOString(),
    }

    const res = await fetch(url, {
      method: method,
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${authStore.token}`,
      },
      body: JSON.stringify(payload),
    })

    if (!res.ok) {
      const errMsg = await res.text()
      throw new Error(errMsg || 'Failed to save event')
    }

    // Reset form and refetch events
    resetForm()
    fetchOrganizerEvents()
  } catch (err) {
    error.value = err.message
  }
}

const resetForm = () => {
  eventForm.value = {
    title: '',
    description: '',
    location: '',
    startTime: '',
    endTime: '',
    capacity: 1,
    categoryId: '',
    contactEmail: '',
  }
  isEditMode.value = false
  currentEventId.value = null
}

const editEvent = (event) => {
  eventForm.value = {
    title: event.title,
    description: event.description,
    location: event.location,
    startTime: new Date(event.start_time).toISOString().slice(0, 16), // format for datetime-local
    endTime: new Date(event.end_time).toISOString().slice(0, 16),
    capacity: event.capacity,
    categoryId: event.category_id,
    contactEmail: event.contact_email,
  }
  isEditMode.value = true
  currentEventId.value = event.id
}

const cancelEvent = async (id) => {
  if (!confirm('Are you sure you want to cancel this event?')) return
  try {
    const res = await fetch(`http://localhost:8080/events/${id}/cancel`, {
      method: 'PUT',
      headers: { 'Authorization': `Bearer ${authStore.token}` }
    })
    if (!res.ok) {
      const errMsg = await res.text()
      throw new Error(errMsg || 'Failed to cancel event')
    }
    fetchOrganizerEvents()
  } catch (err) {
    error.value = err.message
  }
}

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString(undefined, {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  })
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
            <h1 class="text-xl font-bold text-blue-600">Organizer Dashboard</h1>
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
        <h2 class="text-3xl font-extrabold text-gray-900 tracking-tight">Manage Your Events</h2>
        <p class="mt-2 text-lg text-gray-500">Create, edit, and oversee your event portfolio.</p>
      </div>
    </header>

    <main class="max-w-7xl mx-auto py-12 px-4 sm:px-6 lg:px-8">
      <div v-if="authStore.userStatus === 'Pending'" class="bg-yellow-50 border border-yellow-200 text-yellow-800 p-4 rounded mb-6">
        <p class="font-medium">Your organizer account is pending administrator approval.</p>
        <p class="text-sm mt-1">You cannot create or publish events until your account is approved.</p>
      </div>

      <div v-else class="grid grid-cols-1 md:grid-cols-3 gap-8">
        <!-- Event Form -->
        <div class="md:col-span-1 bg-white p-6 rounded-lg shadow-md border border-gray-200">
          <h3 class="text-xl font-bold text-gray-900 mb-6">{{ isEditMode ? 'Edit Event' : 'Create New Event' }}</h3>
          
          <div aria-live="assertive">
            <p v-if="error" class="bg-red-50 text-red-700 p-3 rounded text-sm font-medium border border-red-200 mb-4">
              {{ error }}
            </p>
          </div>

          <form @submit.prevent="handleEventSubmit" class="space-y-4">
            <div>
              <label for="title" class="block text-sm font-medium text-gray-700">Event Title</label>
              <input type="text" id="title" v-model="eventForm.title" required 
                     class="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm focus-visible:ring-2 focus-visible:ring-blue-500">
            </div>
            <div>
              <label for="description" class="block text-sm font-medium text-gray-700">Description</label>
              <textarea id="description" v-model="eventForm.description" rows="4" 
                        class="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm focus-visible:ring-2 focus-visible:ring-blue-500"></textarea>
            </div>
            <div>
              <label for="location" class="block text-sm font-medium text-gray-700">Location</label>
              <input type="text" id="location" v-model="eventForm.location" required 
                     class="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm focus-visible:ring-2 focus-visible:ring-blue-500">
            </div>
            <div class="grid grid-cols-2 gap-4">
              <div>
                <label for="start_time" class="block text-sm font-medium text-gray-700">Start Time</label>
                <input type="datetime-local" id="start_time" v-model="eventForm.startTime" required 
                       class="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm focus-visible:ring-2 focus-visible:ring-blue-500">
              </div>
              <div>
                <label for="end_time" class="block text-sm font-medium text-gray-700">End Time</label>
                <input type="datetime-local" id="end_time" v-model="eventForm.endTime" required 
                       class="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm focus-visible:ring-2 focus-visible:ring-blue-500">
              </div>
            </div>
            <div>
              <label for="capacity" class="block text-sm font-medium text-gray-700">Capacity</label>
              <input type="number" id="capacity" v-model.number="eventForm.capacity" required min="1" 
                     class="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm focus-visible:ring-2 focus-visible:ring-blue-500">
            </div>
            <div>
              <label for="category" class="block text-sm font-medium text-gray-700">Category</label>
              <select id="category" v-model="eventForm.categoryId" required 
                      class="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm focus-visible:ring-2 focus-visible:ring-blue-500">
                <option value="">Select a category</option>
                <option v-for="cat in categories" :key="cat.id" :value="cat.id">{{ cat.name }}</option>
              </select>
            </div>
            <div>
              <label for="contact_email" class="block text-sm font-medium text-gray-700">Contact Email</label>
              <input type="email" id="contact_email" v-model="eventForm.contactEmail" required 
                     class="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm focus-visible:ring-2 focus-visible:ring-blue-500">
            </div>
            <div class="flex space-x-4">
              <button type="submit" 
                      class="flex-1 inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 focus-visible:ring-2 focus-visible:ring-blue-500">
                {{ isEditMode ? 'Update Event' : 'Create Event' }}
              </button>
              <button type="button" @click="resetForm" v-if="isEditMode" 
                      class="flex-1 inline-flex justify-center py-2 px-4 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 focus-visible:ring-2 focus-visible:ring-blue-500">
                Cancel
              </button>
            </div>
          </form>
        </div>

        <!-- Event List -->
        <div class="md:col-span-2 bg-white p-6 rounded-lg shadow-md border border-gray-200">
          <h3 class="text-xl font-bold text-gray-900 mb-6">Your Events</h3>
          
          <div v-if="loading" class="flex justify-center py-8">
            <div class="animate-spin rounded-full h-10 w-10 border-b-2 border-blue-600"></div>
          </div>

          <div v-else-if="events.length === 0" class="text-center py-8 text-gray-500">
            No events created yet. Start by creating one!
          </div>

          <ul v-else class="divide-y divide-gray-200">
            <li v-for="event in events" :key="event.id" class="py-4 flex justify-between items-center">
              <div>
                <h4 class="text-lg font-bold text-gray-900">{{ event.title }}</h4>
                <p class="text-sm text-gray-500">{{ formatDate(event.start_time) }} - {{ event.location }}</p>
                <span :class="[
                  'px-2 py-0.5 mt-1 inline-flex text-xs leading-5 font-semibold rounded-full',
                  event.status === 'published' ? 'bg-green-100 text-green-800' :
                  event.status === 'draft' ? 'bg-gray-100 text-gray-800' :
                  'bg-red-100 text-red-800'
                ]">
                  {{ event.status }}
                </span>
              </div>
              <div class="flex space-x-2">
                <button @click="editEvent(event)" 
                        class="inline-flex items-center px-3 py-1.5 border border-transparent text-sm font-medium rounded-md text-blue-700 bg-blue-100 hover:bg-blue-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 focus-visible:ring-2 focus-visible:ring-blue-500">
                  Edit
                </button>
                <button @click="cancelEvent(event.id)" 
                        class="inline-flex items-center px-3 py-1.5 border border-transparent text-sm font-medium rounded-md text-red-700 bg-red-100 hover:bg-red-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 focus-visible:ring-2 focus-visible:ring-red-500">
                  Cancel
                </button>
                <router-link :to="`/organizer/events/${event.id}/attendees`"
                  class="inline-flex items-center px-3 py-1.5 border border-transparent text-sm font-medium rounded-md text-gray-700 bg-gray-100 hover:bg-gray-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-500 focus-visible:ring-2 focus-visible:ring-gray-500"
                >
                  Attendees
                </router-link>
              </div>
            </li>
          </ul>
        </div>
      </div>
    </main>
  </div>
</template>
