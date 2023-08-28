package model

import (
	"encoding/json"
	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name,omitempty" json:"name"`
	Surname     string             `bson:"surname,omitempty" json:"surname"`
	PhoneNumber string             `bson:"phoneNumber,omitempty" json:"phoneNumber"`
	Address     UserAddress        `bson:"address,omitempty" json:"address"`
	Credentials UserCredentials    `bson:"credentials,omitempty" json:"credentials"`
	Role        UserRole           `bson:"role,omitempty" json:"role"`
	Flights     UserFlights        `bson:"flights" json:"flights"`
	APIKey      APIKey             `bson:"apikey" json:"apikey"`
}

type Users []*User

func (u *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *User) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(u)
}

/*
{
    "name": "Ivan",
    "surname": "Galic",
    "phoneNumber": "0123456789",
    "address":{
		"street": "Tolstojeva",
		"city": "Novi Sad",
		"country": "Srbija"
	},
	"credentials":{
		"username": "galicc",
		"password": "hcijesranje"
	},
	"role": "USER",
	"flights": [

	]
}
*/
