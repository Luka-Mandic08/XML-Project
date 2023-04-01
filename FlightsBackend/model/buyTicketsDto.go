package model

import (
	"encoding/json"
	"io"
)

type BuyTicketDto struct {
	FlightId string `json:"flightId"`
	Amount   int64  `json:"amount"`
	UserId   string `json:"userId"`
}

func (ticket *BuyTicketDto) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ticket)
}

func (ticket *BuyTicketDto) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(ticket)
}
