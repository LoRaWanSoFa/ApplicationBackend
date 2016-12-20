package mqttUplink

import (
	"strings"
	"testing"

	components "github.com/LoRaWanSoFa/LoRaWanSoFa/Components"
	"github.com/LoRaWanSoFa/LoRaWanSoFa/Core/distributor"
	"github.com/LoRaWanSoFa/ttn/core"
	"github.com/LoRaWanSoFa/ttn/core/types"
)

type mockDistributor struct {
	message components.MessageUplinkI
}

func newMockDistributor() distributor.Distributor {
	return new(mockDistributor)
}

var mClient = New()

func (m *mockDistributor) InputUplink(message components.MessageUplinkI) (components.MessageUplinkI, error) {
	m.message = message
	return message, nil
}

func (m *mockDistributor) InputDownlink(message components.MessageDownLink) {
}
func (m *mockDistributor) InputNewSensors(sensors []components.Sensor, devEUI string) {
}
func (m *mockDistributor) DeleteSensors(sensors []components.Sensor, devEUI string) {
}

func TestUplinkMessageHandler(t *testing.T) {
	dist = newMockDistributor()
	mClient.Connect()
	payload := core.DataUpAppReq{}
	payload.Payload = []byte{0x20, 0x42, 0x22, 0xEC,
		0x25, 0xC2, 0x93, 0xDE, 0xD8, 0x01}
	devEui := types.DevEUI{0x00, 0x00, 0x00, 0x00, 0xAB, 0xCD, 0xEF, 0x12}

	uplinkMessageHandler(mClient.GetClient(), types.AppEUI{}, devEui, payload)
	payloads := dist.(*mockDistributor).message.GetPayloads()
	if len(payloads) != 3 {
		t.Errorf("payload should contain 3 arguments, but had %d.", len(payloads))
	}

	payload.Payload[0] = 0x01
	uplinkMessageHandler(mClient.GetClient(), types.AppEUI{}, devEui, payload)
	if errs[0] != "[No valid flag]" {
		t.Errorf("Should have gotten an invalid flag error, but got %s.", errs[0])
	}

	payload.Payload[0] = 0x10
	uplinkMessageHandler(mClient.GetClient(), types.AppEUI{}, devEui, payload)
	if len(errs) != 1 {
		t.Errorf("Should not have gotten any aditional errors, but had: %+v.", errs[:1])
	}

	payload.Payload = payload.Payload[:8]
	uplinkMessageHandler(mClient.GetClient(), types.AppEUI{}, devEui, payload)
	if errs[1] != "[Header of unkown length was send.]" {
		t.Errorf("Should have gotten an unkown header length error, but got: %+v.", errs[1])
	}

	payload.Payload[0] = 0x20
	uplinkMessageHandler(mClient.GetClient(), types.AppEUI{}, devEui, payload)
	if !strings.Contains(errs[2], "not of the right length for the received message.") {
		t.Errorf("Should have header - payload length mismatch error, but got: %+v.", errs[2])
	}

	mClient.Disconnect()
}
