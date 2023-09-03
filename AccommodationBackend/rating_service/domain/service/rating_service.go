package service

import (
	"go.mongodb.org/mongo-driver/mongo"
	"rating_service/domain/model"
	"rating_service/domain/repository"
	"sort"

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
		println("Accommodation rating created: ID - ", result.AccommodationId)
	}
	if err != nil {
		return nil, err
	}

	err = service.guestAccommodationGraphStore.CreateOrUpdateGuestAccommodationConnection(result)
	if err != nil {
		println("Error occurred in graph database: ", err.Error())
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
	oldRating, _ := service.GetAccommodationRatingById(accommodationRating.Id)

	result, err := service.accommodationStore.Update(accommodationRating)
	if result.ModifiedCount == 1 {
		println("Accommodation rating updated: ID - ", accommodationRating.Id.Hex())
	}
	if err != nil {
		return nil, err
	}

	updatedResult, _ := service.GetAccommodationRatingById(accommodationRating.Id)

	err = service.guestAccommodationGraphStore.CreateOrUpdateGuestAccommodationConnection(updatedResult)
	if err != nil {
		println("Error occurred in graph database: ", err.Error())
		revertResult, err := service.accommodationStore.Update(oldRating) //revert change
		if revertResult.ModifiedCount == 1 {
			println("Accommodation rating reverted: ID - ", accommodationRating.Id.Hex())
		}
		if err != nil {
			println("Error occurred when reverting accommodation rating: ID - ", accommodationRating.Id.Hex())
			return nil, err
		}
		return nil, err
	}

	return result, nil
}

func (service *RatingService) DeleteAccommodationRating(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	ratingToDelete, err := service.GetAccommodationRatingById(id)
	if err != nil {
		return nil, err
	}

	deleteResult, err := service.accommodationStore.Delete(id)
	if err != nil {
		return nil, err
	}

	err = service.guestAccommodationGraphStore.DeleteGuestAndConnection(ratingToDelete.GuestId, ratingToDelete.AccommodationId, ratingToDelete.Id.Hex())
	if err != nil {
		return nil, err
	}

	return deleteResult, nil
}

func (service *RatingService) GetRecommendedAccommodations(guestId string) ([]model.RecommendedAccommodation, error) {
	accommodationIds, err := service.guestAccommodationGraphStore.RecommendAccommodationsForGuest(guestId)
	if err != nil {
		return nil, err
	}

	recommendedAccommodations := make([]model.RecommendedAccommodation, 0)

	// Iterate through the recommended accommodation IDs and get their average scores
	for _, accommodationId := range accommodationIds {
		averageScore, err := service.GetAverageScoreForAccommodation(accommodationId)
		if err != nil {
			// Handle the error if needed
			continue
		}

		// Add the accommodation ID and average score to the slice
		recommendedAccommodations = append(recommendedAccommodations, model.RecommendedAccommodation{
			AccommodationID: accommodationId,
			AverageScore:    averageScore,
		})
	}

	// Sort the recommended accommodations by score in descending order
	sort.SliceStable(recommendedAccommodations, func(i, j int) bool {
		return recommendedAccommodations[i].AverageScore > recommendedAccommodations[j].AverageScore
	})

	// Return only the top 10 recommended accommodations
	if len(recommendedAccommodations) > 10 {
		return recommendedAccommodations[:10], nil
	}

	return recommendedAccommodations, nil
}
