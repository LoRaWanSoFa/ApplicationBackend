package restUplink

import (
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"testing"

	components "github.com/LoRaWanSoFa/ApplicationBackend/Components"
)

var testMessage = components.NewMessageUplink(123, "1234")

var testData = struct {
	nodeID    string
	sensorID  string
	longitude float64
	latitude  float64
	payload   url.Values
}{"AB13A02BD", "103", 40.730610, -73.935242, url.Values{"key": {"Value"}}}

var testConnector Connector
var testServer *http.Server
var l net.Listener

func TestMain(m *testing.M) {
	setUp()
	result := m.Run()
	tearDown()
	os.Exit(result)
}

func setUp() {
	var err error
	testConnector = NewRestUplinkConnector(":8080", "test")
	testServer = &http.Server{
		Addr: ":8080",
	}
	l, err = net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		log.Fatal(testServer.Serve(l))
	}()
}

func tearDown() {
	l.Close()
}

func TestNewNode(t *testing.T) {
	testConnector.NewNode(testData.nodeID, testData.longitude, testData.latitude)
}

func TestUpdateNode(t *testing.T) {
	testConnector.UpdateNode(testData.nodeID, testData.longitude, testData.latitude)
}

func TestDeleteNode(t *testing.T) {
	testConnector.DeleteNode(testData.nodeID)
}

func TestNewSensor(t *testing.T) {
	testConnector.NewSensor(testData.nodeID, testData.sensorID)
}

func TestDeleteSensor(t *testing.T) {
	testConnector.DeleteSensor(testData.nodeID, testData.sensorID)
}

func TestNewData(t *testing.T) {
	testConnector.NewData(testData.nodeID, testMessage)
}
