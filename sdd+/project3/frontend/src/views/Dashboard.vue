<template>
  <div>
    <h1 class="text-2xl font-bold text-slate-800 mb-xl">Dashboard Overview</h1>
    
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-md">
      <div v-for="i in 8" :key="i" class="p-lg bg-white rounded-lg border border-slate-200 shadow-soft">
        <h3 class="text-sm font-semibold text-slate-500 uppercase tracking-wider">Statistic {{ i }}</h3>
        <p class="text-3xl font-bold text-primary-600 mt-xs">1,234</p>
        <p class="text-xs text-success-500 mt-xs">↑ 12% from last month</p>
      </div>
    </div>

    <div class="mt-xl bg-white rounded-lg border border-slate-200 shadow-soft">
      <div class="p-lg border-b border-slate-100">
        <h2 class="text-lg font-bold text-slate-800">Recent Activity</h2>
      </div>
      <div class="p-lg text-sm text-slate-600">
        <p>No recent activity to display.</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
</script>
