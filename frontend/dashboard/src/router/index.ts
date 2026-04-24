import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'

/** 路由配置 */
const routes: RouteRecordRaw[] = [
  {
    path: '/',
    redirect: '/Dashboard',
  },
  {
    path: '/Dashboard',
    name: 'Dashboard',
    component: () => import('@/pages/Dashboard/index.vue'),
    meta: {
      title: 'Dashboard',
    },
  },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

export default router
