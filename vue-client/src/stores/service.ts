import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { Service } from '@/types/service'

// JSON import
import servicesData from '@/assets/jsons/services/services.json'

export const useServiceStore = defineStore(
  'service',
  () => {
    // --- State ---
    const services = ref<Service[]>(servicesData.services)

    // Filters
    const division_id = ref<number | null>(null)
    const district_id = ref<number | null>(null)
    const subdistrict_id = ref<number | null>(null)
    const category_id = ref<number | null>(null)
    const subcategory_id = ref<number | null>(null)

    // --- Computed ---
    const isFilterComplete = computed(
      () =>
        division_id.value !== null &&
        district_id.value !== null &&
        subdistrict_id.value !== null &&
        category_id.value !== null &&
        subcategory_id.value !== null,
    )

    const filteredServices = computed(() => {
      return services.value.filter((s) => {
        if (division_id.value && s.division_id !== division_id.value) return false
        if (district_id.value && s.district_id !== district_id.value) return false
        if (subdistrict_id.value && s.subdistrict_id !== subdistrict_id.value) return false
        if (category_id.value && s.category_id !== category_id.value) return false
        if (subcategory_id.value && s.subcategory_id !== subcategory_id.value) return false
        return true
      })
    })

    // --- Actions ---
    function setFilter(
      filter: Partial<{
        division_id: number
        district_id: number
        subdistrict_id: number
        category_id: number
        subcategory_id: number
      }>,
    ) {
      if (filter.division_id !== undefined) {
        division_id.value = filter.division_id
        district_id.value = null
        subdistrict_id.value = null
      }
      if (filter.district_id !== undefined) {
        district_id.value = filter.district_id
        subdistrict_id.value = null
      }
      if (filter.subdistrict_id !== undefined) subdistrict_id.value = filter.subdistrict_id
      if (filter.category_id !== undefined) {
        category_id.value = filter.category_id
        subcategory_id.value = null
      }
      if (filter.subcategory_id !== undefined) subcategory_id.value = filter.subcategory_id
    }

    function resetFilters() {
      division_id.value = null
      district_id.value = null
      subdistrict_id.value = null
      category_id.value = null
      subcategory_id.value = null
    }

    return {
      services,
      division_id,
      district_id,
      subdistrict_id,
      category_id,
      subcategory_id,
      isFilterComplete,
      filteredServices,
      setFilter,
      resetFilters,
    }
  },
  {
    // --- Persistence ---
    persist: {
      key: 'service-filters',
    },
  },
)
