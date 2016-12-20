package components

import (
	"net/url"
	"strconv"
	"time"
)

// MessageUplinkI is used to handle uplink messages more efficiently.
// It gives methods that help handling and displaying the message.
type MessageUplinkI interface {
	AddPayload(p []byte, s Sensor)
	AddPayloadString(str string, s Sensor)
	GetPayloads() []messagePayloadI
	GetID() int64
	GetDevEUI() string
	RemovePayloads()
	ToJSON() url.Values
}

type messageUplink struct {
	ID       int64 //database id
	Time     string
	DevEUI   string // or [8]byte or types.DevEUI
	Payloads []messagePayloadI
}

// MessageDownLink is used for Downlink messages and is created by the webserver
// via restful JSON.
type MessageDownLink struct {
	ID      int64     `json:"id"`
	Deveui  string    `json:"deveui"`
	Payload string    `json:"payload"`
	Time    time.Time `json:"time"`
}

// Messages is a Slice of messages of type MessageDownLink, is used to bundle
// Downlink Messages in webserver.
type Messages []MessageDownLink

// NewMessageUplink is the constructor for the MessageUplinkI interface.
func NewMessageUplink(id int64, devEUI string) MessageUplinkI {
	message := new(messageUplink)
	message.ID = id
	message.DevEUI = devEUI
	message.Payloads = make([]messagePayloadI, 0)
	return message
}

func (m *messageUplink) AddPayload(p []byte, s Sensor) {
	mp := new(messagePayloadByte)
	mp.Payload = p
	mp.Sensor = s
	m.Payloads = append(m.Payloads, mp)
}

func (m *messageUplink) AddPayloadString(str string, s Sensor) {
	mp := new(messagePayloadString)
	mp.Payload = str
	mp.Sensor = s
	m.Payloads = append(m.Payloads, mp)
}

func (m *messageUplink) RemovePayloads() {
	m.Payloads = make([]messagePayloadI, 0)
}

func (m *messageUplink) GetPayloads() []messagePayloadI {
	return m.Payloads
}

func (m *messageUplink) GetID() int64 {
	return m.ID
}

func (m *messageUplink) GetDevEUI() string {
	return m.DevEUI
}

// ToJSON creates a json object from the MessageUplinkI Interface.
func (m *messageUplink) ToJSON() url.Values {
	json := url.Values{}
	payloads := m.GetPayloads()
	for i := range payloads {
		idString := strconv.FormatInt(payloads[i].GetSensor().ID, 10)
		json.Add(idString, payloads[i].GetPayload().(string))
	}
	return json
}
