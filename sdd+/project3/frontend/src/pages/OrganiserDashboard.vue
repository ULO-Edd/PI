<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4 mb-8">
      <div>
        <h1 class="text-3xl font-extrabold text-gray-900 tracking-tight">Organiser Administration</h1>
        <p class="text-gray-500 mt-1">Create, update, and manage your events and registrations.</p>
      </div>
      <button
        v-if="!showForm"
        @click="openCreateForm"
        class="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
      >
        Create Event
      </button>
    </div>

    <!-- Active Form View (Create/Edit) -->
    <div v-if="showForm" class="mb-8">
      <OrganiserEventForm
        :event="selectedEvent"
        :is-edit="isEditMode"
        @save="handleSave"
        @cancel="closeForm"
      />
    </div>

    <!-- Dashboard Main Layout Split -->
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
      <!-- Event List (2/3 width) -->
      <section class="lg:col-span-2 space-y-6" aria-label="Owned Events">
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
          <div class="px-6 py-4 border-b border-gray-200 flex justify-between items-center">
            <h2 class="text-lg font-bold text-gray-900">Your Events</h2>
            <span class="text-sm text-gray-500" role="status">{{ events.length }} total</span>
          </div>

          <div v-if="loadingEvents" class="p-8 text-center" role="status">
            <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-indigo-600 mx-auto"></div>
            <span class="sr-only">Loading events...</span>
          </div>

          <div v-else-if="events.length === 0" class="p-12 text-center">
            <p class="text-gray-500">You haven't created any events yet.</p>
          </div>

          <ul v-else class="divide-y divide-gray-200">
            <li 
              v-for="event in events" 
              :key="event.id" 
              class="p-6 hover:bg-gray-50 transition-colors flex justify-between items-center"
              :class="{ 'bg-indigo-50/50 border-l-4 border-indigo-500': activeEventId === event.id }"
            >
              <div>
                <h3 class="text-base font-bold text-gray-900">{{ event.title }}</h3>
                <p class="text-xs text-gray-500 mt-1">
                  {{ event.location || 'Location TBA' }} · Capacity: {{ event.capacity }} · Tickets Available: {{ event.availableTickets }}
                </p>
              </div>

              <div class="flex items-center gap-2">
                <button
                  @click="viewAttendees(event.id)"
                  class="text-xs font-semibold text-indigo-600 hover:text-indigo-900 px-2.5 py-1.5 rounded bg-indigo-50 hover:bg-indigo-100"
                >
                  Attendees
                </button>
                <button
                  @click="openEditForm(event)"
                  class="text-xs font-semibold text-gray-600 hover:text-gray-900 px-2.5 py-1.5 rounded bg-gray-50 hover:bg-gray-100"
                >
                  Edit
                </button>
                <button
                  @click="deleteEvent(event.id)"
                  class="text-xs font-semibold text-red-600 hover:text-red-900 px-2.5 py-1.5 rounded bg-red-50 hover:bg-red-100"
                >
                  Delete
                </button>
              </div>
            </li>
          </ul>
        </div>
      </section>

      <!-- Attendees list (1/3 width) -->
      <section class="space-y-6" aria-label="Registration details">
        <div class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
          <div class="px-6 py-4 border-b border-gray-200">
            <h2 class="text-lg font-bold text-gray-900">Registered Attendees</h2>
          </div>

          <div v-if="!activeEventId" class="p-8 text-center text-gray-500">
            Select an event to view registrations.
          </div>

          <div v-else-if="loadingAttendees" class="p-8 text-center" role="status">
            <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-indigo-600 mx-auto"></div>
            <span class="sr-only">Loading attendees...</span>
          </div>

          <div v-else-if="attendees.length === 0" class="p-8 text-center text-gray-500">
            No attendees registered for this event yet.
          </div>

          <ul v-else class="divide-y divide-gray-200 max-h-[450px] overflow-y-auto">
            <li 
              v-for="attendee in attendees" 
              :key="attendee.id" 
              class="p-4 hover:bg-gray-50 flex justify-between items-center"
            >
              <div class="min-w-0 flex-1 pr-2">
                <p class="text-sm font-bold text-gray-900 truncate">{{ attendee.fullName }}</p>
                <p class="text-xs text-gray-500 truncate">{{ attendee.email }}</p>
                <span 
                  class="inline-flex items-center px-2 py-0.5 rounded text-2xs font-semibold mt-1"
                  :class="attendee.checkInStatus === 'CheckedIn' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800'"
                >
                  {{ attendee.checkInStatus }}
                </span>
              </div>

              <button
                @click="removeAttendee(attendee.id)"
                class="text-xs font-semibold text-red-600 hover:text-red-900 bg-red-50 hover:bg-red-100 px-2 py-1 rounded"
              >
                Remove
              </button>
            </li>
          </ul>
        </div>
      </section>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { api } from "../services/api";
