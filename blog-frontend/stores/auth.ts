import { defineStore } from 'pinia'

interface User {
  id: number
  username: string
  nickname: string
  avatar: string
}

interface AuthState {
  token: string | null
  user: User | null
}

export const useAuthStore = defineStore('auth', {
  state: (): AuthState => ({
    token: null,
    user: null,
  }),

  getters: {
    isAuthenticated: (state) => !!state.token,
  },

  actions: {
    // 初始化状态（从 localStorage 恢复）
    init() {
      if (import.meta.client) {
        const token = localStorage.getItem('token')
        const userStr = localStorage.getItem('user')
        if (token) {
          this.token = token
        }
        if (userStr) {
          try {
            this.user = JSON.parse(userStr)
          } catch (e) {
            this.user = null
          }
        }
      }
    },

    // 设置登录信息
    setAuth(token: string, user: User) {
      this.token = token
      this.user = user
      if (import.meta.client) {
        localStorage.setItem('token', token)
        localStorage.setItem('user', JSON.stringify(user))
      }
    },

    // 登出
    logout() {
      this.token = null
      this.user = null
      if (import.meta.client) {
        localStorage.removeItem('token')
        localStorage.removeItem('user')
      }
    },
  },
})
