package main

import (
	"fmt"
	"os"
	"time"

	//import the Paho Go MQTT library
	"github.com/LoRaWanSoFa/MQTTClient/NewClient"
	"github.com/TheThingsNetwork/ttn/core"
	"github.com/TheThingsNetwork/ttn/core/types"
	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

//define a function for the default message handler
var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

var messageLoad = make(chan bool)

func messageHandler(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func uplinkMessageHandler(client NewClient.Client, appEUI types.AppEUI, devEUI types.DevEUI, req core.DataUpAppReq) {
	fmt.Println(appEUI.GoString())
	fmt.Println("got a message")
	fmt.Println(req)
}

func main() {
	//oldConnect()
	log.SetHandler(text.New(os.Stderr))

	ctx := log.WithField("Test", "AnotherOne")
	client := NewClient.NewClient(ctx, "ttnctl", "70B3D57ED0001162", "C8nzg+pocQVnC6yjuqyi/yrCEGiV9/s8QSQdEuQVuSE=", "tcp://staging.thethingsnetwork.org:1883")
	if err := client.Connect(); err != nil {
		ctx.WithError(err).Fatal("Could not connect")
	}
	client.Connect()
	u := uplinkMessageHandler
	eui := types.AppEUI{0x70, 0xB3, 0xD5, 0x7E, 0xD0, 0x00, 0x11, 0x62}
	client.SubscribeAppUplink(eui, u)
	time.Sleep(60 * time.Second)
	time.Sleep(60 * time.Second)
	time.Sleep(60 * time.Second)
	client.Disconnect()
}

func oldConnect() {
	//create a ClientOptions struct setting the broker address, clientid, turn
	//off trace output and set the default message handler
	opts := MQTT.NewClientOptions().AddBroker("tcp://staging.thethingsnetwork.org:1883")
	//opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetCleanSession(true)
	opts.SetUsername("70B3D57ED0001162")
	opts.SetPassword("C8nzg+pocQVnC6yjuqyi/yrCEGiV9/s8QSQdEuQVuSE=")
	fmt.Printf("%s \n", "test")
	//opts.SetDefaultPublishHandler(f)

	//create and start a client using the above ClientOptions
	fmt.Printf("%s \n", "test")
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Printf("%s \n", "end test")
	//subscribe to the topic /go-mqtt/sample and request messages to be delivered
	//at a maximum qos of zero, wait for the receipt to confirm the subscription

	if token := c.Subscribe("+/devices/+/up", 0, messageHandler); token.Wait() && token.Error() != nil {
		print("test")
		fmt.Println(token.Error())
		os.Exit(1)

	} else {
		go print("Subscribed:")
	}

	//Publish 5 messages to /go-mqtt/sample at qos 1 and wait for the receipt
	//from the server after sending each message
	for i := 0; i < 5; i++ {
		text := fmt.Sprintf("this is msg #%d!", i)
		token := c.Publish("+/devices/+/up", 0, false, text)
		time.Sleep(5 * time.Second)
		token.Wait()
	}

	//unsubscribe from /go-mqtt/sample
	if token := c.Unsubscribe("+/devices/+/up"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	/*
		messageCount := 0
		for i := 0; i < 100; i++ {

			select {
			case <-messageLoad:
				println("test")
				messageCount++
			}
		}
		fmt.Printf("Received %3d Broker Load messages\n", messageCount)
	*/
	c.Disconnect(250)
}
