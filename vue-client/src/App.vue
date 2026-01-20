<template>
  <div
    class="min-h-screen bg-[#05070A] text-white font-sans selection:bg-blue-500/30 overflow-x-hidden"
  >
    <!-- Global Loader -->
    <GlobalLoader>
      <CircleSpinner />
    </GlobalLoader>

    <!-- App-wide container -->
    <div class="max-w-8xl mx-auto px-3 py-6 space-y-6 md:px-8 md:py-8 md:space-y-8">
      <NavBar />

      <!-- Page content -->
      <RouterView />
    </div>
  </div>
</template>

<script setup lang="ts">
import { RouterView } from 'vue-router'
import NavBar from '@/components/NavBar.vue'
import GlobalLoader from 'vue-global-loader/GlobalLoader.vue'
import CircleSpinner from 'vue-global-loader/CircleSpinner.vue'
</script>
