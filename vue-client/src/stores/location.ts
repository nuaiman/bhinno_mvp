import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { Division, District, Subdistrict } from '@/types/location'
import divisionsData from '@/assets/jsons/location/bd/divisions.json'
import districtsData from '@/assets/jsons/location/bd/districts.json'
import subdistrictsData from '@/assets/jsons/location/bd/subdistricts.json'

export const useLocationStore = defineStore(
  'location',
  () => {
    const divisions = ref<Division[]>(divisionsData.divisions)
    const districts = ref<District[]>(districtsData.districts)
    const subdistricts = ref<Subdistrict[]>(subdistrictsData.subdistricts)

    const selectedDivision = ref<Division | null>(null)
    const selectedDistrict = ref<District | null>(null)
    const selectedSubdistrict = ref<Subdistrict | null>(null)

    const filteredDistricts = computed(() =>
      selectedDivision.value
        ? districts.value.filter((d) => d.division_id === selectedDivision.value!.id)
        : [],
    )

    const filteredSubdistricts = computed(() =>
      selectedDistrict.value
        ? subdistricts.value.filter((s) => s.district_id === selectedDistrict.value!.id)
        : [],
    )

    function setDivision(divisionId: number) {
      selectedDivision.value = divisions.value.find((d) => d.id === divisionId) || null
      selectedDistrict.value = null
      selectedSubdistrict.value = null
    }

    function setDistrict(districtId: number) {
      selectedDistrict.value = districts.value.find((d) => d.id === districtId) || null
      selectedSubdistrict.value = null
    }

    function setSubdistrict(subdistrictId: number) {
      selectedSubdistrict.value = subdistricts.value.find((s) => s.id === subdistrictId) || null
    }

    return {
      divisions,
      districts,
      subdistricts,
      selectedDivision,
      selectedDistrict,
      selectedSubdistrict,
      filteredDistricts,
      filteredSubdistricts,
      setDivision,
      setDistrict,
      setSubdistrict,
    }
  },
  {
    persist: {
      key: 'location-filters',
    },
  },
)
