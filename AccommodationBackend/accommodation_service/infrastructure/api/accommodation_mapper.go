package api

import (
	"accommodation_service/domain/model"
	accommodation "common/proto/accommodation_service"
)

// Func za mapiranje objekata iz modela na proto message
func MapCreateRequestToAccommodation(req *accommodation.CreateRequest) *model.Accommodation {
	address := model.Address{
		Street:  req.Address.Street,
		City:    req.Address.City,
		Country: req.Address.Country,
	}
	acc := model.Accommodation{
		Name:                     req.Name,
		Address:                  address,
		Amenities:                req.Amenities,
		Images:                   req.Images,
		MinGuests:                req.MinGuests,
		MaxGuests:                req.MaxGuests,
		HostId:                   req.HostId,
		PriceIsPerGuest:          req.PriceIsPerGuest,
		HasAutomaticReservations: req.HasAutomaticReservations,
	}
	return &acc
}
