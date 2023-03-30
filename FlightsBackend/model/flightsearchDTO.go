package model

import (
	"encoding/json"
	"io"
	"time"
)

type FlightSearchDTO struct {
	StartDate        time.Time `bson:"startdate" json:"startdate" gorm:"not null"`
	Destination      string    `bson:"destination" json:"destination" gorm:"not null"`
	Start            string    `bson:"start" json:"start" gorm:"not null"`
	RemainingTickets int64     `bson:"remainingtickets" json:"remainingtickets" gorm:"not null"`
}

func (f *FlightSearchDTO) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(f)
}

func (f *FlightSearchDTO) FromJSON(r io.Reader) error {
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
