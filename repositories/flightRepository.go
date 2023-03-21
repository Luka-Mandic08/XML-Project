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
	cli    *mongo.Client
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
		cli:    client,
		logger: logger,
	}, nil
}

func (fr *FlightRepository) Disconnect(ctx context.Context) error {
	err := fr.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (fr *FlightRepository) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check connection -> if no error, connection is established
	err := fr.cli.Ping(ctx, readpref.Primary())
	if err != nil {
		fr.logger.Println(err)
	}

	// Print available databases
	databases, err := fr.cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		fr.logger.Println(err)
	}
	fmt.Println(databases)
}

func (fr *FlightRepository) getCollection() *mongo.Collection {
	flightDatabase := fr.cli.Database("mongoDemo")
	flightCollection := flightDatabase.Collection("flights")
	return flightCollection
}

func (fr *FlightRepository) Insert(flight *model.Flight) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	flightsCollection := fr.getCollection()

	result, err := flightsCollection.InsertOne(ctx, &flight)
	if err != nil {
		fr.logger.Println(err)
		return err
	}
	fr.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (fr *FlightRepository) GetById(id string) (*model.Flight, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	flightsCollection := fr.getCollection()

	var flight model.Flight
	objID, _ := primitive.ObjectIDFromHex(id)
	err := flightsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&flight)

	if err != nil {
		fr.logger.Println(err)
		return nil, err
	}
	return &flight, nil
}

func (fr *FlightRepository) GetAll() (model.Flights, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	flightsCollection := fr.getCollection()

	var flights model.Flights
	flightsCursor, err := flightsCollection.Find(ctx, bson.M{})
	if err != nil {
		fr.logger.Println(err)
		return nil, err
	}
	if err = flightsCursor.All(ctx, &flights); err != nil {
		fr.logger.Println(err)
		return nil, err
	}
	return flights, nil
}

func (fr *FlightRepository) Update(id string, amount int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	flightsCollection := fr.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	update := bson.M{
		"$inc": bson.M{"remainingtickets": -amount}}

	fr.logger.Println(amount)

	result, err := flightsCollection.UpdateOne(ctx, filter, update)
	fr.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	fr.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		fr.logger.Println(err)
		return err
	}
	return nil
}

func (fr *FlightRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	flightsCollection := fr.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	result, err := flightsCollection.DeleteOne(ctx, filter)
	if err != nil {
		fr.logger.Println(err)
		return err
	}

	fr.logger.Printf("Documents (Flights) deleted: %v\n", result.DeletedCount)
	return nil
}
