<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '../../store/auth'
import AccessibleTable from '../../components/organizer/AccessibleTable.vue'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()

const eventId = route.params.id
const attendees = ref([])
const loading = ref(true)
const error = ref('')

const fetchAttendees = async () => {
  loading.value = true
  try {
    const res = await fetch(`http://localhost:8080/events/${eventId}/attendees`, {
      headers: { 'Authorization': `Bearer ${authStore.token}` }
    })
    if (!res.ok) throw new Error('Failed to fetch attendees')
    attendees.value = await res.json()
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
}

onMounted(fetchAttendees)

const updatePresence = async (registrationId, status) => {
  try {
    const res = await fetch(`http://localhost:8080/events/${eventId}/attendees/${registrationId}/presence`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${authStore.token}`
      },
      body: JSON.stringify({ status })
    })
    if (!res.ok) {
      const errMsg = await res.text()
      throw new Error(errMsg || 'Failed to update presence')
    }
    // Update local attendee status
    const index = attendees.value.findIndex(a => a.registration_id === registrationId)
    if (index !== -1) {
      attendees.value[index].status = status
    }
  } catch (err) {
    error.value = err.message
  }
}

const exportCsv = async () => {
  try {
    const res = await fetch(`http://localhost:8080/events/${eventId}/attendees/export`, {
      headers: { 'Authorization': `Bearer ${authStore.token}` }
    })
    if (!res.ok) {
      const errMsg = await res.text()
      throw new Error(errMsg || 'Failed to export CSV')
    }
    const blob = await res.blob()
    const url = window.URL.createObjectURL(blob)
    const a = document.createElement('a')
    a.href = url
    a.download = `attendees_${eventId}.csv`
    document.body.appendChild(a)
    a.click()
    a.remove()
    window.URL.revokeObjectURL(url)
  } catch (err) {
    error.value = err.message
  }
}

const tableHeaders = [
  { label: 'Attendee Name', key: 'attendee_name', class: 'w-1/3', bodyClass: 'font-medium text-gray-900' },
  { label: 'Status', key: 'status', class: 'w-1/3', bodyClass: 'text-gray-500' },
]

const formatAttendeesForTable = computed(() => {
  return attendees.value.map(att => ({
    ...att,
    status: att.status.charAt(0).toUpperCase() + att.status.slice(1).replace('-', ' ') // Capitalize and replace dash
  }))
})
</script>

<template>
  <div class="min-h-screen bg-gray-50">
    <nav class="bg-white shadow-sm h-16 flex items-center mb-8">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 w-full flex justify-between">
        <button @click="router.back()" class="text-blue-600 font-medium flex items-center focus-visible:ring-2 focus-visible:ring-blue-500 rounded px-2">
          <svg class="h-5 w-5 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
          </svg>
          Back to Dashboard
        </button>
      </div>
    </nav>

    <main class="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
      <header class="mb-8">
        <h2 class="text-3xl font-extrabold text-gray-900">Attendees for Event {{ eventId }}</h2>
        <p class="mt-2 text-gray-600">Manage attendee presence and export data.</p>
      </header>

      <div v-if="loading" class="flex justify-center py-12">
        <div class="animate-spin rounded-full h-10 w-10 border-b-2 border-blue-600"></div>
      </div>

      <div v-else-if="error" class="bg-red-50 text-red-700 p-4 rounded-lg border border-red-200 mb-6">
        {{ error }}
      </div>

      <div v-else-if="attendees.length === 0" class="bg-white p-12 text-center rounded-xl shadow-sm border border-gray-200">
        <p class="text-gray-500 mt-1">No attendees registered for this event yet.</p>
      </div>

      <div v-else class="bg-white shadow rounded-lg p-6">
        <div class="flex justify-end mb-4">
          <button @click="exportCsv" class="bg-green-600 hover:bg-green-700 text-white font-medium py-2 px-4 rounded-md focus-visible:ring-2 focus-visible:ring-green-500">
            Export to CSV
          </button>
        </div>
        <AccessibleTable :headers="tableHeaders" :data="formatAttendeesForTable" caption="Attendee Roster">
          <template #actions="{ row }">
            <div class="flex justify-end space-x-2">
              <button 
                @click="updatePresence(row.registration_id, 'Present')"
                :disabled="row.status === 'Present'"
                class="text-blue-600 hover:text-blue-900 focus-visible:ring-2 focus-visible:ring-blue-500 rounded"
              >
                Mark Present
              </button>
              <button 
                @click="updatePresence(row.registration_id, 'No-Show')"
                :disabled="row.status === 'No-Show'"
                class="text-red-600 hover:text-red-900 focus-visible:ring-2 focus-visible:ring-red-500 rounded"
              >
                Mark No-Show
              </button>
            </div>
          </template>
        </AccessibleTable>
      </div>
    </main>
  </div>
</template>
