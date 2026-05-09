<script setup lang="ts">
import { VueQueryDevtools } from '@tanstack/vue-query-devtools'
</script>

<template>
  <div class="grid h-screen place-content-center">
    <div class="flex gap-2">
      <RouterLink to="/">Go to Index</RouterLink>
      <RouterLink to="/about">Go to About</RouterLink>
    </div>
    <RouterView />
    <VueQueryDevtools />
  </div>
</template>
