package mqttUplink

import (
	"errors"
	"fmt"
	"log"

	components "github.com/LoRaWanSoFa/LoRaWanSoFa/Components"
	"github.com/LoRaWanSoFa/LoRaWanSoFa/Core/ByteConverter"
	DBC "github.com/LoRaWanSoFa/LoRaWanSoFa/DBC/DatabaseConnector"
)

type MessageCreator interface {
	CreateMessage(payload []byte, devEui []byte) (components.MessageUplinkI, error)
}

type messageCreator struct {
	byteConverter byteConverter.ByteConverter
}

// A messageCreator is created, the purpose of the MessageCreator is to
// convert the message, that comes in from a payload and the devEUI as bytes,
// to the MessageUplinkI format for further use.
func NewMessageCreator() MessageCreator {
	mc := new(messageCreator)
	mc.byteConverter = byteConverter.New()
	return mc
}

// Creates a MessageUplinkI object from the payload and devEui that were entered
// as bytes. If there is no header found for the input devEui this method will
// return an error.
func (m *messageCreator) CreateMessage(payload []byte, devEui []byte) (components.MessageUplinkI, error) {
	var message components.MessageUplinkI
	var sensors []components.Sensor
	// Convert devEui from bytes into a hexadecimal representation of them as a string.
	devEuiS, err := m.byteConverter.ConvertSingleValue(devEui, 4)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	sensors = DBC.GetNodeSensors(devEuiS)
	if len(sensors) > 0 {
		message, err = DBC.AddMessage(devEuiS)
	} else {
		err = errors.New(fmt.Sprintf("There was no header received for %s", devEuiS))
		return nil, err
	}

	// adding payloads to the newly created message
	b, headerLength := m.checkPayloadLength(payload[1:], sensors)
	if b {
		m.addPayloads(payload[1:], &message, sensors)
	} else {
		err = errors.New(fmt.Sprintf("The existing header for %s is not of the "+
			"right length for the received message. Header length was %d, while "+
			"payload length was %d.", devEuiS, headerLength, len(payload[1:])))
		return nil, err
	}

	return message, nil
}

func (m *messageCreator) addPayloads(payload []byte, message *components.MessageUplinkI, sensors []components.Sensor) {
	for i := range sensors {
		LoV := sensors[i].LenghtOfValues
		NoV := sensors[i].NumberOfValues
		for j := 0; j < NoV; j++ {
			(*message).AddPayload(payload[:LoV], sensors[i])
			payload = payload[LoV:]
		}
	}
}

func (m *messageCreator) checkPayloadLength(payload []byte, sensors []components.Sensor) (bool, int) {
	length := 0
	for i := range sensors {
		LoV := sensors[i].LenghtOfValues
		NoV := sensors[i].NumberOfValues
		length = length + (LoV * NoV)
	}
	return length == len(payload), length
}
