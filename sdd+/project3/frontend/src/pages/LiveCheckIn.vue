<template>
  <div class="flex flex-col min-h-screen bg-gray-50">
    <!-- Non-blocking connection drop warning banner -->
    <NotificationBanner />

    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8 w-full">
      <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4 mb-8">
        <div>
          <h1 class="text-3xl font-extrabold text-gray-900 tracking-tight">Live Attendance Check-In</h1>
          <p class="text-gray-500 mt-1">Scan, track, and synchronise event attendances in real time.</p>
        </div>

        <!-- Event selector dropdown -->
        <div class="w-full md:max-w-xs">
          <label for="event-select" class="block text-xs font-bold text-gray-700 mb-1">
            Active Event
          </label>
          <select
            id="event-select"
            v-model="activeEventId"
            @change="handleEventChange"
            class="block w-full rounded-md border-gray-300 shadow-sm focus:indigo-500 p-2.5 border text-sm"
          >
            <option value="" disabled>Select an event...</option>
            <option 
              v-for="event in events" 
              :key="event.id" 
              :value="event.id"
            >
              {{ event.title }}
            </option>
          </select>
        </div>
      </div>

      <!-- Main Live Check-In Content Split -->
      <div v-if="activeEventId" class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <!-- Roster and Search List (2/3 width) -->
        <section class="lg:col-span-2 bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden" aria-label="Attendee Roster">
          <div class="px-6 py-4 border-b border-gray-200 flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
            <h2 class="text-lg font-bold text-gray-900">Attendee Roster</h2>
            
            <!-- Roster Search -->
            <div class="w-full sm:max-w-xs">
              <label for="roster-search" class="sr-only">Search Roster</label>
              <input
                id="roster-search"
                type="text"
                v-model="searchQuery"
                placeholder="Search attendee by name..."
                class="block w-full rounded-md border-gray-300 shadow-sm focus:indigo-500 p-2 border text-sm"
              />
            </div>
          </div>

          <div v-if="loadingAttendees" class="p-12 text-center" role="status">
            <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-indigo-600 mx-auto"></div>
            <span class="sr-only">Loading attendee roster...</span>
          </div>

          <div v-else-if="filteredAttendees.length === 0" class="p-12 text-center text-gray-500">
            No attendees match the search criteria.
          </div>

          <ul v-else class="divide-y divide-gray-200">
            <li 
              v-for="attendee in filteredAttendees" 
              :key="attendee.id" 
              class="p-6 flex justify-between items-center"
            >
              <div>
                <p class="text-sm font-bold text-gray-900">{{ attendee.fullName }}</p>
                <p class="text-xs text-gray-500">{{ attendee.email }}</p>
                <div class="flex items-center gap-2 mt-1">
                  <span 
                    class="inline-flex items-center px-2 py-0.5 rounded text-2xs font-semibold"
                    :class="getBadgeClass(attendee)"
                  >
                    {{ getStatusLabel(attendee) }}
                  </span>
                </div>
              </div>

              <!-- Check-In Action Button -->
              <div>
                <button
                  v-if="attendee.checkInStatus === 'Pending'"
                  @click="performCheckIn(attendee)"
                  class="inline-flex items-center px-3 py-1.5 border border-transparent text-xs font-semibold rounded shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none"
                >
                  Check In
                </button>
                <span 
                  v-else 
                  class="inline-flex items-center text-sm font-semibold text-green-600 gap-1"
                >
                  <svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M5 13l4 4L19 7" />
                  </svg>
                  Checked In
                </span>
              </div>
            </li>
          </ul>
        </section>

        <!-- Sidebar Sync Stats (1/3 width) -->
        <section class="space-y-6" aria-label="Sync status panel">
          <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
            <h2 class="text-lg font-bold text-gray-900 mb-4">Resilience Console</h2>
            
            <div class="space-y-4">
              <!-- Connection State indicator -->
              <div class="flex justify-between items-center border-b border-gray-100 pb-3">
                <span class="text-sm font-medium text-gray-500">Connection Status</span>
                <span 
                  class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-semibold"
                  :class="checkInStore.isConnected ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'"
                >
                  {{ checkInStore.isConnected ? 'Online' : 'Offline' }}
                </span>
              </div>

              <!-- Queued offline scans count -->
              <div class="flex justify-between items-center border-b border-gray-100 pb-3">
                <span class="text-sm font-medium text-gray-500">Local Offline Queue</span>
                <span class="text-sm font-bold text-gray-900">
                  {{ checkInStore.pendingQueue.length }} pending
                </span>
              </div>

              <div class="text-xs text-gray-500 mt-2">
                All attendee scans triggered in offline mode are preserved locally in browser persistent stores and synchronised immediately upon connection recovery.
              </div>
            </div>
          </div>
        </section>
      </div>

      <div v-else class="text-center py-20 bg-white rounded-lg border border-gray-200 shadow-sm">
        <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 15l-2 5L9 9l11 4-5 2zm0 0l5 5M7.188 2.239l.777 2.897M5.136 7.965l-2.898-.777M13.95 4.05l-2.122 2.122m-5.657 5.656l-2.12 2.122" />
        </svg>
        <h3 class="mt-2 text-sm font-medium text-gray-900">No event selected</h3>
        <p class="mt-1 text-sm text-gray-500">Please select an active event in the top right to open the check-in roster.</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, computed } from "vue";
