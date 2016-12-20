package mqttUplink

import (
	"fmt"
	"os"
	"testing"

	components "github.com/LoRaWanSoFa/LoRaWanSoFa/Components"
	"github.com/LoRaWanSoFa/LoRaWanSoFa/DBC/DatabaseConnector"
)

var devEuiB = []byte{0x00, 0x00, 0x00, 0x00, 0xAB, 0xCD, 0xEF, 0x12}
var devEuiS = "00000000ABCDEF12"
var payload = []byte{0x20, 0x42, 0x22, 0xEC, 0x25, 0xC2, 0x93, 0xDE, 0xD8, 0x01}
var message components.MessageUplinkI
var mc = NewMessageCreator()
var expectedSensors []components.Sensor
var errs []string

func TestMain(m *testing.M) {
	result := 0
	err := DatabaseConnector.Connect()
	if err == nil {
		setUp()
		message, _ = mc.CreateMessage(payload, devEuiS)
		result = m.Run()
		DatabaseConnector.Close()
	}
	os.Exit(result)
}

func setUp() {
	errs = []string{}
	logFatal = func(args ...interface{}) {
		errs = append(errs, fmt.Sprintf("%+v", args))
	}
	expectedSensors = []components.Sensor{
		components.Sensor{IoType: 0, IoAddress: 0, SensorType: 79, LenghtOfValues: 1, NumberOfValues: 1, HeaderOrder: 1},
		components.Sensor{IoType: 0, IoAddress: 0, SensorType: 101, LenghtOfValues: 1, NumberOfValues: 4, HeaderOrder: 2},
		components.Sensor{IoType: 0, IoAddress: 3, SensorType: 79, LenghtOfValues: 1, NumberOfValues: 1, HeaderOrder: 3},
	}
}

func TestAddSimpleMessage(t *testing.T) {
	if message.GetDevEUI() != devEuiS {
		t.Errorf("Expected %s, was %s with input %v", devEuiS, message.GetDevEUI(), devEuiB)
	}
}

func TestCheckPayloads(t *testing.T) {
	gpsSensor := components.NewSensor(3, 0, 0, 0, 0, 2, 4, 1, 2, "", "0", false)
	boolSensor := components.NewSensor(4, 0, 0, 0, 0, 1, 1, 2, 5, "", "0", false)
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

func TestNoHeader(t *testing.T) {
	devEuiNoSensor := "00000000AF1294E5"
	_, err := mc.CreateMessage(payload, devEuiNoSensor)
	if err == nil {
		t.Errorf("The node should not have a header, and throw an error %+v", err)
	}
}
