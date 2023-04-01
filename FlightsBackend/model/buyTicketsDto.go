package model

import (
	"encoding/json"
	"io"
)

type BuyTicketDto struct {
	UserId   string `json:"userId"`
	FlightId string `json:"flightId"`
	Amount   int64  `json:"amount"`
}

func (ticket *BuyTicketDto) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ticket)
}

func (ticket *BuyTicketDto) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(ticket)
}
