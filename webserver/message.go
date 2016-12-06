package webserver

import "time"

type Message struct {
	Id     int       `json:"id"`
	Deveui int       `json:"deveui"`
	Down   bool      `json:"down"`
	Time   time.Time `json:"time"`
}

type Messages []Message
