package repositories

import (
	"Rest/model"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type FlightRepository struct {
	client *mongo.Client
	logger *log.Logger
}

func NewFlightRepository(ctx context.Context, logger *log.Logger) (*FlightRepository, error) {
	dburi := os.Getenv("MONGO_DB_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &FlightRepository{
		client: client,
		logger: logger,
	}, nil
}

func (flightRepository *FlightRepository) Disconnect(ctx context.Context) error {
	err := flightRepository.client.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (flightRepository *FlightRepository) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check connection -> if no error, connection is established
	err := flightRepository.client.Ping(ctx, readpref.Primary())
	if err != nil {
		flightRepository.logger.Println(err)
	}

	// Print available databases
	databases, err := flightRepository.client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		flightRepository.logger.Println(err)
	}
	fmt.Println(databases)
}

func (flightRepository *FlightRepository) getCollection() *mongo.Collection {
	flightDatabase := flightRepository.client.Database("mongoDemo")
	flightCollection := flightDatabase.Collection("flights")
	return flightCollection
}

func (flightRepository *FlightRepository) Insert(flight *model.Flight) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	flightsCollection := flightRepository.getCollection()

	result, err := flightsCollection.InsertOne(ctx, &flight)
	if err != nil {
		flightRepository.logger.Println(err)
		return err
	}
	flightRepository.logger.Printf("Flight document with ID: %v inserted.\n", result.InsertedID)
	return nil
}

func (flightRepository *FlightRepository) GetById(id string) (*model.Flight, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	flightsCollection := flightRepository.getCollection()

	var flight model.Flight
	objID, _ := primitive.ObjectIDFromHex(id)
	err := flightsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&flight)

	if err != nil {
		flightRepository.logger.Println(err)
		return nil, err
	}

	flightRepository.logger.Printf("Flight document with ID: %v found.\n", flight.ID)
	return &flight, nil
}

func (flightRepository *FlightRepository) GetByUser(userFlights model.UserFlights) (model.Flights, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	flightsCollection := flightRepository.getCollection()
	var ids []string
	for i, s := range userFlights {
		ids[i] = s.FlightID
	}

	var flights model.Flights

	flightsCursor, err := flightsCollection.Find(ctx, bson.M{"_id": bson.M{"$in": ids}})
	if err != nil {
		flightRepository.logger.Println(err)
		return nil, err
	}
	if err = flightsCursor.All(ctx, &flights); err != nil {
		flightRepository.logger.Println(err)
		return nil, err
	}

	for _, s := range flights {
		for _, q := range userFlights {
			if s.ID.String() == q.FlightID {
				s.RemainingTickets = q.TicketCount
			}
		}
	}

	return flights, nil
}

func (flightRepository *FlightRepository) GetAll() (model.Flights, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	flightsCollection := flightRepository.getCollection()

	var flights model.Flights
	flightsCursor, err := flightsCollection.Find(ctx, bson.M{})
	if err != nil {
		flightRepository.logger.Println(err)
		return nil, err
	}
	if err = flightsCursor.All(ctx, &flights); err != nil {
		flightRepository.logger.Println(err)
		return nil, err
	}

	return flights, nil
}

func (flightRepository *FlightRepository) GetSearched(dto *model.FlightSearchDTO) (model.Flights, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	flightsCollection := flightRepository.getCollection()
	helperDate := dto.StartDate.Add(time.Hour * 24)

	var flights model.Flights
	var flightsCursor *mongo.Cursor
	var err error

	if dto.StartDate.Year() != 1970 {
		flightsCursor, err = flightsCollection.Find(ctx, bson.M{"start": bson.M{"$regex": dto.Start, "$options": "i"},
			"destination":      bson.M{"$regex": dto.Destination, "$options": "i"},
			"startdate":        bson.M{"$gte": dto.StartDate, "$lt": helperDate},
			"remainingtickets": bson.M{"$gte": dto.RemainingTickets}})
	} else {
		flightsCursor, err = flightsCollection.Find(ctx, bson.M{"start": bson.M{"$regex": dto.Start, "$options": "i"},
			"destination":      bson.M{"$regex": dto.Destination, "$options": "i"},
			"remainingtickets": bson.M{"$gte": dto.RemainingTickets}})
	}
	if err != nil {
		flightRepository.logger.Println(err)
		return nil, err
	}
	if err = flightsCursor.All(ctx, &flights); err != nil {
		flightRepository.logger.Println(err)
		return nil, err
	}

	return flights, nil
}

func (flightRepository *FlightRepository) UpdateFlightRemainingTickets(id string, amount int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	flightsCollection := flightRepository.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	update := bson.M{
		"$inc": bson.M{"remainingtickets": -amount}}

	var flight model.Flight
	err := flightsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&flight)

	if err != nil {
		flightRepository.logger.Println(err)
		return err
	}

	flightRepository.logger.Printf("RT: %d\n", flight.RemainingTickets)
	flightRepository.logger.Printf("Amount: %d\n", amount)

	if flight.RemainingTickets >= amount {
		result, err2 := flightsCollection.UpdateOne(ctx, filter, update)
		flightRepository.logger.Printf("Number of Flight documents matched: %v\n", result.MatchedCount)
		flightRepository.logger.Printf("%d Flight document with id: %v updated.\n", result.ModifiedCount, objID)

		if err2 != nil {
			flightRepository.logger.Println(err2)
			return err2
		}
		return nil
	}

	flightRepository.logger.Println("Not enought tickets left.")
	return nil
}

func (flightRepository *FlightRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	flightsCollection := flightRepository.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	result, err := flightsCollection.DeleteOne(ctx, filter)
	if err != nil {
		flightRepository.logger.Println(err)
		return err
	}

	flightRepository.logger.Printf("%v Flight document with ID: %s deleted.\n", result.DeletedCount, id)
	return nil
}
