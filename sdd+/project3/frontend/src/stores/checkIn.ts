import { defineStore } from "pinia";
import { api } from "../services/api";

export interface CheckInScan {
  scanId: string;
  eventId: string;
  attendeeId: string;
  scannedAt: string;
  syncStatus: "Queued" | "Syncing" | "Synced";
  correlationId: string;
}

const LOCAL_STORAGE_KEY = "sdd3_pending_checkins";

export const useCheckInStore = defineStore("checkIn", {
  state: () => {
    // Load initial queue state from client LocalStorage
    let initialQueue: CheckInScan[] = [];
    try {
      const stored = localStorage.getItem(LOCAL_STORAGE_KEY);
      if (stored) {
        initialQueue = JSON.parse(stored);
      }
    } catch (e) {
      console.error("Failed to parse stored check-ins:", e);
    }

    return {
      pendingQueue: initialQueue,
      isConnected: false,
      isSyncing: false,
    };
  },

  actions: {
    // Persist queue mutations directly to LocalStorage
    persistQueue() {
      try {
        localStorage.setItem(LOCAL_STORAGE_KEY, JSON.stringify(this.pendingQueue));
      } catch (e) {
        console.error("LocalStorage write failed:", e);
      }
    },

    setConnectionStatus(status: boolean) {
      this.isConnected = status;
      if (status) {
        this.triggerQueueSync();
      }
    },

    // Enqueue scan offline
    addScanToQueue(eventId: string, attendeeId: string) {
      // Enforce queue limits (cap at 5,000 entries max) to prevent browser storage crashes
      if (this.pendingQueue.length >= 5000) {
        console.warn("Local offline scan queue capacity limit reached. Scan discarded.");
        return null;
      }

      const newScan: CheckInScan = {
        scanId: crypto.randomUUID(),
        eventId,
        attendeeId,
        scannedAt: new Date().toISOString(),
        syncStatus: "Queued",
        correlationId: `corr-${crypto.randomUUID()}`,
      };

      this.pendingQueue.push(newScan);
      this.persistQueue();
      return newScan;
    },

    // WebSocket broadcast receiver logic
    handleIncomingCheckIn(data: any) {
      console.log("Real-time check-in received:", data);
      // Remove from pending queue if it matches a locally queued sync scan
      this.pendingQueue = this.pendingQueue.filter(
        (scan) => scan.attendeeId !== data.attendeeId
      );
      this.persistQueue();
    },

    // Sequential Queue Synchronization and Flush Engine
    async triggerQueueSync() {
      if (this.isSyncing || this.pendingQueue.length === 0) return;
      this.isSyncing = true;

      console.log(`Starting automated sync for ${this.pendingQueue.length} pending scans...`);

      // Flush scans sequentially to preserve chronological event execution order
      while (this.pendingQueue.length > 0 && this.isConnected) {
        const currentScan = this.pendingQueue[0];
        currentScan.syncStatus = "Syncing";

        try {
          // Perform bulk sync call or single post request to registration validation backend
          await api.post("/api/registrations/sync", {
            eventId: currentScan.eventId,
            attendeeId: currentScan.attendeeId,
            scannedAt: currentScan.scannedAt,
            correlationId: currentScan.correlationId,
          });

          // Upon success, remove from memory queue and persist
          this.pendingQueue.shift();
          this.persistQueue();
          console.log(`Successfully synchronised scan: ${currentScan.scanId}`);
        } catch (error) {
          console.error(`Synchronisation failed for scan ${currentScan.scanId}:`, error);
          currentScan.syncStatus = "Queued";
          // Halt sync flush upon network/connection error to prevent dropouts
          break;
        }
      }

      this.isSyncing = false;
    },
  },
});
