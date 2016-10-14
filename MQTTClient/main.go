package main

import (
	"fmt"
	"os"
	"time"

	"github.com/LoRaWanSoFa/MQTTClient/NewClient"
	//import the Paho Go MQTT library

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

//define a function for the default message handler
var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

var messageLoad = make(chan bool)

func messageHandler(client MQTT.Client, msg MQTT.Message) {
	messageLoad <- true
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func main() {
	NewClient.NewConnect()
}

func oldConnect() {
	//create a ClientOptions struct setting the broker address, clientid, turn
	//off trace output and set the default message handler
	opts := MQTT.NewClientOptions().AddBroker("")
	//opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("")
	opts.SetCleanSession(true)
	opts.SetUsername("")
	opts.SetPassword("")
	opts.SetDefaultPublishHandler(f)

	//opts.SetDefaultPublishHandler(f)

	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
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
		token.Wait()
	}

	time.Sleep(10 * time.Second)

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
