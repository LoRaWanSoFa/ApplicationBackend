package main

import (
	"log"
	"time"

	mdl "github.com/LoRaWanSoFa/LoRaWanSoFa/Components"
	dbc "github.com/LoRaWanSoFa/LoRaWanSoFa/DBC/DatabaseConnector"
	_ "github.com/lib/pq"
)

//mdl "github.com/LoRaWanSoFa/LoRaWanSoFa/Components"

func main() {
	errPing := dbc.Connect()
	checkErr(errPing)
	log.Println("connected")
	//close it when the program ends.
	defer dbc.Close()
	time.Sleep(1000 * time.Millisecond) // wait a bit for simulation purposes
	log.Println("Start sending data now:")

	sensor := mdl.NewHeaderSensor(1, 0, 0, 0, 0, "conversion_expression") //soft_deleted = false
	sensor.Soft_deleted = true
	dbc.ChangeSingleSensorActivationState(sensor)

	time.Sleep(8000 * time.Millisecond) // Let program run for ever till panic?
}

func checkErr(err error) {
	if err != nil {
		log.Print("Fatal error!\n")
		panic(err)
	}
}
