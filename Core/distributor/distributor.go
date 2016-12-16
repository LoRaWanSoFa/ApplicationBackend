package distributor

import (
	"errors"
	"log"

	components "github.com/LoRaWanSoFa/LoRaWanSoFa/Components"
	"github.com/LoRaWanSoFa/LoRaWanSoFa/Core/ByteConverter"
	"github.com/LoRaWanSoFa/LoRaWanSoFa/Core/restUplinkConnector"
	"github.com/LoRaWanSoFa/LoRaWanSoFa/DBC/DatabaseConnector"
)

var logFatal = log.Fatal

type Distributor interface {
	InputUplink(components.MessageUplinkI) (components.MessageUplinkI, error)
	InputDownlink(components.MessageDownLink)
}

type distributor struct {
	byteConverter       byteConverter.ByteConverter
	restUplinkConnector restUplink.RestUplinkConnector
}

// Creates a new Distributor object.
func New() Distributor {
	dist := new(distributor)
	dist.byteConverter = byteConverter.New()
	config := components.GetConfiguration().Rest
	dist.restUplinkConnector = restUplink.NewRestUplinkConnector(config.Ip, config.ApiKey)
	return dist
}

// Receives an Uplink message and distributes the message to the parts of the
// application that need to receive it.
func (d *distributor) InputUplink(message components.MessageUplinkI) (components.MessageUplinkI, error) {
	if d.deduplicate(message) {
		newMessage := d.convertMessage(message)
		err := DatabaseConnector.StoreMessagePayloads(newMessage)
		if err != nil {
			logFatal(err)
		}
		d.restUplinkConnector.NewData(newMessage.GetDevEUI(), newMessage)
		return newMessage, nil
	} else {
		err := errors.New("message was a duplicate")
		return nil, err
	}
}

// Receives a Downlink message and distributes the message to the parts of the
// application that need to receive it.
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

// Converts a message payload from bytes to string.
func (d *distributor) convertMessage(message components.MessageUplinkI) components.MessageUplinkI {
	bytePayloads := message.GetPayloads()
	message.RemovePayloads()
	for i := range bytePayloads {
		payload, ok := bytePayloads[i].GetPayload().([]byte)
		if ok {
			sensor := bytePayloads[i].GetSensor()
			payloadS, err := d.byteConverter.ConvertSingleValue(payload, sensor.DataType)
			if err != nil {
				logFatal(err)
			} else {
				message.AddPayloadString(payloadS, sensor)
			}
		}
	}
	return message
}
