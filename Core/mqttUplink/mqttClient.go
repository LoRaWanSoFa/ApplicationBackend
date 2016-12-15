package mqttUplink

import (
	"errors"
	"os"

	goLog "log"

	"github.com/LoRaWanSoFa/LoRaWanSoFa/Components"
	"github.com/LoRaWanSoFa/LoRaWanSoFa/Core/distributor"
	"github.com/LoRaWanSoFa/ttn/core"
	"github.com/LoRaWanSoFa/ttn/core/types"
	"github.com/LoRaWanSoFa/ttn/mqtt"
	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

type MqttClient interface {
	Connect() error
	Disconnect()
	GetClient() mqtt.Client
}

type mqttClient struct {
	client mqtt.Client
}

func New() MqttClient {
	client := new(mqttClient)
	return client
}

var hHandler = NewHeaderHandler()
var mCreator = NewMessageCreator()
var dist = distributor.New()

func uplinkMessageHandler(client mqtt.Client, appEUI types.AppEUI, devEUI types.DevEUI, req core.DataUpAppReq) {
	if len(req.Payload) > 0 {
		flag := req.Payload[0]
		if flag>>4 == 1 {
			header, err := hHandler.CreateNewHeader(req.Payload, devEUI.GoString())
			if err != nil {
				goLog.Fatal(err)
			} else {
				err = hHandler.StoreHeader(header, devEUI.GoString())
				if err != nil {
					goLog.Fatal(err)
				}
			}
		} else if flag>>5 == 1 {
			message, err := mCreator.CreateMessage(req.Payload, devEUI.GoString())
			if err != nil {
				goLog.Fatal(err)
			} else {
				dist.InputUplink(message)
			}
		} else {
			err := errors.New("No valid flag")
			goLog.Fatal(err)
		}
	}
}

func (m *mqttClient) Connect() error {
	log.SetHandler(text.New(os.Stderr))
	mqttConfig := components.GetConfiguration().Mqtt
	ctx := log.WithField("distributorClient", "mqtt-distributorClient")
	m.client = mqtt.NewClient(ctx, "ttnctl", mqttConfig.AppEUI, mqttConfig.Password, mqttConfig.Address)
	if err := m.client.Connect(); err != nil {
		return errors.New("Could not connect")
	}
	u := uplinkMessageHandler
	m.client.SubscribeUplink(u)
	return nil
}

func (m *mqttClient) Disconnect() {
	m.client.Disconnect()
}

func (m *mqttClient) GetClient() mqtt.Client {
	return m.client
}
