package service

import (
	"go.mongodb.org/mongo-driver/mongo"
	"rating_service/domain/model"
	"rating_service/domain/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RatingService struct {
	hostStore                    repository.HostRatingStore
	accommodationStore           repository.AccommodationRatingStore
	guestAccommodationGraphStore repository.GuestAccommodationGraphStore
}

func NewRatingService(hostStore repository.HostRatingStore, accommodationStore repository.AccommodationRatingStore, guestAccommodationGraphStore repository.GuestAccommodationGraphStore) *RatingService {
	return &RatingService{
		hostStore:                    hostStore,
		accommodationStore:           accommodationStore,
		guestAccommodationGraphStore: guestAccommodationGraphStore,
	}
}

func (service *RatingService) GetHostRatingById(id primitive.ObjectID) (*model.HostRating, error) {
	return service.hostStore.Get(id)
}

func (service *RatingService) GetAllForHost(hostId string) ([]*model.HostRating, error) {
	return service.hostStore.GetAllForHost(hostId)
}

func (service *RatingService) GetAverageScoreForHost(hostId string) (float32, error) {
	return service.hostStore.GetAverageScoreForHost(hostId)
}

func (service *RatingService) InsertHostRating(hostRating *model.HostRating) (*model.HostRating, error) {
	return service.hostStore.Insert(hostRating)
}

func (service *RatingService) UpdateHostRating(HostRating *model.HostRating) (*mongo.UpdateResult, error) {
	return service.hostStore.Update(HostRating)
}

func (service *RatingService) DeleteHostRating(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return service.hostStore.Delete(id)
}

func (service *RatingService) GetAccommodationRatingById(id primitive.ObjectID) (*model.AccommodationRating, error) {
	return service.accommodationStore.Get(id)
}

func (service *RatingService) GetAllForAccommodation(accommodationId string) ([]*model.AccommodationRating, error) {
	return service.accommodationStore.GetAllForAccommodation(accommodationId)
}

func (service *RatingService) GetAverageScoreForAccommodation(accommodationId string) (float32, error) {
	return service.accommodationStore.GetAverageScoreForAccommodation(accommodationId)
}

func (service *RatingService) InsertAccommodationRating(accommodationRating *model.AccommodationRating) (*model.AccommodationRating, error) {
	result, err := service.accommodationStore.Insert(accommodationRating)
	if result != nil {
		println("Accommodation rating created 0: ID - ", result.AccommodationId)
	}
	if err != nil {
		deleteResult, deleteErr := service.accommodationStore.Delete(result.Id)
		if deleteResult.DeletedCount > 0 {
			println("Accommodation rating deleted 0: ID - ", result.AccommodationId)
		}
		if deleteErr != nil {
			println("Error occurred when deleting accommodation rating 0: ID - ", result.AccommodationId)
			return nil, deleteErr
		}
		return nil, err
	}
	/*
		err = service.guestAccommodationGraphStore.CreateGuestNode(result.GuestId)
		if err != nil {
			println("Could not create GuestNode", err.Error())
			deleteResult, deleteErr := service.accommodationStore.Delete(result.Id)
			if deleteResult.DeletedCount > 0 {
				println("Accommodation rating deleted 1: ID - ", result.AccommodationId)
			}
			if deleteErr != nil {
				println("Error occurred when deleting accommodation rating 1: ID - ", result.AccommodationId)
				return nil, deleteErr
			}
			return nil, err
		}

		err = service.guestAccommodationGraphStore.CreateAccommodationNode(result.AccommodationId)
		if err != nil {
			println("Could not create AccommodationNode", err.Error())
			deleteResult, deleteErr := service.accommodationStore.Delete(result.Id)
			if deleteResult.DeletedCount > 0 {
				println("Accommodation rating deleted 2: ID - ", result.AccommodationId)
			}
			if deleteErr != nil {
				println("Error occurred when deleting accommodation rating 2: ID - ", result.AccommodationId)
				return nil, deleteErr
			}
			return nil, err
		}

		err = service.guestAccommodationGraphStore.CreateConnectionBetweenGuestAndAccommodation(result)
		if err != nil {
			println("Could not create ConnectionBetweenGuestAndAccommodation", err.Error())
			return nil, err
		}
	*/
	err = service.guestAccommodationGraphStore.CreateOrUpdateGuestAccommodationConnection(result)
	if err != nil {
		println("Could not create neo4j relation between nodes", err.Error())
		deleteResult, deleteErr := service.accommodationStore.Delete(result.Id)
		if deleteResult.DeletedCount > 0 {
			println("Accommodation rating deleted: ID - ", result.AccommodationId)
		}
		if deleteErr != nil {
			println("Error occurred when deleting accommodation rating: ID - ", result.AccommodationId)
			return nil, deleteErr
		}
		return nil, err
	}

	return result, nil
}

func (service *RatingService) UpdateAccommodationRating(accommodationRating *model.AccommodationRating) (*mongo.UpdateResult, error) {
	return service.accommodationStore.Update(accommodationRating)
}

func (service *RatingService) DeleteAccommodationRating(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return service.accommodationStore.Delete(id)
}
