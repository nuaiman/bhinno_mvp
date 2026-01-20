<template>
  <!-- Overlay -->
  <div class="fixed inset-0 z-50 bg-transparent backdrop-blur-2xl py-4">
    <!-- Scrollable modal content -->
    <div class="flex justify-center w-full shadow-3xl overflow-y-auto">
      <SearchForm @close="$emit('close')" />
    </div>
  </div>
</template>
<script setup lang="ts">
import SearchForm from '@/components/SearchForm.vue'
</script>
