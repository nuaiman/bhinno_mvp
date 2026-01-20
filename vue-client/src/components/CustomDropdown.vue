<!-- CustomDropdown.vue -->
<template>
  <div class="relative w-full" ref="dropdownRef">
    <!-- Label -->
    <label class="block text-[10px] text-gray-400 uppercase font-black mb-2">
      {{ props.label }}
    </label>

    <!-- Input -->
    <div
      class="bg-white/5 rounded-2xl p-4 border border-white/5 text-white cursor-pointer flex justify-between items-center transition-all hover:border-blue-500/30 focus-within:border-blue-500/50 truncate"
      :class="{ 'opacity-50 cursor-not-allowed': props.disabled }"
      @click="!props.disabled && (open = !open)"
    >
      <span class="truncate max-w-full block">{{ selected?.name || props.placeholder }}</span>
      <svg
        class="w-4 h-4 text-gray-400 transition-transform"
        :class="{ 'rotate-180': open }"
        fill="none"
        stroke="currentColor"
        viewBox="0 0 24 24"
        stroke-width="2.5"
      >
        <path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7" />
      </svg>
    </div>

    <!-- Dropdown Menu -->
    <ul
      v-if="open && !props.disabled"
      class="absolute z-50 mt-1 w-full bg-[#0D1017]/80 border border-white/10 rounded-2xl max-h-60 overflow-y-auto shadow-3xl"
      style="scrollbar-width: thin; scrollbar-color: rgba(59, 130, 246, 0.5) transparent"
    >
      <li
        v-for="item in props.options"
        :key="item.id"
        class="p-3 hover:bg-blue-500/20 cursor-pointer transition-colors rounded-xl truncate"
        @click="selectItem(item)"
      >
        {{ item.name }}
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, onMounted, onUnmounted } from 'vue'

const props = defineProps<{
  label: string
  options: { id: number; name: string }[]
  modelValue?: { id: number; name: string } | null
  placeholder?: string
  disabled?: boolean
}>()

const emit = defineEmits<{
  (e: 'update:modelValue', value: { id: number; name: string } | null): void
}>()

const open = ref(false)
const selected = ref<{ id: number; name: string } | null>(props.modelValue || null)
const dropdownRef = ref<HTMLElement | null>(null)

watch(
  () => props.modelValue,
  (val) => (selected.value = val || null),
  { immediate: true },
)

function selectItem(item: { id: number; name: string }) {
  selected.value = item
  emit('update:modelValue', item)
  open.value = false
}

function handleClickOutside(event: MouseEvent) {
  if (dropdownRef.value && !dropdownRef.value.contains(event.target as Node)) {
    open.value = false
  }
}

onMounted(() => document.addEventListener('click', handleClickOutside))
onUnmounted(() => document.removeEventListener('click', handleClickOutside))
</script>
