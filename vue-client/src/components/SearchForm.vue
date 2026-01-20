<template>
  <div
    class="bg-[#0D1017]/80 backdrop-blur-2xl border border-white/10 p-4 md:p-6 rounded-[20px] xs:rounded-[40px] shadow-3xl w-80 sm:w-full max-w-md"
  >
    <!-- Location Selection -->
    <SearchFormLocation :fields="locationFields" />

    <hr class="border border-white/5 my-4 mx-2" />

    <!-- Category Selection -->
    <SearchFormCategory :fields="categoryFields" />

    <!-- Submit Button -->
    <button
      class="w-full mt-4 bg-blue-600 hover:bg-blue-500 rounded-2xl py-4 transition-all shadow-lg shadow-blue-600/30 flex items-center justify-center gap-2 group active:scale-95"
      @click="goToServices"
    >
      <svg
        class="w-5 h-5 stroke-white fill-none group-hover:scale-110 transition-transform"
        stroke-width="2.5"
        viewBox="0 0 24 24"
      >
        <circle cx="11" cy="11" r="8" />
        <path d="m21 21-4.3-4.3" />
      </svg>
      <span class="text-sm font-bold tracking-wide">Find Services</span>
    </button>
  </div>
</template>

<script setup lang="ts">
import SearchFormLocation from './SearchFormLocation.vue'
import SearchFormCategory from './SearchFormCategory.vue'

import { useRouter } from 'vue-router'
import { useLocationStore } from '@/stores/location'
import { useCategoryStore } from '@/stores/category'
import { useServiceStore } from '@/stores/service'

const router = useRouter()
const locationStore = useLocationStore()
const categoryStore = useCategoryStore()
const serviceStore = useServiceStore()

const locationFields = [
  { label: 'Division', placeholder: '' },
  { label: 'District', placeholder: '' },
  { label: 'Sub District', placeholder: '' },
]

const categoryFields = [
  { label: 'Category', placeholder: 'Electrician' },
  { label: 'Sub Category', placeholder: 'Wiring' },
]

const emit = defineEmits<{
  (e: 'close'): void
}>()

function goToServices() {
  // Validate all fields
  if (
    !locationStore.selectedDivision ||
    !locationStore.selectedDistrict ||
    !locationStore.selectedSubdistrict ||
    !categoryStore.selectedCategory ||
    !categoryStore.selectedSubCategory
  ) {
    alert('Please fill all fields first!')
    return
  }

  // Set filters
  serviceStore.division_id = locationStore.selectedDivision.id
  serviceStore.district_id = locationStore.selectedDistrict.id
  serviceStore.subdistrict_id = locationStore.selectedSubdistrict.id
  serviceStore.category_id = categoryStore.selectedCategory.id
  serviceStore.subcategory_id = categoryStore.selectedSubCategory.id

  // Close modal
  emit('close')

  // Go to services page
  router.push({ name: 'services' })
}
</script>
