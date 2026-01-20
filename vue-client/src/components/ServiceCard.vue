<template>
  <div
    class="bg-[#0D1017]/90 backdrop-blur-md border border-white/10 rounded-3xl p-5 shadow-md flex flex-col gap-4 hover:shadow-blue-500/30 transition-shadow duration-300 cursor-pointer"
  >
    <!-- User -->
    <div class="flex items-center gap-5">
      <img
        :src="service.user_avatar"
        alt="Avatar"
        class="w-24 h-24 rounded-2xl object-cover border-2 border-white/20"
      />
      <div class="w-full">
        <p class="font-bold text-white text-lg">{{ service.user_name }}</p>
        <div class="flex justify-between">
          <div class="flex gap-3 text-sm text-gray-400 mt-1">
            <a
              v-if="service.social_media_1_link"
              :href="service.social_media_1_link"
              target="_blank"
              class="hover:text-blue-400 transition-colors"
            >
              {{ service.social_media_1_name }}
            </a>
            <a
              v-if="service.social_media_2_link"
              :href="service.social_media_2_link"
              target="_blank"
              class="hover:text-blue-400 transition-colors"
            >
              {{ service.social_media_2_name }}
            </a>
          </div>
        </div>

        <div class="flex justify-end">
          <a
            :href="`tel:${service.contact_phone}`"
            target="_blank"
            class="text-sm text-gray-400 hover:text-blue-400 transition-colors"
          >
            Phone
          </a>
        </div>
      </div>
    </div>

    <!-- Title & Description -->
    <h3 class="font-semibold text-lg text-white mt-2">{{ service.title }}</h3>
    <p class="text-gray-300 text-sm line-clamp-3">{{ service.description }}</p>
  </div>
</template>

<script setup lang="ts">
import type { Service } from '@/types/service'

defineProps<{
  service: Service
}>()
</script>
