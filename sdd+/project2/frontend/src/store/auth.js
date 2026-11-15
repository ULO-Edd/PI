import { defineStore } from 'pinia'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: localStorage.getItem('token') || '',
    user: JSON.parse(localStorage.getItem('user')) || null,
  }),
  getters: {
    isAuthenticated: (state) => !!state.token,
    userRole: (state) => state.user?.role || '',
    userStatus: (state) => state.user?.status || '',
  },
  actions: {
    async login(username, password) {
      const res = await fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username, password }),
      })
      if (!res.ok) {
        const errMsg = await res.text()
        throw new Error(errMsg || 'Failed to login')
      }
      const data = await res.json()
      
      const payloadBase64 = data.token.split('.')[1]
      const payload = JSON.parse(atob(payloadBase64))

      this.token = data.token
      this.user = {
        id: payload.user_id,
        username: payload.username,
        role: payload.role,
        status: payload.status,
      }

      localStorage.setItem('token', this.token)
      localStorage.setItem('user', JSON.stringify(this.user))
    },
    async register(username, password, role) {
      const res = await fetch('http://localhost:8080/register', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username, password, role }),
      })
      if (!res.ok) {
        const errMsg = await res.text()
        throw new Error(errMsg || 'Failed to register')
      }
      return await res.json()
    },
    logout() {
      this.token = ''
      this.user = null
      localStorage.removeItem('token')
      localStorage.removeItem('user')
    }
  }
})
