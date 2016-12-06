package distributor

import components "github.com/LoRaWanSoFa/LoRaWanSoFa/Components"

type Distributor interface {
	InputUplink()
	InputDownlink()
}

type distributor struct {
}

func New() Distributor {
	dist := new(distributor)
	return dist
}

func (d *distributor) InputUplink() {

}
func (d *distributor) InputDownlink() {

}

func (d *distributor) deduplicate(message components.MessageUplinkI) {

}

func (d *distributor) convertMessage() components.MessageUplinkI {
	return components.NewMessageUplink(1, "1")
}
