package model

import (
	"encoding/json"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Flight struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	StartDate        time.Time          `bson:"startdate" json:"startdate" gorm:"not null"`
	ArrivalDate      time.Time          `bson:"arrivaldate" json:"arrivaldate" gorm:"not null"`
	Destination      string             `bson:"destination" json:"destination" gorm:"not null"`
	Start            string             `bson:"start" json:"start" gorm:"not null"`
	Price            float32            `bson:"price" json:"price" gorm:"not null"`
	RemainingTickets int                `bson:"remainingtickets" json:"remainingtickets" gorm:"not null"`
	TotalTickets     int                `bson:"totaltickets" json:"totaltickets" gorm:"not null"`
}

type Flights []*Flight

func (f *Flights) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(f)
}

func (f *Flight) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(f)
}

func (f *Flight) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(f)
}

/*
{
    "startdate": "2023-05-05T13:00:00Z",
    "arrivaldate": "2023-05-05T19:00:00Z",
    "destination": "Beograd",
    "start": "Berlin",
    "price": 200,
    "remainingtickets": 20,
    "totaltickets": 100
}
*/
