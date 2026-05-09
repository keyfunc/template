import { createApp } from 'vue'
import { createPinia } from 'pinia'
import '@/index.css'
import App from '@/App.vue'
import { router } from '@/router'
import { VueQueryPlugin } from '@tanstack/vue-query'
import { queryClient } from '@/service/tanstack'

const pinia = createPinia()
const app = createApp(App)

app.use(pinia)
app.use(router)
app.use(VueQueryPlugin, { queryClient })
app.mount('#app')