import { api } from "../services/api";
import { wsService } from "../services/websocket";
import { useCheckInStore } from "../stores/checkIn";
import NotificationBanner from "../components/NotificationBanner.vue";

export interface Event {
  id: string;
  title: string;
}

export interface AttendeeRoster {
  id: string;
  fullName: string;
  email: string;
  checkInStatus: "Pending" | "CheckedIn";
  isOfflineQueued?: boolean;
}

const events = ref<Event[]>([]);
const attendees = ref<AttendeeRoster[]>([]);
const searchQuery = ref("");
const activeEventId = ref("");
const loadingAttendees = ref(false);

const checkInStore = useCheckInStore();

const fetchEvents = async () => {
  try {
    const response = await api.get("/api/organiser/events");
    events.value = response.data;
  } catch (error) {
    console.error("Failed to load events for check-in:", error);
  }
};

const handleEventChange = () => {
  if (!activeEventId.value) return;

  // Establish WebSocket connection and subscription for selected event
  wsService.disconnect();
  wsService.connect(activeEventId.value);
  fetchRoster();
};

const fetchRoster = async () => {
  if (!activeEventId.value) return;
  loadingAttendees.value = true;
  try {
    const response = await api.get(`/api/organiser/events/${activeEventId.value}/attendees`);
    attendees.value = response.data;
  } catch (error) {
    console.error("Failed to fetch attendee roster:", error);
  } finally {
    loadingAttendees.value = false;
  }
};

const performCheckIn = async (attendee: AttendeeRoster) => {
  if (!activeEventId.value) return;

  if (checkInStore.isConnected) {
    try {
      // Online check-in: post directly to registrations API
      await api.post(`/api/registrations/check-in`, {
        eventId: activeEventId.value,
        attendeeId: attendee.id,
      });
      attendee.checkInStatus = "CheckedIn";
    } catch (error) {
      console.error("Online check-in failed, queueing offline:", error);
      queueOffline(attendee);
    }
  } else {
    // Offline mode: queue scan locally
    queueOffline(attendee);
  }
};

const queueOffline = (attendee: AttendeeRoster) => {
  if (!activeEventId.value) return;

  const result = checkInStore.addScanToQueue(activeEventId.value, attendee.id);
  if (result) {
    attendee.checkInStatus = "CheckedIn";
    attendee.isOfflineQueued = true;
  }
};

const getBadgeClass = (attendee: AttendeeRoster) => {
  if (attendee.checkInStatus === "Pending") return "bg-gray-100 text-gray-800";
  if (attendee.isOfflineQueued) return "bg-amber-100 text-amber-800";
  return "bg-green-100 text-green-800";
};

const getStatusLabel = (attendee: AttendeeRoster) => {
  if (attendee.checkInStatus === "Pending") return "Ready";
  if (attendee.isOfflineQueued) return "Offline Queued";
  return "Checked In";
};

const filteredAttendees = computed(() => {
  const query = searchQuery.value.toLowerCase().trim();
  if (!query) return attendees.value;
  return attendees.value.filter((attendee) =>
    attendee.fullName.toLowerCase().includes(query)
  );
});

onMounted(() => {
  fetchEvents();
});

onUnmounted(() => {
  wsService.disconnect();
});
</script>
