package model

import (
	"encoding/json"
	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Surname     string             `bson:"surname,omitempty" json:"surname"`
	PhoneNumber string             `bson:"phoneNumber,omitempty" json:"phoneNumber"`
	Address     Address            `bson:"address,omitempty" json:"address"`
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
	}
}
*/
