import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '../store/auth'

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/auth/LoginView.vue'),
    meta: { guestOnly: true }
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import('../views/auth/RegisterView.vue'),
    meta: { guestOnly: true }
  },
  {
    path: '/',
    name: 'Discovery',
    component: () => import('../views/attendee/DiscoveryView.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/dashboard',
    name: 'AttendeeDashboard',
    component: () => import('../views/attendee/DashboardView.vue'),
    meta: { requiresAuth: true, role: 'attendee' }
  },
  {
    path: '/events/:id',
    name: 'EventDetail',
    component: () => import('../views/attendee/EventDetailView.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/organizer/dashboard',
    name: 'OrganizerDashboard',
    component: () => import('../views/organizer/OrganizerDashboardView.vue'),
    meta: { requiresAuth: true, role: 'organiser' }
  },
  {
    path: '/organizer/events/:id/attendees',
    name: 'OrganizerAttendees',
    component: () => import('../views/organizer/EventAttendeesView.vue'),
    meta: { requiresAuth: true, role: 'organiser' }
  },
  {
    path: '/admin/dashboard',
    name: 'AdminDashboard',
    component: () => import('../views/admin/AdminDashboardView.vue'),
    meta: { requiresAuth: true, role: 'administrator' }
  },
  {
    path: '/admin/organizers',
    name: 'AdminOrganizers',
    component: () => import('../views/admin/OrganizerApprovalView.vue'),
    meta: { requiresAuth: true, role: 'administrator' }
  },
  {
    path: '/admin/categories',
    name: 'AdminCategories',
    component: () => import('../views/admin/CategoryManagementView.vue'),
    meta: { requiresAuth: true, role: 'administrator' }
  },
  {
    path: '/admin/stats',
    name: 'AdminStats',
    component: () => import('../views/admin/GlobalStatsView.vue'),
    meta: { requiresAuth: true, role: 'administrator' }
  },
  {
    path: '/forbidden',
    name: 'Forbidden',
    component: () => import('../views/auth/ForbiddenView.vue')
  }
]

export const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()

  if (!authStore.isAuthenticated) {
    if (to.meta.requiresAuth) {
      return next('/login')
    }
    return next()
  }

  // If user is authenticated...
  if (to.meta.guestOnly) {
    if (authStore.userRole === 'administrator') return next('/admin/dashboard')
    if (authStore.userRole === 'organiser') return next('/organizer/dashboard')
    return next('/')
  }

  if (to.path === '/') {
    if (authStore.userRole === 'administrator') return next('/admin/dashboard')
    if (authStore.userRole === 'organiser') return next('/organizer/dashboard')
    return next()
  }

  if (to.meta.role && authStore.userRole !== to.meta.role && authStore.userRole !== 'administrator') {
    return next('/forbidden')
  }

  next()
})
