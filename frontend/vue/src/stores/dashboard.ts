import { defineStore } from 'pinia'
import { ref } from 'vue'

/** Dashboard 页面状态 */
export const useDashboardStore = defineStore('dashboard', () => {
  const userTotal = ref('vue3/ts/tailwindcss 模版')

  return {
    userTotal,
  }
})
