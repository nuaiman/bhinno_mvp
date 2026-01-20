<template>
  <section class="relative min-h-[85vh] flex items-center">
    <HomeHeroBackground
      image-src="https://images.unsplash.com/photo-1577495508048-b635879837f1?auto=format&fit=crop&q=80&w=1500"
      alt-text="Modern city background"
    />

    <div class="relative z-10 grid lg:grid-cols-2 gap-12 w-full px-3 py-6 sm:px-8 sm:py-8">
      <HomeHeroText />
      <div class="flex flex-col items-center lg:items-end w-full">
        <SearchForm />
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import HomeHeroBackground from '@/components/HomeHeroBackground.vue'
import HomeHeroText from '@/components/HomeHeroText.vue'
import SearchForm from '@/components/SearchForm.vue'
</script>
