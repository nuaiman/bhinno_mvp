<template>
  <div class="flex items-center gap-2 cursor-pointer" @click="emit('click')">
    <!-- Squircle logo -->
    <div
      class="bg-blue-600 w-9 h-9 rounded-xl flex items-center justify-center font-bold shadow-lg shadow-blue-600/20 group-hover:scale-110 transition"
    >
      {{ logo }}
    </div>
    <!-- Optional text next to logo -->
    <span v-if="text" class="text-xl font-bold tracking-tight">{{ text }}</span>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  text: string
  logo: string
}>()

const emit = defineEmits<{
  (e: 'click'): void
}>()
</script>
