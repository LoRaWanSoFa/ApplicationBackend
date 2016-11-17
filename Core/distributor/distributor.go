package distributor

import components "github.com/LoRaWanSoFa/Components"

type Distributor interface {
}

type distributor struct {
}

func (d *distributor) deduplicate(message components.Message) {

}
