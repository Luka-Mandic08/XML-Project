package model

type UserFlight struct {
	FlightID    string `bson:"flightID,omitempty" json:"flightID"`
	TicketCount int64  `bson:"ticketCount,omitempty" json:"ticketCount"`
}

type UserFlights []*UserFlight
