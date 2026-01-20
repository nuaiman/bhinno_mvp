<template>
  <div class="space-y-4 min-w-0">
    <!-- Categories -->
    <CustomDropdown
      v-model="categoryStore.selectedCategory"
      label="Category"
      placeholder="Select Category"
      :options="categoriesOptions"
    />

    <!-- Sub Categories -->
    <CustomDropdown
      v-model="categoryStore.selectedSubCategory"
      label="Sub Category"
      placeholder="Select Sub Category"
      :options="subCategoriesOptions"
      :disabled="!categoryStore.selectedCategory"
    />
  </div>
</template>

<script setup lang="ts">
import CustomDropdown from '@/components/CustomDropdown.vue'
import { useCategoryStore } from '@/stores/category'
import { computed, watch } from 'vue'

const categoryStore = useCategoryStore()

// Options
const categoriesOptions = computed(() =>
  categoryStore.categories.map((c) => ({ id: c.id, name: c.name })),
)

const subCategoriesOptions = computed(() =>
  categoryStore.filteredSubCategories.map((sc) => ({ id: sc.id, name: sc.name })),
)

// Watchers to reset child
watch(
  () => categoryStore.selectedCategory,
  () => {
    categoryStore.selectedSubCategory = null
  },
)
</script>
