package service

import (
	"accommodation_service/domain/model"
	"accommodation_service/domain/repository"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccommodationService struct {
	store repository.AccommodationStore
}

func NewAccommodationService(store repository.AccommodationStore) *AccommodationService {
	return &AccommodationService{
		store: store,
	}
}

func (service *AccommodationService) Get(id primitive.ObjectID) (*model.Accommodation, error) {
	return service.store.GetById(id)
}

func (service *AccommodationService) GetByUsername(username string) (*model.Accommodation, error) {
	return service.store.GetByUsername(username)
}

func (service *AccommodationService) Insert(accommodation *model.Accommodation) (*model.Accommodation, error) {
	return service.store.Insert(accommodation)
}

func (service *AccommodationService) Update(accommodation *model.Accommodation) (*mongo.UpdateResult, error) {
	return service.store.Update(accommodation)
}

func (service *AccommodationService) Delete(id string) (*mongo.DeleteResult, error) {
	return service.store.Delete(id)
}
