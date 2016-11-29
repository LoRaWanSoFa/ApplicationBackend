package components

type MessageUplinkI interface {
	AddPayload(p []byte, s Sensor)
	GetPayloads() []messagePayload
}

type MessageUplink struct {
	Id       int64
	Time     string
	DevEUI   string // or [8]byte or types.DevEUI
	Payloads []messagePayload
}

func NewMessageUplink(id int64, devEUI string) MessageUplinkI {
	var messageI MessageUplinkI
	var message MessageUplink
	message = MessageUplink{Id: id, DevEUI: devEUI}
	message.Payloads = make([]messagePayload, 0)
	messageI = message
	return messageI
}

func (m MessageUplink) AddPayload(p []byte, s Sensor) {
	mp := messagePayload{payload: p, sensor: s}
	m.Payloads = append(m.Payloads, mp)
}

func (m MessageUplink) GetPayloads() []messagePayload {
	return m.Payloads
}

type messagePayload struct {
	id      int64
	payload []byte
	sensor  Sensor
}
