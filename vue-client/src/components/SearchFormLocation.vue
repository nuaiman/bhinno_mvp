<template>
  <div class="space-y-4 min-w-0">
    <!-- Division -->
    <CustomDropdown
      v-model="locationStore.selectedDivision"
      label="Division"
      placeholder="Select Division"
      :options="divisionsOptions"
    />

    <!-- District -->
    <CustomDropdown
      v-model="locationStore.selectedDistrict"
      label="District"
      placeholder="Select District"
      :options="districtsOptions"
      :disabled="!locationStore.selectedDivision"
    />

    <!-- Sub District -->
    <CustomDropdown
      v-model="locationStore.selectedSubdistrict"
      label="Sub District"
      placeholder="Select Sub District"
      :options="subdistrictsOptions"
      :disabled="!locationStore.selectedDistrict"
    />
  </div>
</template>

<script setup lang="ts">
import CustomDropdown from '@/components/CustomDropdown.vue'
import { useLocationStore } from '@/stores/location'
import { computed, watch } from 'vue'

const locationStore = useLocationStore()

const divisionsOptions = computed(() =>
  locationStore.divisions.map((d) => ({ id: d.id, name: d.name })),
)

const districtsOptions = computed(() =>
  locationStore.filteredDistricts.map((d) => ({ id: d.id, name: d.name })),
)

const subdistrictsOptions = computed(() =>
  locationStore.filteredSubdistricts.map((s) => ({ id: s.id, name: s.name })),
)

watch(
  () => locationStore.selectedDivision,
  () => {
    locationStore.selectedDistrict = null
    locationStore.selectedSubdistrict = null
  },
)

watch(
  () => locationStore.selectedDistrict,
  () => {
    locationStore.selectedSubdistrict = null
  },
)
</script>
