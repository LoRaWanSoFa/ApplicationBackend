package mqttDownlink

import (
	"errors"
	"os"

	components "github.com/LoRaWanSoFa/LoRaWanSoFa/Components"
	"github.com/LoRaWanSoFa/ttn/core"
	"github.com/LoRaWanSoFa/ttn/core/types"
	"github.com/LoRaWanSoFa/ttn/mqtt"
	apexLog "github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

type MqttClient interface {
	Connect() error
	Disconnect()
	PublishDownlink(appEUI types.AppEUI, devEUI types.DevEUI, payload core.DataDownAppReq)
}

type mqttClient struct {
	client mqtt.Client
}

func New() MqttClient {
	downClient := new(mqttClient)
	return downClient
}

// Connects the downlink mqtt Client to the backend.
func (m *mqttClient) Connect() error {
	apexLog.SetHandler(text.New(os.Stderr))
	mqttConfig := components.GetConfiguration().Mqtt
	ctx := apexLog.WithField("downlinkClient", "mqtt-downlinkClient")
	m.client = mqtt.NewClient(ctx, "ttnctl", mqttConfig.AppEUI, mqttConfig.Password, mqttConfig.Address)
	if err := m.client.Connect(); err != nil {
		return errors.New("Could not connect")
	}
	return nil
}

// Disconnects the mqtt downlink client.
func (m *mqttClient) Disconnect() {
	m.client.Disconnect()
}

// Publishes a downlink message to the backend.
func (m *mqttClient) PublishDownlink(appEUI types.AppEUI, devEUI types.DevEUI, payload core.DataDownAppReq) {
	m.client.PublishDownlink(appEUI, devEUI, payload)
}
