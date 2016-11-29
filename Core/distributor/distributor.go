package distributor

import (
	components "github.com/LoRaWanSoFa/LoRaWanSoFa/Components"
	"github.com/LoRaWanSoFa/LoRaWanSoFa/MQTTClient/NewClient"
	"github.com/LoRaWanSoFa/ttn/core"
	"github.com/LoRaWanSoFa/ttn/core/types"
)

type Distributor interface {
}

type distributor struct {
}

func uplinkMessageHandler(client NewClient.Client, appEUI types.AppEUI, devEUI types.DevEUI, req core.DataUpAppReq) {

}

func (d *distributor) deduplicate(message components.Message) {

}
