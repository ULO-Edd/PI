import { useCheckInStore } from "../stores/checkIn";

let socket: WebSocket | null = null;
let reconnectDelay = 1000;
const maxReconnectDelay = 30000;
let reconnectTimer: any = null;
let currentEventId: string | null = null;

export const wsService = {
  connect(eventId: string) {
    const checkInStore = useCheckInStore();
    currentEventId = eventId;

    if (socket && socket.readyState === WebSocket.OPEN) {
      this.subscribe(eventId);
      return;
    }

    const wsUrl = import.meta.env.VITE_WS_URL || "ws://localhost:5000/ws";
    socket = new WebSocket(wsUrl);

    socket.onopen = () => {
      console.log("WebSocket connected successfully");
      checkInStore.setConnectionStatus(true);
      reconnectDelay = 1000; // Reset delay on successful connect
      if (reconnectTimer) {
        clearTimeout(reconnectTimer);
        reconnectTimer = null;
      }
      this.subscribe(eventId);
    };

    socket.onmessage = (event) => {
      try {
        const data = JSON.parse(event.data);
        if (data.event === "attendee_checked_in") {
          checkInStore.handleIncomingCheckIn(data);
        }
      } catch (err) {
        console.error("Failed to parse WebSocket message:", err);
      }
    };

    socket.onclose = () => {
      console.log("WebSocket closed. Attempting reconnect...");
      checkInStore.setConnectionStatus(false);
      this.scheduleReconnect();
    };

    socket.onerror = (error) => {
      console.error("WebSocket error observed:", error);
      socket?.close();
    };
  },

  subscribe(eventId: string) {
    if (socket && socket.readyState === WebSocket.OPEN) {
      socket.send(JSON.stringify({ action: "subscribe", eventId }));
    }
  },

  disconnect() {
    if (reconnectTimer) {
      clearTimeout(reconnectTimer);
      reconnectTimer = null;
    }
    if (socket) {
      // Remove event listeners to prevent reconnect loops on intentional close
      socket.onopen = null;
      socket.onmessage = null;
      socket.onclose = null;
      socket.onerror = null;
      socket.close();
      socket = null;
    }
    const checkInStore = useCheckInStore();
    checkInStore.setConnectionStatus(false);
    currentEventId = null;
  },

  scheduleReconnect() {
    if (reconnectTimer) return;

    // Jitter calculation (+/- 200ms) to prevent massive client stampedes
    const jitter = Math.floor(Math.random() * 400) - 200;
    const finalDelay = Math.min(reconnectDelay, maxReconnectDelay) + jitter;

    console.log(`Scheduling reconnect in ${finalDelay}ms`);

    reconnectTimer = setTimeout(() => {
      reconnectTimer = null;
      // Exponential backoff multiplier
      reconnectDelay *= 2;
      if (currentEventId) {
        this.connect(currentEventId);
      }
    }, Math.max(100, finalDelay));
  },
};
