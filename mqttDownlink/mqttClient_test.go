package mqttDownlink

import (
	"os"
	"testing"

	"github.com/LoRaWanSoFa/ttn/core"
	"github.com/LoRaWanSoFa/ttn/core/types"
)

var downClient MqttClient

func TestMain(m *testing.M) {
	setUp()
	result := m.Run()
	tearDown()
	os.Exit(result)
}

func setUp() {
	downClient = New()
}

func tearDown() {
	downClient = nil
}

func TestDownlinkClient(t *testing.T) {
	err := downClient.Connect()
	if err == nil {
		t.Error("Should not be able to connect currently.")
	}
	downClient.PublishDownlink(types.AppEUI{}, types.DevEUI{}, core.DataDownAppReq{})
	downClient.Disconnect()
}
