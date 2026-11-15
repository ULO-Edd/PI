<template>
  <div class="min-h-screen bg-slate-50 flex flex-col md:flex-row">
    <!-- Sidebar on desktop, Top Nav on mobile -->
    <aside class="w-full md:w-64 bg-white border-b md:border-r border-slate-200 flex-shrink-0">
      <div class="p-md flex items-center justify-between">
        <span class="text-lg font-bold text-primary-600">Admin Portal</span>
      </div>
      <nav class="p-sm flex flex-col gap-xs">
        <router-link to="/" class="p-sm rounded-md text-sm text-slate-600 hover:bg-slate-100" active-class="bg-primary-50 text-primary-600 font-semibold">
          Dashboard
        </router-link>
        <router-link to="/users" class="p-sm rounded-md text-sm text-slate-600 hover:bg-slate-100" active-class="bg-primary-50 text-primary-600 font-semibold">
          Users
        </router-link>
        <router-link to="/reports" class="p-sm rounded-md text-sm text-slate-600 hover:bg-slate-100" active-class="bg-primary-50 text-primary-600 font-semibold">
          Reports
        </router-link>
      </nav>
    </aside>

    <!-- Main Content Area -->
    <main class="flex-grow p-md md:p-xl overflow-y-auto">
      <router-view />
    </main>
  </div>
</template>

<script setup lang="ts">
</script>
