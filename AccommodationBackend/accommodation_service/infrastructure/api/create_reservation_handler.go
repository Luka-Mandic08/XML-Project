package api

import (
	"accommodation_service/domain/service"
	accommodation "common/proto/accommodation_service"
	events "common/saga/create_reservation"
	saga "common/saga/messaging"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CreateReservationCommandHandler struct {
	accommodationService *service.AccommodationService
	replyPublisher       saga.Publisher
	commandSubscriber    saga.Subscriber
}

func NewCreateReservationCommandHandler(accommodationService *service.AccommodationService, publisher saga.Publisher, subscriber saga.Subscriber) (*CreateReservationCommandHandler, error) {
	o := &CreateReservationCommandHandler{
		accommodationService: accommodationService,
		replyPublisher:       publisher,
		commandSubscriber:    subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *CreateReservationCommandHandler) handle(command *events.CreateReservationCommand) {
	reply := events.CreateReservationReply{Reservation: command.Reservation}

	accommodationId, err := primitive.ObjectIDFromHex(command.Reservation.AccommodationId)
	if err != nil {
		reply.Type = events.AccommodationNotExist
		println("Reply: events.AccommodationNotExist")
		println(err.Error())
		_ = handler.replyPublisher.Publish(reply)
		return
	}

	id, err := primitive.ObjectIDFromHex(command.Reservation.Id)
	if err != nil {
		reply.Type = events.AccommodationNotExist
		println("Reply: events.AccommodationNotExist")
		println(err.Error())
		_ = handler.replyPublisher.Publish(reply)
		return
	}

	switch command.Type {
	case events.CheckAccommodationExists:
		err := handler.accommodationService.CheckAccommodationExists(accommodationId)
		println("Command: events.CheckAccommodationExists")
		if err != nil {
			reply.Type = events.AccommodationNotExist
			println("Reply: events.AccommodationNotExist")
			println(err.Error() + "Id: " + id.Hex())
			break
		}
		reply.Type = events.AccommodationExists
		println("Reply: events.AccommodationExists")
	case events.CheckAvailableAccommodation:
		println("Command: events.CheckAvailableAccommodation")
		request := accommodation.CheckAvailabilityRequest{
			Accommodationid: command.Reservation.AccommodationId,
			DateFrom:        timestamppb.New(command.Reservation.Start),
			DateTo:          timestamppb.New(command.Reservation.End),
			NumberOfGuests:  command.Reservation.NumberOfGuests,
		}
		accommodation, err := handler.accommodationService.CheckAccommodationAvailability(&request)
		if err != nil {
			reply.Type = events.AccommodationNotAvailable
			println("Reply: events.AccommodationNotAvailable")
			println(err.Error())
			break
		}
		reply.Reservation.Price = accommodation.TotalPrice

		reply.Type = events.AccommodationAvailable
		println("Reply: events.AccommodationAvailable")
	case events.ChangeAvailability:
		println("Command: events.ChangeAvailability")
		availabilitiesToUpdate := handler.accommodationService.GetAllAvailability(command.Reservation.Start, command.Reservation.End, accommodationId.Hex())
		id, err = primitive.ObjectIDFromHex(command.Reservation.AccommodationId)
		accommodation, err := handler.accommodationService.Get(id)
		if accommodation.HasAutomaticReservations == true {
			err = handler.accommodationService.ChangeAvailability(availabilitiesToUpdate, false)
			if err != nil {
				reply.Type = events.AvailabilityNotChanged
				println("Reply: events.AvailabilityNotChanged")
				println(err.Error())
				break
			}
		}
		reply.Type = events.AvailabilityChanged
		println("Reply: events.AvailabilityChanged")
	case events.CheckAutomaticApproveReservation:
		println("Command: events.CheckAutomaticApproveReservation")
		accommodation, err := handler.accommodationService.Get(accommodationId)
		if err != nil {
			println(err.Error())
			break
		}
		if accommodation.HasAutomaticReservations {
			reply.Type = events.AutoApproveReservation
			println("Reply: events.AutoApproveReservation")
			break
		}
		reply.Type = events.ManualPendingReservation
		println("Reply: events.ManualPendingReservation")
	case events.RevertAvailability:
		println("Command: events.RevertAvailability")
		accommodationAvailability := accommodation.CheckAvailabilityRequest{
			Accommodationid: command.Reservation.Id,
			DateFrom:        timestamppb.New(command.Reservation.Start),
			DateTo:          timestamppb.New(command.Reservation.End),
			NumberOfGuests:  0,
		}

		id, err := primitive.ObjectIDFromHex(command.Reservation.Id)
		if err != nil {
			reply.Type = events.AvailabilityNotReverted
			println("Reply: events.AvailabilityNotReverted")
			break
		}
		accommodation, err := handler.accommodationService.Get(id)
		if err != nil {
			reply.Type = events.AvailabilityNotReverted
			println("Reply: events.AvailabilityNotReverted")
			break
		}

		_, availability, err := handler.accommodationService.CheckDateAvailability(&accommodationAvailability, accommodation)
		if err != nil {
			reply.Type = events.AvailabilityNotReverted
			println("Reply: events.AvailabilityNotReverted")
			break
		}

		err = handler.accommodationService.ChangeAvailability(availability, false)
		if err != nil {
			reply.Type = events.AvailabilityNotReverted
			println("Reply: events.AvailabilityNotReverted")
			break
		}
		reply.Type = events.AvailabilityReverted
		println("Reply: events.AvailabilityReverted")
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
