package model

import (
	"encoding/json"
	"io"
)

type Credentials struct {
	Username string `bson:"username,omitempty" json:"username"`
	Password string `bson:"password,omitempty" json:"password"`
}

func (c *Credentials) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}

func (c *Credentials) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(c)
}

/*
{
	"username": "galicc",
	"password": "hcijesranje"
}
*/
