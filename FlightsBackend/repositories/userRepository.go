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

type UserRepository struct {
	cli    *mongo.Client
	logger *log.Logger
}

// NoSQL: Constructor which reads db configuration from environment
func NewUserRepository(ctx context.Context, logger *log.Logger) (*UserRepository, error) {
	dburi := os.Getenv("MONGO_DB_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &UserRepository{
		cli:    client,
		logger: logger,
	}, nil
}

// Disconnect from database
func (pr *UserRepository) Disconnect(ctx context.Context) error {
	err := pr.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Check database connection
func (pr *UserRepository) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check connection -> if no error, connection is established
	err := pr.cli.Ping(ctx, readpref.Primary())
	if err != nil {
		pr.logger.Println(err)
	}

	// Print available databases
	databases, err := pr.cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		pr.logger.Println(err)
	}
	fmt.Println(databases)
}

// CRUD -- CREATE
func (pr *UserRepository) Insert(user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	usersCollection := pr.getCollection()

	result, err := usersCollection.InsertOne(ctx, &user)
	if err != nil {
		pr.logger.Println(err)
		return err
	}
	pr.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

// CRUD -- READ
func (pr *UserRepository) GetAll() (model.Users, error) {
	// Initialise context (after 5 seconds timeout, abort operation)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	usersCollection := pr.getCollection()

	var users model.Users
	usersCursor, err := usersCollection.Find(ctx, bson.M{})
	if err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	if err = usersCursor.All(ctx, &users); err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	return users, nil
}

func (pr *UserRepository) GetById(id string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	usersCollection := pr.getCollection()

	var user model.User
	objID, _ := primitive.ObjectIDFromHex(id)
	err := usersCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	return &user, nil
}

func (pr *UserRepository) GetByName(name string) (model.Users, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	usersCollection := pr.getCollection()

	var users model.Users
	patientsCursor, err := usersCollection.Find(ctx, bson.M{"name": name})
	if err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	if err = patientsCursor.All(ctx, &users); err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	return users, nil
}

// CRUD -- UPDATE
func (pr *UserRepository) Update(id string, user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	usersCollection := pr.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{
		"name":    user.Name,
		"surname": user.Surname,
	}}
	result, err := usersCollection.UpdateOne(ctx, filter, update)
	pr.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	pr.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		pr.logger.Println(err)
		return err
	}
	return nil
}

func (pr *UserRepository) UpdateAddress(id string, address *model.Address) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	usersCollection := pr.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.M{"$set": bson.M{
		"address": address,
	}}
	result, err := usersCollection.UpdateOne(ctx, filter, update)
	pr.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	pr.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		pr.logger.Println(err)
		return err
	}
	return nil
}

func (pr *UserRepository) ChangePhone(id string, index int, phoneNumber string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	usersCollection := pr.getCollection()

	// What happens if set value for index=10, but we only have 3 phone numbers?
	// -> Every value in between will be set to an empty string
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.M{"$set": bson.M{
		"phoneNumber": phoneNumber,
	}}
	result, err := usersCollection.UpdateOne(ctx, filter, update)
	pr.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	pr.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		pr.logger.Println(err)
		return err
	}
	return nil
}

// CRUD -- DELETE
func (pr *UserRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	usersCollection := pr.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	result, err := usersCollection.DeleteOne(ctx, filter)
	if err != nil {
		pr.logger.Println(err)
		return err
	}
	pr.logger.Printf("Documents deleted: %v\n", result.DeletedCount)
	return nil
}

// BONUS
func (pr *UserRepository) getCollection() *mongo.Collection {
	patientDatabase := pr.cli.Database("mongoDemo")
	patientsCollection := patientDatabase.Collection("users")
	return patientsCollection
}
