package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Subdistrict struct
type Subdistrict struct {
	ID         int    `json:"id"`
	DistrictID int    `json:"district_id"`
	Name       string `json:"name"`
}

// District struct
type District struct {
	ID           int           `json:"id"`
	DivisionID   int           `json:"division_id"`
	Name         string        `json:"name"`
	Subdistricts []Subdistrict `json:"subdistricts,omitempty"`
}

// Division struct
type Division struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Districts []District `json:"districts,omitempty"`
}

// Wrapper structs for reading JSON
type DivisionsFile struct {
	Divisions []Division `json:"divisions"`
}

type DistrictsFile struct {
	Districts []District `json:"districts"`
}

type SubdistrictsFile struct {
	Subdistricts []Subdistrict `json:"subdistricts"`
}

func main() {
	// 1️⃣ Read Divisions
	divBytes, err := ioutil.ReadFile("divisions.json")
	if err != nil {
		log.Fatal(err)
	}
	var divFile DivisionsFile
	if err := json.Unmarshal(divBytes, &divFile); err != nil {
		log.Fatal(err)
	}

	// 2️⃣ Read Districts
	disBytes, err := ioutil.ReadFile("districts.json")
	if err != nil {
		log.Fatal(err)
	}
	var disFile DistrictsFile
	if err := json.Unmarshal(disBytes, &disFile); err != nil {
		log.Fatal(err)
	}

	// 3️⃣ Read Subdistricts
	subBytes, err := ioutil.ReadFile("subdistricts.json")
	if err != nil {
		log.Fatal(err)
	}
	var subFile SubdistrictsFile
	if err := json.Unmarshal(subBytes, &subFile); err != nil {
		log.Fatal(err)
	}

	// 4️⃣ Nest subdistricts into districts
	for i := range disFile.Districts {
		var children []Subdistrict
		for _, sub := range subFile.Subdistricts {
			if sub.DistrictID == disFile.Districts[i].ID {
				children = append(children, sub)
			}
		}
		disFile.Districts[i].Subdistricts = children
	}

	// 5️⃣ Nest districts into divisions
	for i := range divFile.Divisions {
		var children []District
		for _, d := range disFile.Districts {
			if d.DivisionID == divFile.Divisions[i].ID {
				children = append(children, d)
			}
		}
		divFile.Divisions[i].Districts = children
	}

	// 6️⃣ Write nested JSON
	outBytes, err := json.MarshalIndent(divFile, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("bd_nested.json", outBytes, 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ bd_nested.json created successfully!")
}
