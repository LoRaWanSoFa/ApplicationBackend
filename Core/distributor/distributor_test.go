package distributor

import (
	"testing"

	components "github.com/LoRaWanSoFa/LoRaWanSoFa/Components"
	"github.com/LoRaWanSoFa/LoRaWanSoFa/DBC/DatabaseConnector"
)

var dist = New()
var devEuiS = "00000000ABCDEF12"

func TestConvertMessage(t *testing.T) {
	DatabaseConnector.Connect()
	gpsSensor := components.NewSensor(3, 0, 0, 0, 2, 4, 1, 2, "", "0")
	inputMessage := components.NewMessageUplink(123, devEuiS)
	inputMessage.AddPayload([]byte{0x42, 0x22, 0xEC, 0x25}, gpsSensor)
	inputMessage.AddPayload([]byte{0xC2, 0x93, 0xDE, 0xD8}, gpsSensor)
	expectedMessage := components.NewMessageUplink(123, devEuiS)
	expectedMessage.AddPayloadString("40.730610", gpsSensor)
	expectedMessage.AddPayloadString("-73.935242", gpsSensor)
	mp, _ := dist.InputUplink(inputMessage)
	payloads := mp.GetPayloads()
	for i := range payloads {
		inputPayload := payloads[i]
		expectedPayload := expectedMessage.GetPayloads()[i]
		if !inputPayload.Equals(expectedPayload) {
			t.Errorf("The payload of the message should be %s, but was %s.",
				expectedPayload.GetPayload(), inputPayload.GetPayload())
		}
	}
	DatabaseConnector.Close()
}
