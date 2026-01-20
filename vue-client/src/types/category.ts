export interface Category {
  id: number
  name: string
}

export interface Subcategory {
  id: number
  category_id: number
  name: string
}
