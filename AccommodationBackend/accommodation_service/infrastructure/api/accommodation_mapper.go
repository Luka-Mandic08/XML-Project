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

func MapAccommodationsToSearchRequest(accs []*model.Accommodation, prices []float64, numberOfDays, numberOfGuests int) *accommodation.SearchResponse {
	var searchAccommodations = []*accommodation.SearchResultAccommodation{}
	for i, acc := range accs {
		address := accommodation.Address{
			Street:  acc.Address.Street,
			City:    acc.Address.City,
			Country: acc.Address.Country,
		}
		var searchAccommodation = accommodation.SearchResultAccommodation{
			Id:         acc.Id.Hex(),
			Name:       acc.Name,
			Address:    &address,
			Amenities:  acc.Amenities,
			Images:     acc.Images,
			UnitPrice:  float32(prices[i] / float64(numberOfDays)),
			TotalPrice: float32(prices[i]),
		}
		if acc.PriceIsPerGuest {
			searchAccommodation.TotalPrice *= float32(numberOfGuests)
			searchAccommodation.UnitPrice *= float32(numberOfGuests)
		}
		searchAccommodations = append(searchAccommodations, &searchAccommodation)
	}
	return &accommodation.SearchResponse{Accommodations: searchAccommodations}
}
