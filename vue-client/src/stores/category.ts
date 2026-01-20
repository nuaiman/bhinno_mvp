import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { Category, Subcategory } from '@/types/category'
import categoriesData from '@/assets/jsons/categories/categories.json'
import subCategoriesData from '@/assets/jsons/categories/subcategories.json'

export const useCategoryStore = defineStore(
  'category',
  () => {
    const categories = ref<Category[]>(categoriesData.categories)
    const subCategories = ref<Subcategory[]>(subCategoriesData.subcategories)

    const selectedCategory = ref<Category | null>(null)
    const selectedSubCategory = ref<Subcategory | null>(null)

    const filteredSubCategories = computed(() =>
      selectedCategory.value
        ? subCategories.value.filter((c) => c.category_id === selectedCategory.value!.id)
        : [],
    )

    function setCategory(categoryId: number) {
      selectedCategory.value = categories.value.find((d) => d.id === categoryId) || null
      selectedSubCategory.value = null
    }

    function setSubCategory(subCategoryId: number) {
      selectedSubCategory.value = subCategories.value.find((sc) => sc.id === subCategoryId) || null
    }

    return {
      categories,
      subCategories,
      selectedCategory,
      selectedSubCategory,
      filteredSubCategories,
      setCategory,
      setSubCategory,
    }
  },
  {
    persist: {
      key: 'category-filters',
    },
  },
)
