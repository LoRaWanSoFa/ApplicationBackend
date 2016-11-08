package components

type Message interface {
}

type message struct {
	payload []byte
	time    string
	sensors []Sensor
	devEUI  int // or [8]byte or types.DevEUI
}
