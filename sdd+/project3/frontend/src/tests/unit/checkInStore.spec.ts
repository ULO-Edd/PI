import { describe, it, expect, beforeEach, vi } from "vitest";
import { setActivePinia, createPinia } from "pinia";
import { useCheckInStore } from "../../stores/checkIn";
import { api } from "../../services/api";

// Mock axios base instance
vi.mock("../../services/api", () => {
  return {
    api: {
      post: vi.fn(),
    },
  };
});

describe("Pinia CheckInStore Offline Queue & Auto-Sync Unit Tests", () => {
  beforeEach(() => {
    setActivePinia(createPinia());
    localStorage.clear();
    vi.clearAllMocks();
  });

  it("should initialize with an empty queue and correct defaults", () => {
    const store = useCheckInStore();
    expect(store.pendingQueue).toEqual([]);
    expect(store.isConnected).toBe(false);
    expect(store.isSyncing).toBe(false);
  });

  it("should safely queue scanned check-ins locally when offline", () => {
    const store = useCheckInStore();
    const eventId = "event-1111-2222";
    const attendeeId = "attendee-3333-4444";

    const scan = store.addScanToQueue(eventId, attendeeId);

    expect(scan).not.toBeNull();
    expect(store.pendingQueue.length).toBe(1);
    expect(store.pendingQueue[0].eventId).toBe(eventId);
    expect(store.pendingQueue[0].attendeeId).toBe(attendeeId);
    expect(store.pendingQueue[0].syncStatus).toBe("Queued");

    // Verify LocalStorage matches
    const stored = localStorage.getItem("sdd3_pending_checkins");
    expect(stored).not.toBeNull();
    expect(JSON.parse(stored!)[0].attendeeId).toBe(attendeeId);
  });

  it("should sequentially flush and sync offline scans immediately on reconnect", async () => {
    const store = useCheckInStore();
    
    // Setup mock responses for sync API POST requests
    vi.mocked(api.post).mockResolvedValue({ status: 200, data: { success: true } });

    // Queue 2 offline scans
    store.addScanToQueue("event-1", "attendee-1");
    store.addScanToQueue("event-1", "attendee-2");

    expect(store.pendingQueue.length).toBe(2);

    // Trigger online status transition
    store.isConnected = true; // Set connection state manually
    await store.triggerQueueSync(); // Trigger sync flush

    // Assert Axios post was called for each scan sequentially
    expect(api.post).toHaveBeenCalledTimes(2);
    expect(api.post).toHaveBeenNthCalledWith(1, "/api/registrations/sync", expect.objectContaining({
      attendeeId: "attendee-1",
    }));
    expect(api.post).toHaveBeenNthCalledWith(2, "/api/registrations/sync", expect.objectContaining({
      attendeeId: "attendee-2",
    }));

    // Verify queue is completely empty after successful sync
    expect(store.pendingQueue).toEqual([]);
    expect(JSON.parse(localStorage.getItem("sdd3_pending_checkins")!)).toEqual([]);
  });

  it("should halt the sync sequence and preserve remaining items on request failure", async () => {
    const store = useCheckInStore();

    // Mock first sync request succeeding, second request throwing an error
    vi.mocked(api.post)
      .mockResolvedValueOnce({ status: 200, data: { success: true } })
      .mockRejectedValueOnce(new Error("Network timeout"));

    store.addScanToQueue("event-1", "attendee-1");
    store.addScanToQueue("event-1", "attendee-2");

    store.isConnected = true;
    await store.triggerQueueSync();

    // Verify Axios post was attempted for both, but stopped on the second failure
    expect(api.post).toHaveBeenCalledTimes(2);

    // Verify queue only contains the second scan, preserved with Queued state
    expect(store.pendingQueue.length).toBe(1);
    expect(store.pendingQueue[0].attendeeId).toBe("attendee-2");
    expect(store.pendingQueue[0].syncStatus).toBe("Queued");

    // LocalStorage should reflect the un-synced scan
    expect(JSON.parse(localStorage.getItem("sdd3_pending_checkins")!)[0].attendeeId).toBe("attendee-2");
  });
});
