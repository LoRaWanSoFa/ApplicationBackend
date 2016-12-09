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
	//dbc.GetNodeSensorsb("")
	sensors := dbc.GetNodeSensors("A4C12BF")
	log.Printf("sensors: %+v", sensors)

	//m := mdl.NewMessageUplink(76, "A4C12BF")
	// m, err := dbc.AddMessage("A4C12BF")
	// checkErr(err)
	// m.RemovePayloads()
	// m.AddPayloadString("Howdee1", sensors[0])
	// m.AddPayloadString("Howdee2", sensors[1])
	// m.AddPayloadString("Howdee3", sensors[0])
	//
	// //m.AddPayload([]byte("payload"), sensors[0])
	// //m.AddPayload([]byte("hmmm"), sensors[1])
	// log.Printf("message with payloads: %+v", m)
	// err = dbc.StoreMessagePayloads(m)
	// checkErr(err)

	//store downlink message
	dm := new(mdl.MessageDownLink)
	dm.Deveui = "A4C12BF"
	dm.Payload = "DOWNLINKMESSAGE"
	dm.Time = time.Now()

	log.Printf("Message: %+v", dm)
	dmerror := dbc.StoreDownlinkMessage(dm)
	checkErr(dmerror)
	log.Printf("Message: %+v", dm)

	//Adding a new message in the database
	// message, err := dbc.AddMessage("A4C12BF")
	// checkErr(err)
	// log.Printf("Message: %+v", message)
	// message.AddPayload([]byte("payload"), sensors[0])
	// message.AddPayload([]byte("hmmm"), sensors[1])
	//when it goes wrong you will get an error
	// message, err = dbc.AddMessage("DOES NOT EXIST")
	// log.Println("check errors")
	// if err != nil {
	// 	log.Println(err)
	// 	//handle error
	// }
	// err = dbc.StoreMessagePayloads(message, sensors)
	// checkErr(err)
	//log.Printf("Message: %+v", message)
	// for index := 0; index < 100000; index++ {
	// 	go dbc.CheckDevEUI("1")
	// 	go dbc.CheckDevEUI("2")
	// 	go dbc.CheckDevEUI("3")
	// 	go dbc.CheckDevEUI("A4C12BF")
	// }
	time.Sleep(8000 * time.Millisecond) // Let program run for ever till panic?
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
