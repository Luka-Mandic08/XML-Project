package service

import (
	"accommodation_service/domain/model"
	"accommodation_service/domain/repository"
	accommodation "common/proto/accommodation_service"
	reservation "common/proto/reservation_service"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"time"
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
		return nil, errors.New("this accommodation accepts a minimum of " + strconv.Itoa(int(acc.MinGuests)) + " guests, requested " + strconv.Itoa(int(request.NumberOfGuests)))
	}
	if acc.MaxGuests < request.NumberOfGuests {
		return nil, errors.New("this accommodation accepts a maximum of " + strconv.Itoa(int(acc.MaxGuests)) + " guests, requested " + strconv.Itoa(int(request.NumberOfGuests)))
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
			println(err.Error())
			return 0, nil, err
		}
		if acc.PriceIsPerGuest {
			totalPrice += float32(request.NumberOfGuests) * av.Price
			availabilitiesToUpdate = append(availabilitiesToUpdate, av)
			continue
		}
		totalPrice += av.Price
		availabilitiesToUpdate = append(availabilitiesToUpdate, av)
	}
	return totalPrice, availabilitiesToUpdate, nil
}

func (service *AccommodationService) ChangeAvailability(availabilities []*model.Availability, isAvailable bool) error {

	for i := 0; i < len(availabilities); i++ {
		av := availabilities[i]
		av.IsAvailable = isAvailable
		err := service.availabilityStore.Upsert(av)
		if err != nil {
			return err
		}
	}
	return nil
}

func (service *AccommodationService) Search(req *accommodation.SearchRequest, hostIds *reservation.GetAllOutstandingHostsResponse) ([]*model.Accommodation, []float64, int, error) {
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
		objectId, _ := primitive.ObjectIDFromHex(id)
		println("ACC: ", id, " PRICE: ", prices[i])
		acc, _ := service.accommodationStore.GetForSearch(objectId, req, hostIds.GetIds())
		if acc != nil && acc.CheckPrice(prices[i], req.MaxPrice, req.NumberOfGuests) && acc.ContainsAllAmenities(req.GetAmenities()) {
			accommodations = append(accommodations, acc)
			realPrices = append(realPrices, prices[i])
		}
	}

	if len(accommodations) == 0 {
		return nil, nil, 0, errors.New("no accommodations found")
	}
	return accommodations, realPrices, numberOfDays, nil
}

func (service *AccommodationService) CheckAccommodationExists(id primitive.ObjectID) error {
	_, err := service.accommodationStore.GetById(id)
	return err
}

func (service *AccommodationService) GetAllAvailability(dateFrom time.Time, dateTo time.Time, accommodationId string) []*model.Availability {
	var availabilitiesToUpdate []*model.Availability
	for date := dateFrom; !date.After(dateTo); date = date.Add(time.Hour * 24) {
		av, _ := service.availabilityStore.GetByDateAndAccommodation(accommodationId, date)
		availabilitiesToUpdate = append(availabilitiesToUpdate, av)
	}
	return availabilitiesToUpdate
}

func (service *AccommodationService) GetAllByHostId(hostId string) ([]*model.Accommodation, error) {
	return service.accommodationStore.GetAllByHostId(hostId)
}

func (service *AccommodationService) GetById(id primitive.ObjectID) (*model.Accommodation, error) {
	return service.accommodationStore.GetById(id)
}

func (service *AccommodationService) GetAll(page int) ([]*model.Accommodation, error) {
	return service.accommodationStore.GetAll(page)
}

func (service *AccommodationService) GetAvailabilitiesForAccommodation(request *accommodation.GetAvailabilitiesRequest) ([]*model.Availability, error) {
	dateFrom := request.DateFrom.AsTime()
	dateTo := request.DateTo.AsTime()
	return service.availabilityStore.GetAvailabilitiesForAccommodation(dateFrom, dateTo, request.Accommodationid)
}

func (service *AccommodationService) DeleteAllForHost(hostId string) (*mongo.DeleteResult, error) {
	accommodations, _ := service.accommodationStore.GetAllByHostId(hostId)
	for _, a := range accommodations {
		service.availabilityStore.DeleteAllForAccommodation(a.Id.Hex())
	}
	return service.accommodationStore.DeleteAllForHost(hostId)
}

func (service *AccommodationService) CheckCanApprove(request *accommodation.CheckCanApproveRequest) (*accommodation.CheckCanApproveResponse, error) {
	var availabilitiesToUpdate []*model.Availability
	for date := request.Start.AsTime(); !date.After(request.End.AsTime()); date = date.Add(time.Hour * 24) {
		av, err := service.availabilityStore.GetByDateAndAccommodation(request.AccommodationId, date)
		if av == nil {
			response := accommodation.CheckCanApproveResponse{
				CanApprove: "false",
			}
			return &response, nil
		}
		if err != nil {
			return nil, err
		}
		availabilitiesToUpdate = append(availabilitiesToUpdate, av)
	}
	err := service.ChangeAvailability(availabilitiesToUpdate, false)
	if err != nil {
		return nil, err
	}
	response := accommodation.CheckCanApproveResponse{
		CanApprove: "true",
	}
	return &response, nil
}

func (service *AccommodationService) GetAndCancelAllAvailabilitiesToCancel(request *accommodation.GetAndCancelAllAvailabilitiesToCancelRequest) (*accommodation.GetAndCancelAllAvailabilitiesToCancelResponse, error) {
	toChange := []*model.Availability{}
	response := []string{}
	for date := request.Start.AsTime(); !date.After(request.End.AsTime()); date = date.Add(time.Hour * 24) {
		av, err := service.availabilityStore.GetByDateAndAccommodationAllToCancel(request.AccommodationId, date)
		if err != nil {
			return nil, err
		}
		toChange = append(toChange, av)
		response = append(response, av.Id.Hex())
	}

	err := service.ChangeAvailability(toChange, true)
	if err != nil {
		return nil, err
	}
	return &accommodation.GetAndCancelAllAvailabilitiesToCancelResponse{ToCancel: response}, nil
}

func (service *AccommodationService) GetAllForHostByAccommodationId(request *accommodation.GetByIdRequest) (*accommodation.GetAllForHostByAccommodationIdResponse, error) {
	objectId, err := primitive.ObjectIDFromHex(request.GetId())
	if err != nil {
		return nil, err
	}
	ids, hostId, err := service.accommodationStore.GetAllForHostByAccommodationId(objectId)
	if err != nil {
		return nil, err
	}
	return &accommodation.GetAllForHostByAccommodationIdResponse{AccommodationIds: ids, HostId: hostId}, nil
}

func (service *AccommodationService) GetAllAvailabilitiesForRevering(request *accommodation.CheckAvailabilityRequest, acc *model.Accommodation) (float32, []*model.Availability, error) {
	dateFrom := request.DateFrom.AsTime()
	dateTo := request.DateTo.AsTime()
	var availabilitiesToUpdate []*model.Availability
	for date := dateFrom; !date.After(dateTo); date = date.Add(time.Hour * 24) {
		av, err := service.availabilityStore.GetAllAvailabilitiesForRevert(request.Accommodationid, date)
		if err != nil {
			println(err.Error())
			return 0, nil, err
		}
		availabilitiesToUpdate = append(availabilitiesToUpdate, av)
	}
	return 0, availabilitiesToUpdate, nil
}
