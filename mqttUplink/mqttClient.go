package mqttUplink

import (
	"errors"
	"log"
	"os"

	"github.com/LoRaWanSoFa/ApplicationBackend/Components"
	"github.com/LoRaWanSoFa/ApplicationBackend/distributor"
	"github.com/LoRaWanSoFa/ttn/core"
	"github.com/LoRaWanSoFa/ttn/core/types"
	"github.com/LoRaWanSoFa/ttn/mqtt"
	apexLog "github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

var logFatal = log.Fatal

// MqttClient is the client that will receive message from the
// network backend.
type MqttClient interface {
	Connect() error
	Disconnect()
	GetClient() mqtt.Client
}

type mqttClient struct {
	client mqtt.Client
}

// New is the constructor for the Mqtt-Uplink-Client.
func New() MqttClient {
	client := new(mqttClient)
	return client
}

var hHandler = NewHeaderHandler()
var mCreator = NewMessageCreator()
var dist = distributor.New()

// Handles uplink messages received from the backend. Depending on the flag send
// with the payload (First byte), either a header is created / changed or a
// message is added to the database.
func uplinkMessageHandler(client mqtt.Client, appEUI types.AppEUI, devEUI types.DevEUI, req core.DataUpAppReq) {
	if len(req.Payload) > 0 {
		flag := req.Payload[0]
		if flag>>4 == 1 {
			header, err := hHandler.CreateNewHeader(req.Payload, devEUI.GoString())
			if err != nil {
				logFatal(err)
			} else {
				newSensors, activationChanged, err := hHandler.StoreHeader(header, devEUI.GoString())
				if err != nil {
					logFatal(err)
				} else {
					dist.InputNewSensors(newSensors, devEUI.GoString())
					dist.DeleteSensors(activationChanged, devEUI.GoString())
				}
			}
		} else if flag>>5 == 1 {
			message, err := mCreator.CreateMessage(req.Payload, devEUI.GoString())
			if err != nil {
				logFatal(err)
			} else {
				dist.InputUplink(message)
			}
		} else {
			err := errors.New("No valid flag")
			logFatal(err)
		}
	}
}

// Method that connects the mqtt client to the backend, received messages are
// send to the uplinkMessageHandler.
func (m *mqttClient) Connect() error {
	apexLog.SetHandler(text.New(os.Stderr))
	mqttConfig := components.GetConfiguration().Mqtt
	ctx := apexLog.WithField("distributorClient", "mqtt-distributorClient")
	m.client = mqtt.NewClient(ctx, "ttnctl", mqttConfig.AppEUI, mqttConfig.Password, mqttConfig.Address)
	if err := m.client.Connect(); err != nil {
		return errors.New("Could not connect")
	}
	u := uplinkMessageHandler
	m.client.SubscribeUplink(u)
	return nil
}

// Disconnects the mqtt client, used for a graceful shutdown.
func (m *mqttClient) Disconnect() {
	m.client.Disconnect()
}

// Gets the internal mqtt client used.
func (m *mqttClient) GetClient() mqtt.Client {
	return m.client
}
