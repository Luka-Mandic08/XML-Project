package api

import (
	rating "common/proto/rating_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rating_service/domain/model"
)

func MapHostRatingToResponse(r *model.HostRating) *rating.HostRating {
	var mappedHostRating = rating.HostRating{
		Id:      r.Id.Hex(),
		GuestId: r.GuestId,
		HostId:  r.HostId,
		Score:   r.Score,
		Comment: r.Comment,
	}
	return &mappedHostRating
}

func MapManyHostRatingsToResponse(rs []*model.HostRating) *rating.GetAllRatingsForHostResponse {
	var ratings []*rating.HostRating

	for _, r := range rs {
		var mappedHostRating = rating.HostRating{
			Id:      r.Id.Hex(),
			GuestId: r.GuestId,
			HostId:  r.HostId,
			Score:   r.Score,
			Comment: r.Comment,
		}
		ratings = append(ratings, &mappedHostRating)
	}

	return &rating.GetAllRatingsForHostResponse{HostRatings: ratings}
}

func MapCreateRequestToHostRating(r *rating.CreateHostRatingRequest) *model.HostRating {
	var mappedHostRating = model.HostRating{
		GuestId: r.GuestId,
		HostId:  r.HostId,
		Score:   r.Score,
		Comment: r.Comment,
	}
	return &mappedHostRating
}

func MapToHostRating(r *rating.HostRating, objectId primitive.ObjectID) *model.HostRating {
	var mappedHostRating = model.HostRating{
		Id:      objectId,
		GuestId: r.GuestId,
		HostId:  r.HostId,
		Score:   r.Score,
		Comment: r.Comment,
	}
	return &mappedHostRating
}

func MapAccommodationRatingToResponse(r *model.AccommodationRating) *rating.AccommodationRating {
	var mappedAccommodationRating = rating.AccommodationRating{
		Id:              r.Id.Hex(),
		GuestId:         r.GuestId,
		AccommodationId: r.AccommodationId,
		Score:           r.Score,
		Comment:         r.Comment,
	}
	return &mappedAccommodationRating
}

func MapManyAccommodationRatingsToResponse(rs []*model.AccommodationRating) *rating.GetAllRatingsForAccommodationResponse {
	var ratings []*rating.AccommodationRating

	for _, r := range rs {
		var mappedAccommodationRating = rating.AccommodationRating{
			Id:              r.Id.Hex(),
			GuestId:         r.GuestId,
			AccommodationId: r.AccommodationId,
			Score:           r.Score,
			Comment:         r.Comment,
		}
		ratings = append(ratings, &mappedAccommodationRating)
	}

	return &rating.GetAllRatingsForAccommodationResponse{AccommodationRatings: ratings}
}

func MapCreateRequestToAccommodationRating(r *rating.CreateAccommodationRatingRequest) *model.AccommodationRating {
	var mappedAccommodationRating = model.AccommodationRating{
		GuestId:         r.GuestId,
		AccommodationId: r.AccommodationId,
		Score:           r.Score,
		Comment:         r.Comment,
	}
	return &mappedAccommodationRating
}

func MapToAccommodationRating(r *rating.AccommodationRating, objectId primitive.ObjectID) *model.AccommodationRating {
	var mappedAccommodationRating = model.AccommodationRating{
		Id:              objectId,
		GuestId:         r.GuestId,
		AccommodationId: r.AccommodationId,
		Score:           r.Score,
		Comment:         r.Comment,
	}
	return &mappedAccommodationRating
}
