package distributor

import (
	"log"

	components "github.com/LoRaWanSoFa/LoRaWanSoFa/Components"
	"github.com/LoRaWanSoFa/LoRaWanSoFa/Core/MessageConverter"
)

type Distributor interface {
	InputUplink(components.MessageUplinkI) components.MessageUplinkI
	InputDownlink(components.MessageDownLink)
}

type distributor struct {
	messageConverter MessageConverter.MessageConverter
}

func New() Distributor {
	dist := new(distributor)
	dist.messageConverter = MessageConverter.New()
	return dist
}

func (d *distributor) InputUplink(message components.MessageUplinkI) components.MessageUplinkI {
	return d.convertMessage(message)
}
func (d *distributor) InputDownlink(message components.MessageDownLink) {

}

//The deduplicate method should deduplicate messages that come in once from the
//TTN side of things(semi-private) as well as our own private backend and return
// true only if the message has not been received yet.
func (d *distributor) deduplicate(message components.MessageUplinkI) bool {
	// TODO: deduplicate messages that could come in checking with the database
	// or createing a small cache for it.
	return true
}

func (d *distributor) convertMessage(message components.MessageUplinkI) components.MessageUplinkI {
	bytePayloads := message.GetPayloads()
	message.RemovePayloads()
	for i := range bytePayloads {
		payload, ok := bytePayloads[i].GetPayload().([]byte)
		if ok {
			sensor := bytePayloads[i].GetSensor()
			payloadS, err := d.messageConverter.ConvertSingleValue(payload, sensor.DataType)
			if err != nil {
				log.Fatal(err)
			} else {
				message.AddPayloadString(payloadS, sensor)
			}
		}
	}
	return message
}
