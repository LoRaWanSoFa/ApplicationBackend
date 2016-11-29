package main

import (
	"log"
	"time"

	mdl "github.com/LoRaWanSoFa/Components"
	dbc "github.com/LoRaWanSoFa/DBC/DatabaseConnector"
	_ "github.com/lib/pq"
)

//"github.com/LoRaWanSoFa/DBC/DatabaseConnector"

func main() {
	errPing := dbc.Connect()
	checkErr(errPing)
	log.Println("connected")
	//close it when the program ends.
	defer dbc.Close()
	time.Sleep(2000 * time.Millisecond) // wait a bit for simulation purposes
	//log.Println(dbc.CheckDevEUI("A4C12BF"))
	log.Println("Start sending data now:")
	//log.Println(dbc.GetNodeSensors("A4C12BF"))
	//log.Println(dbc.GetNodeSensors("A4C12B2F"))
	//log.Println("DONE")

	sensors := dbc.GetNodeSensors("A4C12BF")
	m := mdl.NewMessageUplink(4, "A4C12BF")
	m.AddPayload([]byte("payload"), sensors[0])
	m.AddPayload([]byte("hmmm"), sensors[1])
	log.Printf("message with payloads: %+v", m)

	//Adding a new message in the database
	message, err := dbc.AddMessage("A4C12BF")
	checkErr(err)
	log.Printf("Message: %+v", message)
	message.AddPayload([]byte("payload"), sensors[0])
	message.AddPayload([]byte("hmmm"), sensors[1])
	//when it goes wrong you will get an error
	message, err = dbc.AddMessage("DOES NOT EXIST")
	log.Println("check errors")
	if err != nil {
		log.Println(err)
		//handle error
	}
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
