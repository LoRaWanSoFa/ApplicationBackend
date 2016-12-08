package webserver

import "time"

type Message struct {
	Id      int       `json:"id"`
	Deveui  string    `json:"deveui"`
	Payload string    `json:"payload"`
	Time    time.Time `json:"time"`
}

type Messages []Message
