package mqttUplink

import (
	"fmt"
	"log"

	components "github.com/LoRaWanSoFa/LoRaWanSoFa/Components"
	"github.com/LoRaWanSoFa/LoRaWanSoFa/Core/MessageConverter"
	DBC "github.com/LoRaWanSoFa/LoRaWanSoFa/DBC/DatabaseConnector"
)

type MessageCreator interface {
	CreateMessage(payload []byte, devEui []byte) components.MessageUplinkI
}

type messageCreator struct {
	messageConverter MessageConverter.MessageConverter
}

func NewMessageCreator() MessageCreator {
	mc := new(messageCreator)
	mc.messageConverter = MessageConverter.New()
	return mc
}

// Creates a MessageUplinkI object from the payload and devEui that were entered
// as bytes.
func (m *messageCreator) CreateMessage(payload []byte, devEui []byte) components.MessageUplinkI {
	var message components.MessageUplinkI
	var sensors []components.Sensor
	// Convert devEui from bytes into a hexadecimal representation of them as a string.
	devEuiS, err := m.messageConverter.ConvertSingleValue(devEui, 4)
	if err != nil {
		log.Fatal(err)
	}

	// Database entry creation.
	err = DBC.Connect()
	if err != nil {
		log.Fatal(err)
	} else {
		message, err = DBC.AddMessage(devEuiS)
		sensors = DBC.GetNodeSensors(devEuiS)
		err = DBC.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

	m.addPayloads(payload, &message, sensors)

	return message
}

func (m *messageCreator) addPayloads(payload []byte, message *components.MessageUplinkI, sensors []components.Sensor) {

	fmt.Println("enter add Payloads")
	for i := range sensors {
		LoV := sensors[i].LenghtOfValues
		NoV := sensors[i].NumberOfValues
		for j := 0; j < NoV; j++ {
			(*message).AddPayload(payload[:LoV], sensors[i])
			fmt.Println(payload)
			payload = payload[LoV:]
			fmt.Println((*message).GetPayloads())
		}
	}
}
