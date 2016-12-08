package mqttUplink

import (
	"testing"

	components "github.com/LoRaWanSoFa/LoRaWanSoFa/Components"
)

var devEuiB = []byte{0x00, 0x00, 0x00, 0x00, 0xAB, 0xCD, 0xEF, 0x12}
var devEuiS = "00000000ABCDEF12"
var payload = []byte{0x42, 0x22, 0xEC, 0x25, 0xC2, 0x93, 0xDE, 0xD8, 0x01}
var mc = NewMessageCreator()
var message = mc.CreateMessage(payload, devEuiB)

func TestAddSimpleMessage(t *testing.T) {

	if message.GetDevEUI() != devEuiS {
		t.Errorf("Expected %s, was %s with input %v", devEuiS, message.GetDevEUI(), devEuiB)
	}
}

func TestCheckPayloads(t *testing.T) {

	gpsSensor := components.NewSensor(3, 0, 0, 2, 4, 1, 2, "", "0")
	boolSensor := components.NewSensor(4, 0, 0, 1, 1, 2, 5, "", "0")
	expectedMessage := components.NewMessageUplink(123, devEuiS)
	expectedMessage.AddPayload([]byte{0x42, 0x22, 0xEC, 0x25}, gpsSensor)
	expectedMessage.AddPayload([]byte{0xC2, 0x93, 0xDE, 0xD8}, gpsSensor)
	expectedMessage.AddPayload([]byte{0x01}, boolSensor)
	if len(message.GetPayloads()) != len(expectedMessage.GetPayloads()) {
		t.Errorf("There should have been %d elements in the payloads of message, but there were %d.", len(expectedMessage.GetPayloads()), len(message.GetPayloads()))
	}
	for i := range message.GetPayloads() {
		if !message.GetPayloads()[i].Equals(expectedMessage.GetPayloads()[i]) {
			t.Errorf("Message payloads should be equal, expected message %+v, but received message %+v",
				expectedMessage.GetPayloads()[i], message.GetPayloads()[i])
		}

	}
}
