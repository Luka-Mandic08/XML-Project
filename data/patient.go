package data

import (
	"encoding/json"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Patient struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name         string             `bson:"name" json:"name"`
	Surname      string             `bson:"surname,omitempty" json:"surname"`
	PhoneNumbers []string           `bson:"phoneNumbers,omitempty" json:"phoneNumbers"`
	Address      Address            `bson:"address,omitempty" json:"address"`
	Anamnesis    []Anamnesis        `bson:"anamnesis,omitempty" json:"anamnesis"`
	Therapy      []Therapy          `bson:"therapy,omitempty" json:"therapy"`
}

type Address struct {
	Street  string `bson:"street,omitempty" json:"street"`
	City    string `bson:"city,omitempty" json:"city"`
	Country string `bson:"country,omitempty" json:"country"`
}

type Anamnesis struct {
	Symptom   string    `bson:"symptom,omitempty" json:"symptom"`
	Intensity string    `bson:"intensity,omitempty" json:"intensity"`
	StartDate time.Time `bson:"startDate,omitempty" json:"startDate"`
}

type Therapy struct {
	Name  string  `bson:"name,omitempty" json:"name"`
	Price float32 `bson:"price,omitempty" json:"price"`
}

type Patients []*Patient

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

func (p *Patients) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Patient) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Patient) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *Anamnesis) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Anamnesis) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *Therapy) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Therapy) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *Address) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Address) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *Flight) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Flight) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}
