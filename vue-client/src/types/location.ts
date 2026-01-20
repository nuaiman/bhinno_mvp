export interface Division {
  id: number
  name: string
}

export interface District {
  id: number
  division_id: number
  name: string
}

export interface Subdistrict {
  id: number
  district_id: number
  name: string
}
