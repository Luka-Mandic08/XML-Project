package model

import (
	"encoding/json"
	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Role UserRole           `bson:"role,omitempty" json:"role"`
}

func (u *Session) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *Session) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(u)
}
