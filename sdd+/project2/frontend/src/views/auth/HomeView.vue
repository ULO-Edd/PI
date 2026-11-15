<script setup>
import { useRouter } from 'vue-router'
import { useAuthStore } from '../../store/auth'

const authStore = useAuthStore()
const router = useRouter()

const handleLogout = () => {
  authStore.logout()
  router.push('/login')
}
</script>

<template>
  <div class="min-h-screen bg-gray-100">
    <nav class="bg-white shadow">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <div class="flex items-center">
            <h1 class="text-xl font-bold text-gray-900">Event Portal</h1>
          </div>
          <div class="flex items-center space-x-4">
            <span class="text-gray-700">Hello, <strong>{{ authStore.user?.username }}</strong> ({{ authStore.userRole }})</span>
            <button 
              @click="handleLogout"
              class="bg-red-600 hover:bg-red-700 text-white px-3 py-1.5 rounded text-sm font-medium focus-visible:ring-2 focus-visible:ring-red-500"
            >
              Logout
            </button>
          </div>
        </div>
      </div>
    </nav>

    <main class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
      <!-- Status Box for Pending/Suspended accounts -->
      <div v-if="authStore.userStatus === 'Pending'" class="bg-yellow-50 border border-yellow-200 text-yellow-800 p-4 rounded mb-6">
        <p class="font-medium">Your account is pending administrator approval.</p>
        <p class="text-sm mt-1">You will not be able to publish events until your account is approved.</p>
      </div>

      <div class="bg-white shadow rounded-lg p-6">
        <h2 class="text-2xl font-bold text-gray-900 mb-4">Welcome back!</h2>
        <p class="text-gray-600">This is the dashboard placeholder. In the next steps, we will build Event Discovery and Organizer panels!</p>
      </div>
    </main>
  </div>
</template>
