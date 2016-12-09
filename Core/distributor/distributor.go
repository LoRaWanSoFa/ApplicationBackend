package distributor

import (
	"errors"
	"log"

	components "github.com/LoRaWanSoFa/LoRaWanSoFa/Components"
	"github.com/LoRaWanSoFa/LoRaWanSoFa/Core/MessageConverter"
	"github.com/LoRaWanSoFa/LoRaWanSoFa/Core/restUplinkConnector"
	"github.com/LoRaWanSoFa/LoRaWanSoFa/DBC/DatabaseConnector"
)

type Distributor interface {
	InputUplink(components.MessageUplinkI) (components.MessageUplinkI, error)
	InputDownlink(components.MessageDownLink)
}

type distributor struct {
	messageConverter    MessageConverter.MessageConverter
	restUplinkConnector restUplink.RestUplinkConnector
}

func New() Distributor {
	dist := new(distributor)
	dist.messageConverter = MessageConverter.New()
	config := components.GetConfiguration().Rest
	dist.restUplinkConnector = restUplink.NewRestUplinkConnector(config.Ip, config.ApiKey)
	return dist
}

func (d *distributor) InputUplink(message components.MessageUplinkI) (components.MessageUplinkI, error) {
	if d.deduplicate(message) {
		newMessage := d.convertMessage(message)
		err := DatabaseConnector.StoreMessagePayloads(newMessage)
		if err != nil {
			log.Fatal(err)
		}
		d.restUplinkConnector.NewData(newMessage.GetDevEUI(), newMessage)
		return newMessage, nil
	} else {
		err := errors.New("message was a duplicate")
		return nil, err
	}
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
