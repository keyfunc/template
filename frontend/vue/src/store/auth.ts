import { defineStore } from 'pinia'
import { computed, ref } from 'vue'

const tokenKey = 'authToken'

export const useAuthStore = defineStore('auth', () => {
  const token = ref<string | null>(localStorage.getItem(tokenKey))
  const isLogin = computed(() => Boolean(token.value))

  // 设置本地 token
  function setToken(tokenValue: string) {
    token.value = tokenValue
    localStorage.setItem(tokenKey, tokenValue)
  }

  // 清除本地 token
  function clearToken() {
    token.value = null
    localStorage.removeItem(tokenKey)
  }

  return { token, isLogin, setToken, clearToken }
})
