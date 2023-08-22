package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"reservation_service/domain/model"
	"time"
)

type ReservationMongoDBStore struct {
	reservations *mongo.Collection
}

func NewReservationMongoDBStore(client *mongo.Client) ReservationStore {
	reservations := client.Database(DATABASE).Collection(COLLECTION)
	return &ReservationMongoDBStore{
		reservations: reservations,
	}
}

func (store *ReservationMongoDBStore) Get(id primitive.ObjectID) (*model.Reservation, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *ReservationMongoDBStore) GetAllByUserId(id primitive.ObjectID) ([]*model.Reservation, error) {
	filter := bson.M{"user": id.Hex()}
	return store.filter(filter)
}

func (store *ReservationMongoDBStore) GetAll() ([]*model.Reservation, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *ReservationMongoDBStore) Insert(reservation *model.Reservation) (*model.Reservation, error) {
	result, err := store.reservations.InsertOne(context.TODO(), reservation)
	if err != nil {
		return nil, err
	}
	reservation.Id = result.InsertedID.(primitive.ObjectID)
	return reservation, nil
}

func (store *ReservationMongoDBStore) Delete(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": id}
	result, err := store.reservations.DeleteOne(context.TODO(), filter)
	return result, err
}

func (store *ReservationMongoDBStore) Update(reservation *model.Reservation) (*mongo.UpdateResult, error) {
	update := bson.D{{"$set",
		bson.D{
			{"accommodation", reservation.AccommodationId},
			{"start", reservation.Start},
			{"end", reservation.End},
			{"user", reservation.UserId},
			{"numberOfGuests", reservation.NumberOfGuests},
			{"status", reservation.Status},
			{"price", reservation.Price},
		},
	}}
	result, err := store.reservations.UpdateByID(context.TODO(), reservation.Id, update)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (store *ReservationMongoDBStore) GetActiveByUserId(id string) ([]*model.Reservation, error) {
	statuses := []string{"Accepted", "Pending"}
	today := time.Now()
	filter := bson.M{"user": id, "status": bson.M{"$in": statuses}, "end": bson.M{"$gt": today}}
	return store.filter(filter)
}
func (store *ReservationMongoDBStore) GetActiveForAccommodations(ids []string) ([]*model.Reservation, error) {
	statuses := []string{"Accepted", "Pending"}
	today := time.Now()
	filter := bson.M{"accommodation": bson.M{"$in": ids}, "status": bson.M{"$in": statuses}, "end": bson.M{"$gt": today}}
	return store.filter(filter)
}

func (store *ReservationMongoDBStore) GetPastByUserId(guestId, accommodationId string) ([]*model.Reservation, error) {
	today := time.Now()
	filter := bson.M{"user": guestId, "accommodation": accommodationId, "status": "Approved", "end": bson.M{"$lt": today}}
	return store.filter(filter)
}
func (store *ReservationMongoDBStore) GetPastForAccommodations(guestId string, ids []string) ([]*model.Reservation, error) {
	today := time.Now()
	filter := bson.M{"user": guestId, "accommodation": bson.M{"$in": ids}, "status": "Approved", "end": bson.M{"$lt": today}}
	return store.filter(filter)
}

func (store *ReservationMongoDBStore) GetAllIntercepting(reservation *model.Reservation) ([]*model.Reservation, error) {
	filter := bson.M{
		"accommodation": reservation.AccommodationId,
		"status":        "Pending",
	}
	return store.filter(filter)
}

func (store *ReservationMongoDBStore) filter(filter interface{}) ([]*model.Reservation, error) {
	cursor, err := store.reservations.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decodeReservations(cursor)
}

func (store *ReservationMongoDBStore) filterOne(filter interface{}) (reservation *model.Reservation, err error) {
	result := store.reservations.FindOne(context.TODO(), filter)
	err = result.Decode(&reservation)
	return
}

func decodeReservations(cursor *mongo.Cursor) (reservations []*model.Reservation, err error) {
	for cursor.Next(context.TODO()) {
		var reservation model.Reservation
		err = cursor.Decode(&reservation)
		if err != nil {
			return
		}
		reservations = append(reservations, &reservation)
	}
	err = cursor.Err()
	return
}

func (store *ReservationMongoDBStore) GetReservationsForAccommodationsByStatus(accommodationIds []string, status string) ([]*model.Reservation, error) {
	filter := bson.M{"accommodation": bson.M{"$in": accommodationIds}, "status": status}
	return store.filter(filter)
}
