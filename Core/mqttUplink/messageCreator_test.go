package mqttUplink

import (
	"testing"

	components "github.com/LoRaWanSoFa/LoRaWanSoFa/Components"
)

var devEuiB = []byte{0x00, 0x00, 0x00, 0x00, 0xAB, 0xCD, 0xEF, 0x12}
var devEuiS = "00000000ABCDEF12"
var payload = []byte{0x42, 0x22, 0xEC, 0x25, 0xC2, 0x93, 0xDE, 0xD8, 0x01}
var mc = NewMessageCreator()

func TestAddSimpleMessage(t *testing.T) {
	message := mc.CreateMessage(payload, devEuiB)

	if message.GetDevEUI() != devEuiS {
		t.Errorf("Expected %s, was %s with input %v", devEuiS, message.GetDevEUI(), devEuiB)
	}
}

func TestCheckPayloads(t *testing.T) {
	gpsSensor := components.NewSensor(3, 1, 1, 2, 4, 1, 2, "GPS", "0")
	boolSensor := components.NewSensor(4, 1, 2, 1, 1, 2, 5, "BOOL", "0")
	expectedMessage := components.NewMessageUplink(123, devEuiS)
	expectedMessage.AddPayload([]byte{0x42, 0x22, 0xEC, 0x25, 0xC2, 0x93, 0xDE, 0xD8}, gpsSensor)
	expectedMessage.AddPayload([]byte{0x01}, boolSensor)

}
