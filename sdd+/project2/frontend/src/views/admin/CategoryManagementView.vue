<script setup>
import { ref, onMounted } from 'vue'
import { useAuthStore } from '../../store/auth'
import AccessibleTable from '../../components/organizer/AccessibleTable.vue'

const authStore = useAuthStore()
const categories = ref([])
const loading = ref(true)
const error = ref('')
const newCategoryName = ref('')
const newCategoryDescription = ref('')

const fetchCategories = async () => {
  loading.value = true
  try {
    const res = await fetch('http://localhost:8080/categories')
    if (!res.ok) throw new Error('Failed to fetch categories')
    categories.value = await res.json()
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
}

onMounted(fetchCategories)

const createCategory = async () => {
  try {
    const res = await fetch('http://localhost:8080/categories', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${authStore.token}`
      },
      body: JSON.stringify({ name: newCategoryName.value, description: newCategoryDescription.value })
    })
    if (!res.ok) {
      const errMsg = await res.text()
      throw new Error(errMsg || 'Failed to create category')
    }
    newCategoryName.value = ''
    newCategoryDescription.value = ''
    fetchCategories()
  } catch (err) {
    error.value = err.message
  }
}

const tableHeaders = [
  { label: 'Name', key: 'name', class: 'w-1/3' },
  { label: 'Description', key: 'description', class: 'w-2/3' },
]
</script>

<template>
  <div class="p-6 bg-white shadow rounded-lg">
    <h3 class="text-xl font-bold text-gray-900 mb-4">Manage Categories</h3>

    <form @submit.prevent="createCategory" class="mb-6 flex space-x-4">
      <div class="flex-grow">
        <label for="category-name" class="sr-only">Category Name</label>
        <input type="text" id="category-name" v-model="newCategoryName" placeholder="New category name" required 
               class="block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm focus-visible:ring-2 focus-visible:ring-blue-500">
      </div>
      <div class="flex-grow">
        <label for="category-description" class="sr-only">Category Description</label>
        <input type="text" id="category-description" v-model="newCategoryDescription" placeholder="Description" 
               class="block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm focus-visible:ring-2 focus-visible:ring-blue-500">
      </div>
      <button type="submit" class="bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-md focus-visible:ring-2 focus-visible:ring-blue-500">
        Create
      </button>
    </form>

    <div v-if="loading" class="text-center py-8">
      <div class="animate-spin rounded-full h-10 w-10 border-b-2 border-blue-600 mx-auto"></div>
    </div>

    <div v-else-if="error" class="bg-red-50 text-red-700 p-4 rounded-lg border border-red-200 mb-6">
      {{ error }}
    </div>

    <AccessibleTable v-else :headers="tableHeaders" :data="categories" caption="Event Categories"/>
  </div>
</template>
