<script setup>
import { ref, onMounted } from 'vue'
import { useAuthStore } from '../../store/auth'
import AccessibleTable from '../../components/organizer/AccessibleTable.vue'

const authStore = useAuthStore()
const organizers = ref([])
const loading = ref(true)
const error = ref('')

const fetchOrganizers = async () => {
  loading.value = true
  try {
    const res = await fetch('http://localhost:8080/admin/organizers', {
      headers: { 'Authorization': `Bearer ${authStore.token}` }
    })
    if (!res.ok) throw new Error('Failed to fetch organizers')
    organizers.value = await res.json()
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
}

onMounted(fetchOrganizers)

const updateStatus = async (organizerId, status) => {
  try {
    const res = await fetch(`http://localhost:8080/admin/organizers/${organizerId}/status`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${authStore.token}`
      },
      body: JSON.stringify({ status })
    })
    if (!res.ok) {
      const errMsg = await res.text()
      throw new Error(errMsg || 'Failed to update status')
    }
    const index = organizers.value.findIndex(o => o.id === organizerId)
    if (index !== -1) {
      organizers.value[index].status = status
    }
  } catch (err) {
    error.value = err.message
  }
}

const tableHeaders = [
  { label: 'Username', key: 'username', class: 'w-1/3' },
  { label: 'Status', key: 'status', class: 'w-1/3' },
]
</script>

<template>
  <div class="p-6 bg-white shadow rounded-lg">
    <h3 class="text-xl font-bold text-gray-900 mb-4">Manage Organizers</h3>
    
    <div v-if="loading" class="text-center py-8">
      <div class="animate-spin rounded-full h-10 w-10 border-b-2 border-blue-600 mx-auto"></div>
    </div>

    <div v-else-if="error" class="bg-red-50 text-red-700 p-4 rounded-lg border border-red-200 mb-6">
      {{ error }}
    </div>

    <AccessibleTable v-else :headers="tableHeaders" :data="organizers" caption="Organizer Accounts">
      <template #actions="{ row }">
        <div class="flex justify-end space-x-2">
          <button 
            @click="updateStatus(row.id, 'Approved')"
            :disabled="row.status === 'Approved'"
            class="text-blue-600 hover:text-blue-900 focus-visible:ring-2 focus-visible:ring-blue-500 rounded"
          >
            Approve
          </button>
          <button 
            @click="updateStatus(row.id, 'Suspended')"
            :disabled="row.status === 'Suspended'"
            class="text-red-600 hover:text-red-900 focus-visible:ring-2 focus-visible:ring-red-500 rounded"
          >
            Suspend
          </button>
        </div>
      </template>
    </AccessibleTable>
  </div>
</template>
