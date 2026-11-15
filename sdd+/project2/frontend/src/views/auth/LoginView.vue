<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../../store/auth'

const username = ref('')
const password = ref('')
const error = ref('')

const authStore = useAuthStore()
const router = useRouter()

const handleLogin = async () => {
  error.value = ''
  try {
    await authStore.login(username.value, password.value)
    router.push('/')
  } catch (err) {
    error.value = err.message || 'Invalid credentials'
  }
}
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-full space-y-8 bg-white p-8 rounded-lg shadow-md">
      <div>
        <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">
          Sign in to your account
        </h2>
        <p class="mt-2 text-center text-sm text-gray-600">
          Or
          <router-link to="/register" class="font-medium text-blue-600 hover:text-blue-500 focus-visible:ring-2 focus-visible:ring-blue-500 rounded">
            create a new account
          </router-link>
        </p>
      </div>

      <div aria-live="assertive">
        <p v-if="error" class="bg-red-50 text-red-700 p-3 rounded text-sm font-medium border border-red-200">
          {{ error }}
        </p>
      </div>

      <form class="mt-8 space-y-6" @submit.prevent="handleLogin">
        <div class="rounded-md shadow-sm space-y-4">
          <div>
            <label for="username" class="block text-sm font-medium text-gray-700">Username</label>
            <input 
              id="username" 
              name="username" 
              type="text" 
              required 
              v-model="username"
              class="appearance-none rounded relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm mt-1 focus-visible:ring-2 focus-visible:ring-blue-500" 
              placeholder="Your username"
            />
          </div>

          <div>
            <label for="password" class="block text-sm font-medium text-gray-700">Password</label>
            <input 
              id="password" 
              name="password" 
              type="password" 
              required 
              v-model="password"
              class="appearance-none rounded relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm mt-1 focus-visible:ring-2 focus-visible:ring-blue-500" 
              placeholder="Your password"
            />
          </div>
        </div>

        <div>
          <button 
            type="submit" 
            class="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 focus-visible:ring-2 focus-visible:ring-blue-500"
          >
            Sign in
          </button>
        </div>
      </form>
    </div>
  </div>
</template>
