<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-full space-y-8 bg-white p-8 rounded-lg shadow-md text-center">
      <h2 class="text-3xl font-extrabold text-red-600">
        403 - Forbidden
      </h2>
      <p class="text-gray-600 mt-2">
        You do not have permission to access this page.
      </p>
      <div class="mt-6">
        <router-link to="/" class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded font-medium focus-visible:ring-2 focus-visible:ring-blue-500">
          Back to Dashboard
        </router-link>
      </div>
    </div>
  </div>
</template>
