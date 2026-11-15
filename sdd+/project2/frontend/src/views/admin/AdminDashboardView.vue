<script setup>
import { useRouter } from 'vue-router'
import { useAuthStore } from '../../store/auth'

const router = useRouter()
const authStore = useAuthStore()

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
            <h1 class="text-xl font-bold text-blue-600">Admin Dashboard</h1>
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

    <main class="max-w-7xl mx-auto py-12 px-4 sm:px-6 lg:px-8">
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
        <router-link to="/admin/organizers" class="bg-white p-6 rounded-lg shadow-md border border-gray-200 hover:shadow-lg transition-shadow">
          <h3 class="text-xl font-bold text-gray-900 mb-2">Manage Organizers</h3>
          <p class="text-gray-600">Approve or suspend organizer accounts.</p>
        </router-link>
        <router-link to="/admin/categories" class="bg-white p-6 rounded-lg shadow-md border border-gray-200 hover:shadow-lg transition-shadow">
          <h3 class="text-xl font-bold text-gray-900 mb-2">Manage Categories</h3>
          <p class="text-gray-600">Create, edit, and delete event categories.</p>
        </router-link>
        <router-link to="/admin/stats" class="bg-white p-6 rounded-lg shadow-md border border-gray-200 hover:shadow-lg transition-shadow">
          <h3 class="text-xl font-bold text-gray-900 mb-2">Global Statistics</h3>
          <p class="text-gray-600">View platform-wide statistics.</p>
        </router-link>
      </div>
    </main>
  </div>
</template>
