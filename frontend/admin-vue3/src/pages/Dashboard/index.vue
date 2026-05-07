<script setup lang="ts">
import { useDashboardStore } from '@/stores/dashboard'

const dashboardStore = useDashboardStore()
</script>

<template>
  <div class="grid h-screen place-content-center">{{ dashboardStore.userTotal }}</div>
</template>
