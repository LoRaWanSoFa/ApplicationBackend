package components

import "log"

type MessageUplinkI interface {
	AddPayload(p []byte, s Sensor)
	GetPayloads() []messagePayloadI
	GetId() int64
}

type MessageUplink struct {
	Id       int64
	Time     string
	DevEUI   string // or [8]byte or types.DevEUI
	Payloads []messagePayloadI
}

func NewMessageUplink(id int64, devEUI string) MessageUplinkI {
	message := new(MessageUplink)
	message.Id = id
	message.DevEUI = devEUI
	message.Payloads = make([]messagePayloadI, 0)
	return message
}

func (m *MessageUplink) AddPayload(p []byte, s Sensor) {
	mp := new(messagePayloadByte)
	mp.Payload = p
	mp.Sensor = s
	log.Printf("Payload with payloads: %+v", mp)
	m.Payloads = append(m.Payloads, mp)
}

func (m *MessageUplink) GetPayloads() []messagePayloadI {
	return m.Payloads
}

func (m *MessageUplink) GetId() int64 {
	return m.Id
}
