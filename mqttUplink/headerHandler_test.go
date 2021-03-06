package mqttUplink

import (
	"testing"

	components "github.com/LoRaWanSoFa/ApplicationBackend/Components"
)

var h = NewHeaderHandler()
var headerPayload = []byte{0x10, 0x00, 0x4F, 0x08, 0x00, 0x65, 0x20, 0x0C, 0x4F, 0x08}

func TestCreateNewHeader(t *testing.T) {
	sensors, _ := h.CreateNewHeader(headerPayload, "")
	if len(expectedSensors) != len(sensors) {
		t.Errorf("Expected the header to have %d sensors, but had %d.",
			len(expectedSensors), len(sensors))
	}
	for i := range expectedSensors {
		if expectedSensors[i] != sensors[i] {
			t.Errorf("Sensors should be the same \n expected sensor: %+v \n "+
				"got sensor: %+v", expectedSensors[i], sensors[i])
		}
	}
}

func TestHeaderWithWrongLength(t *testing.T) {
	wrongLengthPayload := []byte{0x10, 0x00}
	_, err := h.CreateNewHeader(wrongLengthPayload, "")
	if err == nil {
		t.Errorf("Payload with %d, should generate an error in the header handler", len(wrongLengthPayload))
	}
}

func TestStorePayload(t *testing.T) {
	_, _, err := h.StoreHeader(expectedSensors, devEuiS)
	if err != nil {
		t.Errorf("The addition of %+v \n Should not fail, but got %+v.", expectedSensors, err)
	}

	differentSensor := components.Sensor{IoType: 1, IoAddress: 1, SensorType: 0, LenghtOfValues: 4, NumberOfValues: 2, HeaderOrder: 3}
	expectedSensors[2] = differentSensor
	_, _, err = h.StoreHeader(expectedSensors, devEuiS)
	if err != nil {
		t.Errorf("The addition of %+v \n Should not fail, but got %+v.", expectedSensors, err)
	}

}
