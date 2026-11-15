<script setup>
import { defineProps, defineEmits } from 'vue'

const props = defineProps({
  headers: {
    type: Array,
    required: true
  },
  data: {
    type: Array,
    required: true
  },
  caption: {
    type: String,
    default: ''
  }
})

// Emitting custom events for row actions if needed
// const emit = defineEmits(['rowClick', 'action'])
</script>

<template>
  <div class="overflow-x-auto rounded-lg shadow-sm border border-gray-200">
    <table class="min-w-full divide-y divide-gray-200" :aria-label="caption">
      <caption v-if="caption" class="sr-only">{{ caption }}</caption>
      <thead class="bg-gray-50">
        <tr>
          <th v-for="header in headers" :key="header.key" scope="col"
            class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
            :class="header.class">
            {{ header.label }}
          </th>
          <!-- Optional: actions column header -->
          <th v-if="$slots.actions" scope="col" class="relative px-6 py-3"><span class="sr-only">Actions</span></th>
        </tr>
      </thead>
      <tbody class="bg-white divide-y divide-gray-200">
        <tr v-for="(row, rowIndex) in data" :key="rowIndex">
          <td v-for="header in headers" :key="header.key" 
            class="px-6 py-4 whitespace-nowrap text-sm"
            :class="header.bodyClass || 'text-gray-900'">
            {{ row[header.key] }}
          </td>
          <!-- Optional: actions column -->
          <td v-if="$slots.actions" class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
            <slot name="actions" :row="row"></slot>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
