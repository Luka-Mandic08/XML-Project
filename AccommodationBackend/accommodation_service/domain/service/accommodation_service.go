package service

import (
	"accommodation_service/domain/model"
	"accommodation_service/domain/repository"
	accommodation "common/proto/accommodation_service"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccommodationService struct {
	accommodationStore repository.AccommodationStore
	availabilityStore  repository.AvailabilityStore
}

func NewAccommodationService(accommodationStore repository.AccommodationStore, availabilityStore repository.AvailabilityStore) *AccommodationService {
	return &AccommodationService{
		accommodationStore: accommodationStore,
		availabilityStore:  availabilityStore,
	}
}

func (service *AccommodationService) Get(id primitive.ObjectID) (*model.Accommodation, error) {
	return service.accommodationStore.GetById(id)
}

func (service *AccommodationService) Insert(accommodation *model.Accommodation) (*model.Accommodation, error) {
	acc, _ := service.accommodationStore.GetByAddress(accommodation.Address)
	if acc != nil {
		return nil, errors.New("an accommodation already exists in this location")
	}
	return service.accommodationStore.Insert(accommodation)
}

func (service *AccommodationService) Update(accommodation *model.Accommodation) (*mongo.UpdateResult, error) {
	return service.accommodationStore.Update(accommodation)
}

func (service *AccommodationService) Delete(id string) (*mongo.DeleteResult, error) {
	return service.accommodationStore.Delete(id)
}

func (service *AccommodationService) UpdateAvailability(request accommodation.UpdateAvailabilityRequest) error {
	dateFrom := request.DateFrom.AsTime()
	dateTo := request.DateTo.AsTime()
	for date := dateFrom; !date.After(dateTo); date = date.Add(time.Hour * 24) {
		err := service.availabilityStore.Upsert(&model.Availability{
			AccommodationId: request.Accommodationid,
			Date:            date,
			Price:           request.Price,
			IsAvailable:     true,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (service *AccommodationService) CheckAccommodationAvailability(request *accommodation.CheckAvailabilityRequest) (*accommodation.CheckAvailabilityResponse, error) {
	id, _ := primitive.ObjectIDFromHex(request.Accommodationid)
	acc, _ := service.accommodationStore.GetById(id)
	if acc == nil {
		return nil, errors.New("this accommodation does not exist")
	}
	if acc.MinGuests > request.NumberOfGuests {
		return nil, errors.New("this accommodation accepts a minimum of " + strconv.Itoa(int(acc.MinGuests)) + " guests")
	}
	if acc.MaxGuests < request.NumberOfGuests {
		return nil, errors.New("this accommodation accepts a maximum of " + strconv.Itoa(int(acc.MaxGuests)) + " guests")
	}

	totalPrice, availabilitiesToUpdate, err := service.CheckDateAvailability(request, acc)
	if err != nil {
		return nil, err
	}
	if acc.HasAutomaticReservations {
		err := service.ChangeAvailability(availabilitiesToUpdate, false)
		if err != nil {
			return nil, err
		}
	}
	return &accommodation.CheckAvailabilityResponse{
		ShouldCreateAutomaticReservation: acc.HasAutomaticReservations,
		TotalPrice:                       totalPrice,
	}, nil
}

func (service *AccommodationService) CheckDateAvailability(request *accommodation.CheckAvailabilityRequest, acc *model.Accommodation) (float32, []*model.Availability, error) {
	dateFrom := request.DateFrom.AsTime()
	dateTo := request.DateTo.AsTime()
	var totalPrice float32
	var availabilitiesToUpdate []*model.Availability
	for date := dateFrom; !date.After(dateTo); date = date.Add(time.Hour * 24) {
		av, err := service.availabilityStore.GetByDateAndAccommodation(request.Accommodationid, date)
		if av == nil {
			return 0, nil, errors.New("this accommodation is unavailable")
		}
		if err != nil {
			return 0, nil, err
		}
		if acc.PriceIsPerGuest {
			totalPrice += float32(request.NumberOfGuests) * av.Price
			continue
		}
		totalPrice += av.Price
		availabilitiesToUpdate = append(availabilitiesToUpdate, av)
	}
	return totalPrice, availabilitiesToUpdate, nil
}

func (service *AccommodationService) ChangeAvailability(availabilities []*model.Availability, isAvailable bool) error {
	for _, av := range availabilities {
		av.IsAvailable = isAvailable
		err := service.availabilityStore.Upsert(av)
		if err != nil {
			return err
		}
	}
	return nil
}

func (service *AccommodationService) Search(req *accommodation.SearchRequest) ([]*model.Accommodation, []float64, int, error) {
	dateFrom := req.DateFrom.AsTime()
	dateTo := req.DateTo.AsTime()
	numberOfDays := 0
	for date := dateFrom; !date.After(dateTo); date = date.Add(time.Hour * 24) {
		numberOfDays++
	}
	ids, prices, err := service.availabilityStore.FindAndGroupAvailableDates(dateFrom, dateTo, numberOfDays)
	if err != nil {
		return nil, nil, 0, err
	}
	var accommodations []*model.Accommodation
	var realPrices []float64
	for i, id := range ids {
		id, _ := primitive.ObjectIDFromHex(id)
		acc, _ := service.accommodationStore.GetById(id)
		if acc.MaxGuests >= req.NumberOfGuests && acc.MinGuests <= req.NumberOfGuests {
			if strings.Contains(strings.ToLower(acc.Address.City), strings.ToLower(req.City)) && strings.Contains(strings.ToLower(acc.Address.Country), strings.ToLower(req.Country)) {
				accommodations = append(accommodations, acc)
				realPrices = append(realPrices, prices[i])
			}
		}
	}
	if len(accommodations) == 0 {
		return nil, nil, 0, errors.New("no accommodations found")
	}
	return accommodations, realPrices, numberOfDays, nil
}

func (service *AccommodationService) GetAllByHostId(hostId string) ([]*model.Accommodation, error) {
	return service.accommodationStore.GetAllByHostId(hostId)
}
