package model

import (
	"encoding/json"
	"io"
)

type UserAddress struct {
	Street  string `bson:"street,omitempty" json:"street"`
	City    string `bson:"city,omitempty" json:"city"`
	Country string `bson:"country,omitempty" json:"country"`
}

func (a *UserAddress) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(a)
}

func (a *UserAddress) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(a)
}

/*
{
	"street": "Tolstojeva",
	"city": "Novi Sad",
	"country": "Srbija"
}
*/
