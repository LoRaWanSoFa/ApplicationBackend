package mqttUplink

import (
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
	//	uplinkMessageHandler(client, appEUI, devEUI, req)
}
