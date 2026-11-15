import { createRouter, createWebHistory } from "vue-router";
import { useAuthStore } from "../stores/auth";

const routes = [
  {
    path: "/",
    name: "LandingPage",
    component: () => import("../pages/LandingPage.vue"),
  },
  {
    path: "/",
    component: () => import("../layouts/DashboardLayout.vue"),
    children: [
      {
        path: "dashboard",
        name: "OrganiserDashboard",
        component: () => import("../pages/OrganiserDashboard.vue"),
        meta: { requiresAuth: true },
      },
      {
        path: "check-in",
        name: "LiveCheckIn",
        component: () => import("../pages/LiveCheckIn.vue"),
        meta: { requiresAuth: true },
      },
    ]
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});

// Vue Router Navigation Guard enforcing protected paths
router.beforeEach(async (to) => {
  const authStore = useAuthStore();

  // If first visit, verify session validity with backend
  if (!authStore.isAuthenticated && authStore.user === null) {
    await authStore.checkSession();
  }

  const requiresAuth = to.matched.some((record) => record.meta.requiresAuth);

  if (requiresAuth && !authStore.isAuthenticated) {
    // Redirect unauthenticated requests to public landing page
    return { name: "LandingPage" };
  }
  
  return true;
});
