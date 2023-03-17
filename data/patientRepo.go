package data

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	// NoSQL: module containing Mongo api client
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// NoSQL: ProductRepo struct encapsulating Mongo api client
type PatientRepo struct {
	cli    *mongo.Client
	logger *log.Logger
}

// NoSQL: Constructor which reads db configuration from environment
func New(ctx context.Context, logger *log.Logger) (*PatientRepo, error) {
	dburi := os.Getenv("MONGO_DB_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &PatientRepo{
		cli:    client,
		logger: logger,
	}, nil
}

// Disconnect from database
func (pr *PatientRepo) Disconnect(ctx context.Context) error {
	err := pr.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Check database connection
func (pr *PatientRepo) Ping() {
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

func (pr *PatientRepo) GetAll() (Patients, error) {
	// Initialise context (after 5 seconds timeout, abort operation)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	patientsCollection := pr.getCollection()

	var patients Patients
	patientsCursor, err := patientsCollection.Find(ctx, bson.M{})
	if err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	if err = patientsCursor.All(ctx, &patients); err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	return patients, nil
}

func (pr *PatientRepo) GetById(id string) (*Patient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	patientsCollection := pr.getCollection()

	var patient Patient
	objID, _ := primitive.ObjectIDFromHex(id)
	err := patientsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&patient)
	if err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	return &patient, nil
}

func (pr *PatientRepo) GetByName(name string) (Patients, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	patientsCollection := pr.getCollection()

	var patients Patients
	patientsCursor, err := patientsCollection.Find(ctx, bson.M{"name": name})
	if err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	if err = patientsCursor.All(ctx, &patients); err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	return patients, nil
}

func (pr *PatientRepo) Insert(patient *Patient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	patientsCollection := pr.getCollection()

	result, err := patientsCollection.InsertOne(ctx, &patient)
	if err != nil {
		pr.logger.Println(err)
		return err
	}
	pr.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (pr *PatientRepo) Update(id string, patient *Patient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	patientsCollection := pr.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{
		"name":    patient.Name,
		"surname": patient.Surname,
	}}
	result, err := patientsCollection.UpdateOne(ctx, filter, update)
	pr.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	pr.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		pr.logger.Println(err)
		return err
	}
	return nil
}

func (pr *PatientRepo) AddPhoneNumber(id string, phoneNumber string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	patientsCollection := pr.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	update := bson.M{"$push": bson.M{
		"phoneNumbers": phoneNumber,
	}}
	result, err := patientsCollection.UpdateOne(ctx, filter, update)
	pr.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	pr.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		pr.logger.Println(err)
		return err
	}
	return nil
}

func (pr *PatientRepo) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	patientsCollection := pr.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	result, err := patientsCollection.DeleteOne(ctx, filter)
	if err != nil {
		pr.logger.Println(err)
		return err
	}
	pr.logger.Printf("Documents deleted: %v\n", result.DeletedCount)
	return nil
}

func (pr *PatientRepo) AddAnamnesis(id string, anamnesis *Anamnesis) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	patientsCollection := pr.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.M{"$push": bson.M{
		"anamnesis": anamnesis,
	}}
	result, err := patientsCollection.UpdateOne(ctx, filter, update)
	pr.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	pr.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		pr.logger.Println(err)
		return err
	}
	return nil
}

func (pr *PatientRepo) AddTherapy(id string, therapy *Therapy) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	patientsCollection := pr.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.M{"$push": bson.M{
		"therapy": therapy,
	}}
	result, err := patientsCollection.UpdateOne(ctx, filter, update)
	pr.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	pr.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		pr.logger.Println(err)
		return err
	}
	return nil
}

func (pr *PatientRepo) UpdateAddress(id string, address *Address) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	patientsCollection := pr.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.M{"$set": bson.M{
		"address": address,
	}}
	result, err := patientsCollection.UpdateOne(ctx, filter, update)
	pr.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	pr.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		pr.logger.Println(err)
		return err
	}
	return nil
}

func (pr *PatientRepo) ChangePhone(id string, index int, phoneNumber string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	patientsCollection := pr.getCollection()

	// What happens if set value for index=10, but we only have 3 phone numbers?
	// -> Every value in between will be set to an empty string
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.M{"$set": bson.M{
		"phoneNumbers." + strconv.Itoa(index): phoneNumber,
	}}
	result, err := patientsCollection.UpdateOne(ctx, filter, update)
	pr.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	pr.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		pr.logger.Println(err)
		return err
	}
	return nil
}

// BONUS
func (pr *PatientRepo) Receipt(id string) (float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	patientsCollection := pr.getCollection()

	objID, _ := primitive.ObjectIDFromHex(id)
	matchStage := bson.D{{"$match", bson.D{{"_id", objID}}}}
	sumStage := bson.D{{"$addFields",
		bson.D{{"total", bson.D{{"$add",
			bson.D{{"$sum", "$therapy.price"}},
		}},
		}},
	}}
	projectStage := bson.D{{"$project",
		bson.D{{"total", 1}},
	}}

	cursor, err := patientsCollection.Aggregate(ctx, mongo.Pipeline{matchStage, sumStage, projectStage})
	if err != nil {
		pr.logger.Println(err)
		return -1, err
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		pr.logger.Println(err)
		return -1, err
	}
	for _, result := range results {
		pr.logger.Println(result)
		return result["total"].(float64), nil
	}
	return -1, nil
}

func (pr *PatientRepo) Report() (map[string]float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	patientsCollection := pr.getCollection()

	sumStage := bson.D{{"$addFields",
		bson.D{{"total", bson.D{{"$add",
			bson.D{{"$sum", "$therapy.price"}},
		}},
		}},
	}}
	projectStage := bson.D{{"$project",
		bson.D{{"name", 1}, {"surname", 1}, {"total", 1}},
	}}

	cursor, err := patientsCollection.Aggregate(ctx, mongo.Pipeline{sumStage, projectStage})
	if err != nil {
		pr.logger.Println(err)
		return nil, err
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		pr.logger.Println(err)
		return nil, err
	}
	report := make(map[string]float64)
	for _, result := range results {
		pr.logger.Println(result)
		report[result["_id"].(primitive.ObjectID).Hex()] = result["total"].(float64)
	}
	return report, nil
}

func (pr *PatientRepo) getCollection() *mongo.Collection {
	patientDatabase := pr.cli.Database("mongoDemo")
	patientsCollection := patientDatabase.Collection("patients")
	return patientsCollection
}