import OrganiserEventForm from "../components/OrganiserEventForm.vue";
import type { Event } from "../components/EventCard.vue";

export interface Attendee {
  id: string;
  fullName: string;
  email: string;
  checkInStatus: "Pending" | "CheckedIn";
}

const events = ref<Event[]>([]);
const attendees = ref<Attendee[]>([]);
const activeEventId = ref<string | null>(null);

const loadingEvents = ref(false);
const loadingAttendees = ref(false);

const showForm = ref(false);
const isEditMode = ref(false);
const selectedEvent = ref<Event | null>(null);

const fetchEvents = async () => {
  loadingEvents.value = true;
  try {
    const response = await api.get("/api/organiser/events");
    events.value = response.data;
  } catch (error) {
    console.error("Failed to load organiser events:", error);
  } finally {
    loadingEvents.value = false;
  }
};

const viewAttendees = async (eventId: string) => {
  activeEventId.value = eventId;
  loadingAttendees.value = true;
  try {
    const response = await api.get(`/api/organiser/events/${eventId}/attendees`);
    attendees.value = response.data;
  } catch (error) {
    console.error("Failed to load attendees:", error);
  } finally {
    loadingAttendees.value = false;
  }
};

const deleteEvent = async (eventId: string) => {
  if (!confirm("Are you sure you want to permanently delete this event?")) return;
  try {
    await api.delete(`/api/organiser/events/${eventId}`);
    events.value = events.value.filter((e) => e.id !== eventId);
    if (activeEventId.value === eventId) {
      activeEventId.value = null;
      attendees.value = [];
    }
  } catch (error) {
    console.error("Failed to delete event:", error);
  }
};

const removeAttendee = async (attendeeId: string) => {
  if (!confirm("Are you sure you want to remove this attendee's registration?")) return;
  try {
    await api.delete(`/api/organiser/attendees/${attendeeId}`);
    attendees.value = attendees.value.filter((a) => a.id !== attendeeId);
    // Update local ticket availability counter
    const activeEvent = events.value.find((e) => e.id === activeEventId.value);
    if (activeEvent) {
      activeEvent.availableTickets = Math.min(activeEvent.capacity, activeEvent.availableTickets + 1);
    }
  } catch (error) {
    console.error("Failed to remove attendee:", error);
  }
};

const openCreateForm = () => {
  selectedEvent.value = null;
  isEditMode.value = false;
  showForm.value = true;
};

const openEditForm = (event: Event) => {
  selectedEvent.value = event;
  isEditMode.value = true;
  showForm.value = true;
};

const closeForm = () => {
  showForm.value = false;
  selectedEvent.value = null;
};

const handleSave = async (formData: Omit<Event, "id" | "availableTickets"> & { id?: string; availableTickets?: number }) => {
  try {
    if (isEditMode.value && formData.id) {
      const response = await api.put(`/api/organiser/events/${formData.id}`, formData);
      const index = events.value.findIndex((e) => e.id === formData.id);
      if (index !== -1) {
        events.value[index] = response.data;
      }
    } else {
      const response = await api.post("/api/organiser/events", formData);
      events.value.push(response.data);
    }
    closeForm();
  } catch (error) {
    console.error("Failed to save event:", error);
  }
};

onMounted(() => {
  fetchEvents();
});
</script>
