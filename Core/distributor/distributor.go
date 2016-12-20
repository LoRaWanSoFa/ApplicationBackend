package distributor

import (
	"errors"
	"log"
	"strconv"

	components "github.com/LoRaWanSoFa/LoRaWanSoFa/Components"
	"github.com/LoRaWanSoFa/LoRaWanSoFa/Core/ByteConverter"
	"github.com/LoRaWanSoFa/LoRaWanSoFa/Core/mqttDownlink"
	"github.com/LoRaWanSoFa/LoRaWanSoFa/Core/restUplinkConnector"
	"github.com/LoRaWanSoFa/LoRaWanSoFa/DBC/DatabaseConnector"
)

var logFatal = log.Fatal

// Distributor is used to distribute messages through the system.
// When it receives a Uplink message, it transfers that to the
// restUplinkConnector and the DatabaseConnector.
// When receiving a Downlink message, it transfers that to the DatabaseConnector
// and the mqttDownlink Client.
// When the header comes in, it also sends messages to the restUplink that
// Sensors has been either removed in the node header.
type Distributor interface {
	InputUplink(components.MessageUplinkI) (components.MessageUplinkI, error)
	InputDownlink(components.MessageDownLink)
	InputNewSensors(sensors []components.Sensor, devEUI string)
	DeleteSensors(sensors []components.Sensor, devEUI string)
}

type distributor struct {
	byteConverter       byteConverter.ByteConverter
	restUplinkConnector restUplink.Connector
	mqttDownlink        mqttDownlink.MqttClient
}

//New Creates a new Distributor object.
func New() Distributor {
	dist := new(distributor)
	dist.byteConverter = byteConverter.New()
	config := components.GetConfiguration().Rest
	dist.restUplinkConnector = restUplink.NewRestUplinkConnector(config.IP, config.APIKey)
	dist.mqttDownlink = mqttDownlink.New()
	dist.mqttDownlink.Connect()
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
	}
	err := errors.New("message was a duplicate")
	return nil, err
}

// Receives a Downlink message and distributes the message to the parts of the
// application that need to receive it.
func (d *distributor) InputDownlink(message components.MessageDownLink) {
	//TODO: Convert the values that are received from the message here to the
	// 			types that are expected in the call to mqttDownlink.PublishDownlink.
	//Important here:
	// 	The message payload is a string as it is input, but the mqttDownlink
	// 	expects a req.DataDownAppReq which has a slice of bytes as payload. And the
	// 	node itself also can only accept slices of bytes.
	//	This means the string that was received has to be converted to another
	//	data type and then to a slice of bytes again.
}

// Adds sensor to the frontend of the aplication.
func (d *distributor) InputNewSensors(sensors []components.Sensor, devEUI string) {
	for i := range sensors {
		d.restUplinkConnector.NewSensor(devEUI, strconv.FormatInt(sensors[i].ID, 10))
	}
}

// Deletes a sensor from the frontend of the application.
func (d *distributor) DeleteSensors(sensors []components.Sensor, devEUI string) {
	for i := range sensors {
		if sensors[i].SoftDeleted == true {
			d.restUplinkConnector.DeleteSensor(devEUI, strconv.FormatInt(sensors[i].ID, 10))
		}
	}
}

//The deduplicate method should deduplicate messages that come in once from the
//TTN side of things(semi-private) as well as our own private backend and return
// true only if the message has not been received yet.
func (d *distributor) deduplicate(message components.MessageUplinkI) bool {
	// TODO: deduplicate messages that could come in checking with the database
	// 			 or createing a small cache for it.
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
