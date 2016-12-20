package distributor

import (
	"os"
	"testing"

	components "github.com/LoRaWanSoFa/LoRaWanSoFa/Components"
	"github.com/LoRaWanSoFa/LoRaWanSoFa/DBC/DatabaseConnector"
)

var dist Distributor
var devEuiS = "00000000ABCDEF12"

func TestMain(m *testing.M) {
	setUp()
	result := m.Run()
	tearDown()
	os.Exit(result)
}

func setUp() {
	DatabaseConnector.Connect()
	dist = New()
}

func tearDown() {
	DatabaseConnector.Close()
}

func TestConvertMessage(t *testing.T) {
	gpsSensor := components.NewSensor(3, 0, 0, 0, 0, 2, 4, 1, 2, "", "0", false)
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
}

func TestInputDownlink(t *testing.T) {
	dist.InputDownlink(components.MessageDownLink{})
	//TODO: Mock the mqttDownlink client to test the method.
}

func TestInputNewSensors(t *testing.T) {
	dist.InputNewSensors([]components.Sensor{}, "")
	// TODO: Mock the restUplinkConnector to test the method.
}

func TestDeleteSensors(t *testing.T) {
	dist.DeleteSensors([]components.Sensor{}, "")
	// TODO: Mock the restUplinkConnector to test the method.
}
