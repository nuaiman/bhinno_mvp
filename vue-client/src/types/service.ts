export interface Service {
  id: number
  user_id: number
  user_name: string
  user_avatar: string
  social_media_1_name: string
  social_media_1_link: string
  social_media_2_name?: string
  social_media_2_link?: string
  division_id: number
  district_id: number
  subdistrict_id: number
  category_id: number
  subcategory_id: number
  title: string
  description: string
  created_at: string
  active: boolean
  price: string
  rating?: number
  dynamic_fields?: Record<string, string | undefined>
  contact_phone: string
  tags?: string[]
}
