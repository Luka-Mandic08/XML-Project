package repository

import (
	"auth_service/domain/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const (
	DATABASE   = "auth"
	COLLECTION = "auth"
)

type AuthMongoDBStore struct {
	accounts *mongo.Collection
}

func NewAuthMongoDBStore(client *mongo.Client) AuthStore {
	accounts := client.Database(DATABASE).Collection(COLLECTION)
	return &AuthMongoDBStore{
		accounts: accounts,
	}
}

func (store *AuthMongoDBStore) GetById(id primitive.ObjectID) (*model.Account, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *AuthMongoDBStore) Insert(account *model.Account) (*model.Account, error) {
	result, err := store.accounts.InsertOne(context.TODO(), account)
	if err != nil {
		return nil, err
	}
	account.Id = result.InsertedID.(primitive.ObjectID)
	return account, nil
}

func (store *AuthMongoDBStore) GetByUsername(username string) (*model.Account, error) {
	filter := bson.M{"username": username}
	return store.filterOne(filter)
}

func (store *AuthMongoDBStore) GetByUserId(userId string) (*model.Account, error) {
	filter := bson.M{"userid": userId}
	return store.filterOne(filter)
}

func (store *AuthMongoDBStore) Update(account *model.Account) (*mongo.UpdateResult, error) {
	filter := bson.M{"userid": account.UserId}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	update := bson.D{{"$set",
		bson.D{
			{"username", account.Username},
			{"password", string(hashedPassword)},
		},
	}}
	result, err := store.accounts.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (store *AuthMongoDBStore) Delete(id string) (*mongo.DeleteResult, error) {
	filter := bson.M{"userid": id}
	result, err := store.accounts.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (store *AuthMongoDBStore) filter(filter interface{}) ([]*model.Account, error) {
	cursor, err := store.accounts.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *AuthMongoDBStore) filterOne(filter interface{}) (account *model.Account, err error) {
	result := store.accounts.FindOne(context.TODO(), filter)
	err = result.Decode(&account)
	return
}

func decode(cursor *mongo.Cursor) (accounts []*model.Account, err error) {
	for cursor.Next(context.TODO()) {
		var account model.Account
		err = cursor.Decode(&account)
		if err != nil {
			return
		}
		accounts = append(accounts, &account)
	}
	err = cursor.Err()
	return
}
