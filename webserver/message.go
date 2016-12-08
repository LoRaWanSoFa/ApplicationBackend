package webserver

import "time"

type Message struct {
	Id     int       `json:"id"`
	Deveui string    `json:"deveui"`
	Down   bool      `json:"down"`
	Time   time.Time `json:"time"`
}

type Messages []Message
