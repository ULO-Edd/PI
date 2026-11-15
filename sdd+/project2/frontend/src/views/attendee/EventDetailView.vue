<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '../../store/auth'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()

const event = ref(null)
const loading = ref(true)
const error = ref('')
const registering = ref(false)
const successMsg = ref('')

const fetchEvent = async () => {
  loading.value = true
  try {
    const res = await fetch(`http://localhost:8080/events/${route.params.id}`)
    if (!res.ok) throw new Error('Event not found')
    event.value = await res.json()
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
}

const handleRegister = async () => {
  registering.value = true
  error.value = ''
  successMsg.value = ''
  
  try {
    const res = await fetch(`http://localhost:8080/events/${event.value.id}/register`, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${authStore.token}`,
        'Content-Type': 'application/json'
      }
    })
    
    if (!res.ok) {
      const msg = await res.text()
      throw new Error(msg || 'Failed to register')
    }
    
    successMsg.value = 'Successfully registered for this event!'
    // Update local state if needed
  } catch (err) {
    error.value = err.message
  } finally {
    registering.value = false
  }
}

onMounted(fetchEvent)

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleString(undefined, {
    weekday: 'long',
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}
</script>

<template>
  <div class="min-h-screen bg-gray-50 pb-12">
    <nav class="bg-white shadow-sm h-16 flex items-center mb-8">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 w-full flex justify-between">
        <button @click="router.back()" class="text-blue-600 font-medium flex items-center focus-visible:ring-2 focus-visible:ring-blue-500 rounded px-2">
          <svg class="h-5 w-5 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
          </svg>
          Back to Discovery
        </button>
      </div>
    </nav>

    <main class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
      <div v-if="loading" class="text-center py-20">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600 mx-auto"></div>
      </div>

      <div v-else-if="error && !event" class="bg-white p-8 rounded-lg shadow text-center">
        <h2 class="text-2xl font-bold text-red-600 mb-4">Error</h2>
        <p class="text-gray-600">{{ error }}</p>
        <router-link to="/" class="mt-6 inline-block bg-blue-600 text-white px-6 py-2 rounded font-medium">Return Home</router-link>
      </div>

      <article v-else-if="event" class="bg-white rounded-xl shadow-lg overflow-hidden border border-gray-100">
        <div class="bg-blue-600 h-32 w-full"></div>
        
        <div class="p-8 -mt-16">
          <div class="flex flex-col md:flex-row md:justify-between items-start md:items-end mb-6 gap-4">
            <div class="bg-white p-2 rounded-lg shadow-md inline-block">
               <span class="px-4 py-1 rounded bg-blue-100 text-blue-800 text-sm font-bold uppercase tracking-wider">
                 {{ event.status }}
               </span>
            </div>
          </div>

          <h1 class="text-4xl font-extrabold text-gray-900 mb-4">{{ event.title }}</h1>
          
          <div class="grid grid-cols-1 md:grid-cols-3 gap-8 mb-8">
            <div class="md:col-span-2 space-y-6">
              <section>
                <h2 class="text-lg font-bold text-gray-900 border-b pb-2 mb-4">Description</h2>
                <p class="text-gray-700 leading-relaxed whitespace-pre-line">{{ event.description }}</p>
              </section>

              <section class="bg-gray-50 p-6 rounded-lg border border-gray-200">
                <h2 class="text-sm font-bold text-gray-500 uppercase tracking-widest mb-4">Event Details</h2>
                <dl class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                  <div>
                    <dt class="text-xs font-medium text-gray-500 uppercase">Starts</dt>
                    <dd class="text-sm font-semibold text-gray-900">{{ formatDate(event.start_time) }}</dd>
                  </div>
                  <div>
                    <dt class="text-xs font-medium text-gray-500 uppercase">Ends</dt>
                    <dd class="text-sm font-semibold text-gray-900">{{ formatDate(event.end_time) }}</dd>
                  </div>
                  <div>
                    <dt class="text-xs font-medium text-gray-500 uppercase">Location</dt>
                    <dd class="text-sm font-semibold text-gray-900">{{ event.location }}</dd>
                  </div>
                  <div>
                    <dt class="text-xs font-medium text-gray-500 uppercase">Contact</dt>
                    <dd class="text-sm font-semibold text-gray-900">{{ event.contact_email || 'N/A' }}</dd>
                  </div>
                </dl>
              </section>
            </div>

            <aside class="space-y-6">
              <div class="bg-white border-2 border-blue-50 rounded-xl p-6 shadow-sm sticky top-24">
                <div class="mb-4">
                  <span class="text-3xl font-bold text-gray-900">{{ event.capacity }}</span>
                  <span class="text-gray-500 ml-2">Total Capacity</span>
                </div>

                <div v-aria-live="polite">
                  <p v-if="successMsg" class="bg-green-50 text-green-700 p-4 rounded-lg text-sm mb-4 border border-green-200 font-medium">
                    {{ successMsg }}
                  </p>
                  <p v-if="error" class="bg-red-50 text-red-700 p-4 rounded-lg text-sm mb-4 border border-red-200 font-medium">
                    {{ error }}
                  </p>
                </div>

                <button 
                  @click="handleRegister"
                  :disabled="registering || event.status !== 'published'"
                  class="w-full py-4 px-6 rounded-lg font-bold text-white transition-all duration-200 focus-visible:ring-4 focus-visible:ring-blue-300 disabled:opacity-50 disabled:cursor-not-allowed"
                  :class="event.status === 'published' ? 'bg-blue-600 hover:bg-blue-700 shadow-blue-200 shadow-lg' : 'bg-gray-400'"
                >
                  <span v-if="registering" class="flex items-center justify-center">
                    <svg class="animate-spin h-5 w-5 mr-3 text-white" viewBox="0 0 24 24">
                       <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                       <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                    </svg>
                    Processing...
                  </span>
                  <span v-else>Register Now</span>
                </button>
                
                <p class="text-center text-xs text-gray-400 mt-4">
                  By registering, you agree to receive event notifications via email.
                </p>
              </div>
            </aside>
          </div>
        </div>
      </article>
    </main>
  </div>
</template>
