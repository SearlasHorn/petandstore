package models

type PhotoUrlsList []string

type PetStatus string

const (
	AvailablePetStatus PetStatus = "available"
	PendingPetStatus   PetStatus = "pending"
	SoldPetStatus      PetStatus = "sold"
)

type Pet struct {
	Id        int64         `json:"id"`
	Category  Category      `json:"category"`
	Name      string        `json:"name"`
	PhotoUrls PhotoUrlsList `json:"photoUrls"`
	Tags      []Tag         `json:"tags"`
	Status    PetStatus     `json:"status"`
}


type Tag struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Category struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}