package model

import (
	"encoding/json"
	"io"
)

type LinkUserDTO struct {
	ApiKey   APIKey `bson:"apikey,omitempty" json:"apikey"`
	Username string `bson:"username,omitempty" json:"username"`
	Password string `bson:"password,omitempty" json:"password"`
}

func (dto *LinkUserDTO) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(dto)
}

func (dto *LinkUserDTO) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(dto)
}
