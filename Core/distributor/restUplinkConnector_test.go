package distributor

import (
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"testing"
)

var testData = struct {
	nodeId    string
	sensorId  string
	longitude float64
	latitude  float64
	payload   url.Values
}{"AB13A02BD", "103", 40.730610, -73.935242, url.Values{"key": {"Value"}}}

var testConnector RestUplinkConnector
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
	testConnector.NewNode(testData.nodeId, testData.longitude, testData.latitude)
}

func TestUpdateNode(t *testing.T) {
	testConnector.UpdateNode(testData.nodeId, testData.longitude, testData.latitude)
}

func TestDeleteNode(t *testing.T) {
	testConnector.DeleteSensor(testData.nodeId, testData.sensorId)
}

func TestNewSensor(t *testing.T) {
	testConnector.NewSensor(testData.nodeId, testData.sensorId)
}

func TestDeleteSensor(t *testing.T) {
	testConnector.DeleteSensor(testData.nodeId, testData.sensorId)
}

func TestNewData(t *testing.T) {
	testConnector.NewData(testData.nodeId, testData.payload)
}
