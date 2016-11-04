package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/LoRaWanSoFa/MQTTClient/NewClient"
	"github.com/TheThingsNetwork/ttn/core"
	"github.com/TheThingsNetwork/ttn/core/types"
	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	yaml "gopkg.in/yaml.v2"
)

func uplinkMessageHandler(client NewClient.Client, appEUI types.AppEUI, devEUI types.DevEUI, req core.DataUpAppReq) {
	fmt.Println(appEUI.GoString())
	fmt.Println("got a message")
	fmt.Println(fmt.Sprintf("%s", devEUI))
	fmt.Println(req)
}

func main() {
	//oldConnect()
	log.SetHandler(text.New(os.Stderr))
	// START: yaml config block
	goPath := os.Getenv("GOPATH")
	yamlFile, err := ioutil.ReadFile(filepath.Join(goPath, "/src/github.com/LoRaWanSoFa/config.yaml"))
	if err != nil {
		panic(err)
	}
	var applicationData NewClient.ApplicationData
	err = yaml.Unmarshal(yamlFile, &applicationData)
	if err != nil {
		panic(err)
	}
	// END: yaml config block

	ctx := log.WithField("Test", "AnotherOne")
	client := NewClient.NewClient(ctx, "ttnctl", applicationData.AppEUI, applicationData.Password, "tcp://staging.thethingsnetwork.org:1883")
	if err = client.Connect(); err != nil {
		ctx.WithError(err).Fatal("Could not connect")
	}
	client.Connect()
	u := uplinkMessageHandler
	//var eui []byte
	eui := make([]byte, 8)
	eui, err = hex.DecodeString(applicationData.AppEUI)

	//appEUI := types.AppEUI(eui)
	EUI := types.AppEUI{eui[0], eui[1], eui[2], eui[3], eui[4], eui[5], eui[6], eui[7]}
	if err != nil {
		panic(err)
	}
	client.SubscribeAppUplink(EUI, u)
	time.Sleep(18000 * time.Second)
	client.Disconnect()
}
