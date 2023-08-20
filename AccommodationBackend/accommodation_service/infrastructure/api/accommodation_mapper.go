package api

import (
	"accommodation_service/domain/model"
	accommodation "common/proto/accommodation_service"
	"github.com/golang/protobuf/ptypes"
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

func MapAccommodationsToSearchRequest(accs []*model.Accommodation, prices []float64, numberOfDays, numberOfGuests, page int) *accommodation.SearchResponse {
	var searchAccommodations = []*accommodation.SearchResultAccommodation{}
	startIndex := (page - 1) * 9
	endIndex := startIndex + 9
	if endIndex > len(accs) {
		endIndex = len(accs)
	}

	for i := startIndex; i < endIndex; i++ {
		acc := accs[i]
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

func MapAccommodations(accs []*model.Accommodation) (*accommodation.GetAllByHostIdResponse, *accommodation.GetAllResponse) {
	var accommodations = []*accommodation.Accommodation{}
	for _, acc := range accs {
		address := accommodation.Address{
			Street:  acc.Address.Street,
			City:    acc.Address.City,
			Country: acc.Address.Country,
		}
		var newAccommodation = accommodation.Accommodation{
			Name:                     acc.Name,
			Address:                  &address,
			Amenities:                acc.Amenities,
			Images:                   acc.Images,
			MinGuests:                acc.MinGuests,
			MaxGuests:                acc.MaxGuests,
			PriceIsPerGuest:          acc.PriceIsPerGuest,
			HasAutomaticReservations: acc.HasAutomaticReservations,
			HostId:                   acc.HostId,
			Id:                       acc.Id.Hex(),
		}

		accommodations = append(accommodations, &newAccommodation)
	}
	return &accommodation.GetAllByHostIdResponse{Accommodations: accommodations}, &accommodation.GetAllResponse{Accommodations: accommodations}
}

func MapAccommodation(acc *model.Accommodation) *accommodation.GetByIdResponse {
	address := accommodation.Address{
		Street:  acc.Address.Street,
		City:    acc.Address.City,
		Country: acc.Address.Country,
	}
	var newAccommodation = accommodation.Accommodation{
		Name:                     acc.Name,
		Address:                  &address,
		Amenities:                acc.Amenities,
		Images:                   acc.Images,
		MinGuests:                acc.MinGuests,
		MaxGuests:                acc.MaxGuests,
		PriceIsPerGuest:          acc.PriceIsPerGuest,
		HasAutomaticReservations: acc.HasAutomaticReservations,
		HostId:                   acc.HostId,
		Id:                       acc.Id.Hex(),
	}

	return &accommodation.GetByIdResponse{Accommodation: &newAccommodation}
}

func MapAvailabilities(availabilities []*model.Availability) *accommodation.GetAvailabilitiesResponse {
	var a = []*accommodation.Availability{}
	for _, av := range availabilities {
		protoTimestamp, _ := ptypes.TimestampProto(av.Date)
		var availability = accommodation.Availability{
			Date:        protoTimestamp,
			IsAvailable: av.IsAvailable,
			Price:       av.Price,
		}
		a = append(a, &availability)
	}
	return &accommodation.GetAvailabilitiesResponse{AvailabilityDates: a}
}
