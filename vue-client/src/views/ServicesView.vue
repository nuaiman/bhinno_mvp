<template>
  <div class="py-4 md:py-6 relative">
    <!-- Search Form Modal -->
    <SearchFormModal v-if="showSearchModal" @close="closeModal" />

    <!-- Services List -->
    <div v-else>
      <!-- Filters & Stats Bar -->
      <div class="mb-8">
        <div class="relative overflow-hidden">
          <div class="flex flex-col gap-2">
            <!-- Location Breadcrumb -->
            <div
              class="flex flex-wrap items-center gap-1 text-xs font-bold uppercase tracking-widest text-blue-400/80"
            >
              <span>{{ divisionName }}</span>
              <span class="text-gray-600">/</span>
              <span>{{ districtName }}</span>
              <span class="text-gray-600">/</span>
              <span class="text-white">{{ subdistrictName }}</span>
            </div>

            <!-- Results Info -->
            <div class="flex justify-between items-center mt-2 -mb-2">
              <p class="text-xs font-medium text-gray-500 uppercase tracking-tighter leading-none">
                Selected Category
              </p>

              <div class="flex items-center gap-3">
                <span class="relative flex h-2 w-2">
                  <span
                    class="animate-ping absolute inline-flex h-full w-full rounded-full bg-blue-400 opacity-75"
                  ></span>
                  <span class="relative inline-flex rounded-full h-2 w-2 bg-blue-500"></span>
                </span>

                <span class="text-sm font-medium text-gray-200">
                  <span class="text-white font-bold">{{ filteredServices.length }}</span> Results
                  Found
                </span>
              </div>
            </div>

            <!-- Category & Subcategory -->
            <div class="flex flex-col md:flex-row md:items-center justify-between gap-4">
              <div>
                <h2 class="text-2xl md:text-3xl font-light tracking-tight text-white">
                  {{ categoryName }}
                </h2>
                <h3 class="text-md font-light tracking-tight text-white">
                  <span class="font-semibold text-blue-500">{{ subcategoryName }}</span>
                </h3>
              </div>

              <!-- Search Button -->
              <button
                class="mt-2 md:mt-0 bg-blue-600 hover:bg-blue-500 text-white rounded-xl px-4 py-2 transition-all shadow-md shadow-blue-600/30"
                @click="openModal"
              >
                🔍︎
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Empty State -->
      <div v-if="filteredServices.length === 0" class="text-gray-400 text-center py-10">
        No services found for the selected filters.
      </div>

      <!-- Services Grid -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        <ServiceCard v-for="service in filteredServices" :key="service.id" :service="service" />
      </div>

      <!-- Scroll to Top -->
      <button
        v-show="showScrollTop"
        @click="scrollToTop"
        class="fixed bottom-6 right-6 h-12 w-12 bg-blue-600 hover:bg-blue-500 text-white rounded-full shadow-lg shadow-blue-600/40 transition-all"
      >
        ↑
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useServiceStore } from '@/stores/service'
import { useLocationStore } from '@/stores/location'
import { useCategoryStore } from '@/stores/category'

import SearchFormModal from '@/components/SearchFormModal.vue'
import ServiceCard from '@/components/ServiceCard.vue'

const serviceStore = useServiceStore()
const locationStore = useLocationStore()
const categoryStore = useCategoryStore()

// --- Modal State ---
const showSearchModal = ref(false)

function openModal() {
  showSearchModal.value = true
}

function closeModal() {
  showSearchModal.value = false
}

// Auto-open modal on first mount if filters incomplete
onMounted(() => {
  if (!serviceStore.isFilterComplete) {
    showSearchModal.value = true
  }
})

// --- Scroll To Top ---
const showScrollTop = ref(false)

function handleScroll() {
  showScrollTop.value = window.scrollY > 300
}

function scrollToTop() {
  window.scrollTo({ top: 0, behavior: 'smooth' })
}

onMounted(() => {
  window.addEventListener('scroll', handleScroll)
})

onUnmounted(() => {
  window.removeEventListener('scroll', handleScroll)
})

// --- Computed ---
const filteredServices = computed(() => serviceStore.filteredServices)

const divisionName = computed(
  () => locationStore.divisions.find((d) => d.id === serviceStore.division_id)?.name ?? '-',
)
const districtName = computed(
  () => locationStore.districts.find((d) => d.id === serviceStore.district_id)?.name ?? '-',
)
const subdistrictName = computed(
  () => locationStore.subdistricts.find((s) => s.id === serviceStore.subdistrict_id)?.name ?? '-',
)
const categoryName = computed(
  () => categoryStore.categories.find((c) => c.id === serviceStore.category_id)?.name ?? '-',
)
const subcategoryName = computed(
  () =>
    categoryStore.subCategories.find((sc) => sc.id === serviceStore.subcategory_id)?.name ?? '-',
)
</script>
