<template>
  <div
    class="group bg-white/5 rounded-2xl p-4 border border-white/5 transition-all hover:border-blue-500/30 focus-within:border-blue-500/50 focus-within:bg-white/8 relative"
  >
    <label class="block text-[10px] text-gray-400 uppercase font-black mb-1 tracking-wider">
      {{ label }}
    </label>
    <div class="relative flex items-center">
      <select
        v-model="modelValue"
        @change="
          $emit('update:modelValue', modelValue)
          $emit('change')
        "
        :class="[!modelValue ? 'text-gray-500' : 'text-white']"
        class="bg-transparent w-full text-sm font-medium focus:outline-none appearance-none pr-8 cursor-pointer relative z-10"
      >
        <option :value="''" disabled class="bg-[#121212] text-gray-500">
          {{ placeholder }}
        </option>
        <option
          v-for="option in options"
          :key="option.id"
          :value="option.id"
          class="bg-[#1c1c1c] text-white"
        >
          {{ option.name }}
        </option>
      </select>

      <svg
        class="w-4 h-4 absolute right-0 text-gray-500 group-hover:text-blue-400 transition-colors pointer-events-none"
        fill="none"
        stroke="currentColor"
        viewBox="0 0 24 24"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2.5"
          d="M19 9l-7 7-7-7"
        />
      </svg>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const props = defineProps<{
  label: string
  placeholder: string
  options: Array<{ id: string | number; name: string }>
  modelValue: string | number
}>()

const modelValue = ref(props.modelValue)
</script>
