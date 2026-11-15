<template>
  <div class="bg-white border border-slate-200 rounded-lg p-6 shadow-sm max-w-lg mx-auto">
    <h3 class="text-lg font-bold text-slate-800 mb-4">
      {{ isEdit ? 'Edit Event' : 'Create New Event' }}
    </h3>

    <form @submit.prevent="handleSubmit" class="space-y-4">
      <!-- Title input -->
      <div>
        <label for="title" class="block text-sm font-semibold text-slate-700 mb-1">
          Event Title *
        </label>
        <input
          id="title"
          type="text"
          v-model="form.title"
          class="block w-full rounded-md border-slate-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm p-2.5 border"
          required
        />
        <p v-if="errors.title" class="mt-1 text-xs text-danger-500" role="alert">
          {{ errors.title }}
        </p>
      </div>

      <!-- Description input -->
      <div>
        <label for="description" class="block text-sm font-semibold text-slate-700 mb-1">
          Description
        </label>
        <textarea
          id="description"
          v-model="form.description"
          rows="3"
          class="block w-full rounded-md border-slate-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm p-2.5 border"
        ></textarea>
      </div>

      <!-- DateTime input -->
      <div>
        <label for="dateTime" class="block text-sm font-semibold text-slate-700 mb-1">
          Date and Time *
        </label>
        <input
          id="dateTime"
          type="datetime-local"
          v-model="form.dateTime"
          class="block w-full rounded-md border-slate-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm p-2.5 border"
          required
        />
        <p v-if="errors.dateTime" class="mt-1 text-xs text-danger-500" role="alert">
          {{ errors.dateTime }}
        </p>
      </div>

      <!-- Location input -->
      <div>
        <label for="location" class="block text-sm font-semibold text-slate-700 mb-1">
          Location
        </label>
        <input
          id="location"
          type="text"
          v-model="form.location"
          class="block w-full rounded-md border-slate-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm p-2.5 border"
        />
      </div>

      <!-- Capacity input -->
      <div>
        <label for="capacity" class="block text-sm font-semibold text-slate-700 mb-1">
          Capacity *
        </label>
        <input
          id="capacity"
          type="number"
          v-model.number="form.capacity"
          min="1"
          class="block w-full rounded-md border-slate-300 shadow-sm focus:border-primary-500 focus:ring-primary-500 sm:text-sm p-2.5 border"
          required
        />
        <p v-if="errors.capacity" class="mt-1 text-xs text-danger-500" role="alert">
          {{ errors.capacity }}
        </p>
      </div>

      <!-- Action buttons -->
      <div class="flex justify-end gap-3 pt-4 border-t border-slate-100">
        <button
          type="button"
          @click="$emit('cancel')"
          class="px-4 py-2 border border-slate-300 rounded-md text-sm font-medium text-slate-700 bg-white hover:bg-slate-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500"
        >
          Cancel
        </button>
        <button
          type="submit"
          class="px-4 py-2 border border-transparent rounded-md text-sm font-medium text-white bg-primary-600 hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500"
        >
          Save
        </button>
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from "vue";
import type { Event } from "./EventCard.vue";

const props = defineProps<{
  event?: Event | null;
  isEdit?: boolean;
}>();

const emit = defineEmits<{
  (e: "save", data: Omit<Event, "id" | "availableTickets"> & { id?: string; availableTickets?: number }): void;
  (e: "cancel"): void;
}>();

const form = ref({
  title: "",
  description: "",
  dateTime: "",
  location: "",
  capacity: 100,
});

const errors = ref({
  title: "",
  dateTime: "",
  capacity: "",
});

// Watch for initial event payload to populate edit forms
watch(
  () => props.event,
  (newVal) => {
    if (newVal) {
      // Format DateTime from ISO string to datetime-local friendly string (YYYY-MM-DDTHH:MM)
      const date = new Date(newVal.dateTime);
      const formattedDate = date.toISOString().slice(0, 16);

      form.value = {
        title: newVal.title,
        description: newVal.description || "",
        dateTime: formattedDate,
        location: newVal.location || "",
        capacity: newVal.capacity,
      };
    } else {
      form.value = {
        title: "",
        description: "",
        dateTime: "",
        location: "",
        capacity: 100,
      };
    }
  },
  { immediate: true }
);

const validate = (): boolean => {
  let isValid = true;
  errors.value = { title: "", dateTime: "", capacity: "" };

  // Title validation
  if (!form.value.title.trim()) {
    errors.value.title = "Title is required.";
    isValid = false;
  } else if (form.value.title.length > 100) {
    errors.value.title = "Title cannot exceed 100 characters.";
    isValid = false;
  }

  // DateTime validation
  if (!form.value.dateTime) {
    errors.value.dateTime = "Date and Time is required.";
    isValid = false;
  } else {
    const selectedDate = new Date(form.value.dateTime);
    const now = new Date();
    if (selectedDate <= now && !props.isEdit) {
      errors.value.dateTime = "Event Date and Time must be in the future.";
      isValid = false;
    }
  }

  // Capacity validation
  if (form.value.capacity < 1) {
    errors.value.capacity = "Capacity must be at least 1.";
    isValid = false;
  }

  return isValid;
};

const handleSubmit = () => {
  if (!validate()) return;

  // Transform local input date back to standard ISO string
  const payload = {
    ...form.value,
    dateTime: new Date(form.value.dateTime).toISOString(),
    id: props.event?.id,
    availableTickets: props.event?.availableTickets ?? form.value.capacity,
  };

  emit("save", payload);
};
</script>
