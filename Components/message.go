package components

type MessageUplinkI interface {
	AddPayload(p []byte, s Sensor)
	GetPayloads() []messagePayloadI
	GetId() int64
	GetDevEUI() string
	RemovePayloads()
	AddPayloadString(str string, s Sensor)
}

type MessageUplink struct {
	Id       int64 //database id
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
	m.Payloads = append(m.Payloads, mp)
}

func (m *MessageUplink) AddPayloadString(str string, s Sensor) {
	mp := new(messagePayloadString)
	mp.Payload = str
	mp.Sensor = s
	m.Payloads = append(m.Payloads, mp)
}

func (m *MessageUplink) RemovePayloads() {
	m.Payloads = make([]messagePayloadI, 0)
}

func (m *MessageUplink) GetPayloads() []messagePayloadI {
	return m.Payloads
}

func (m *MessageUplink) GetId() int64 {
	return m.Id
}

func (m *MessageUplink) GetDevEUI() string {
	return m.DevEUI
}
