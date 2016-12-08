package distributor

import components "github.com/LoRaWanSoFa/LoRaWanSoFa/Components"
import webserver "github.com/LoRaWanSoFa/LoRaWanSoFa/webserver"

type Distributor interface {
	InputUplink()
	InputDownlink(webserver.Message)
}

type distributor struct {
}

func New() Distributor {
	dist := new(distributor)
	return dist
}

func (d *distributor) InputUplink() {

}
func (d *distributor) InputDownlink(message webserver.Message) {

}

func (d *distributor) deduplicate(message components.MessageUplinkI) {

}

func (d *distributor) convertMessage() components.MessageUplinkI {
	return components.NewMessageUplink(1, "1")
}
