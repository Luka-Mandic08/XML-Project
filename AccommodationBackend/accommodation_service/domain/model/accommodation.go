package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

type Accommodation struct {
	Id                       primitive.ObjectID `bson:"_id,omitempty"`
	Name                     string             `bson:"name"`
	Address                  Address            `bson:"address"`
	Amenities                []string           `bson:"amenities"`
	Images                   []string           `bson:"image"`
	MinGuests                int32              `bson:"minGuests"`
	MaxGuests                int32              `bson:"maxGuests"`
	HostId                   string             `bson:"hostid"`
	PriceIsPerGuest          bool               `bson:"priceIsPerGuest"`
	HasAutomaticReservations bool               `bson:"hasAutomaticReservations"`
}

func (a *Accommodation) ContainsAllAmenities(amenities []string) bool {
	for _, amenity := range amenities {
		for i, accommodationAmenity := range a.Amenities {
			if strings.Contains(strings.ToLower(accommodationAmenity), strings.ToLower(amenity)) {
				break
			}
			if i == len(a.Amenities)-1 {
				return false
			}
		}
	}
	return true
}

func (a *Accommodation) CheckPrice(price float64, maxPrice float32, numberOfGuests int32) bool {
	if maxPrice == 0 {
		return true
	}
	if a.PriceIsPerGuest {
		return price*float64(numberOfGuests) <= float64(maxPrice)
	}
	return price <= float64(maxPrice)
}
