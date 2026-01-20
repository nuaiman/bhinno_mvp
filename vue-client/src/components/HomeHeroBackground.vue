<!-- HomeHeroBackground.vue -->
<template>
  <div class="absolute inset-0 z-0">
    <div
      class="relative w-full h-full overflow-hidden rounded-3xl border border-white/10 shadow-2xl"
    >
      <img
        :src="imageSrc"
        :alt="altText"
        class="w-full h-full object-cover object-center grayscale-[0.5] brightness-[0.4]"
      />
      <div
        class="absolute inset-0 bg-linear-to-r from-[#05070A] via-[#05070A]/60 to-transparent"
      ></div>
      <div
        class="absolute inset-0 bg-linear-to-t from-[#05070A] via-transparent to-transparent"
      ></div>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  imageSrc: string
  altText?: string
}>()
</script>
