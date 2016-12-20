// Package restUplink is used to send uplink messages to the ApplicationFrontend.
package restUplink

import (
	"fmt"
	"net/http"

	components "github.com/LoRaWanSoFa/ApplicationBackend/Components"
)

// Connector is the client that connects to the Rest API
// of the application frontend
type Connector interface {
	NewNode(nodeID string, longitude float64, latitude float64) (*http.Response, error)
	UpdateNode(nodeID string, longitude float64, latitude float64) (*http.Response, error)
	DeleteNode(nodeID string) (*http.Response, error)
	NewSensor(nodeID string, sensorID string) (*http.Response, error)
	DeleteSensor(nodeID string, sensorID string) (*http.Response, error)
	NewData(nodeID string, message components.MessageUplinkI) (*http.Response, error)
}

type restUplinkConnector struct {
	basicURL string
}

// NewRestUplinkConnector is the constructor for a Connector
func NewRestUplinkConnector(ip string, apiKey string) Connector {
	connector := new(restUplinkConnector)
	connector.basicURL = createBasicURL(ip, apiKey)
	return connector
}

func createBasicURL(ip string, apiKey string) string {
	return fmt.Sprintf("http://%s/api/%s", ip, apiKey)
}

func (r *restUplinkConnector) NewNode(nodeID string, longitude float64, latitude float64) (*http.Response, error) {
	//Post Request:
	//http://ip/api/{apikey}/newnode/{id}/{longitude}/{latitude}
	//Needs to be called if a new node is added to the backend.
	url := fmt.Sprintf("%s/newnode/%s/%f/%f", r.basicURL, nodeID, longitude, latitude)
	resp, err := http.PostForm(url, nil)
	return resp, err
}

func (r *restUplinkConnector) UpdateNode(nodeID string, longitude float64, latitude float64) (*http.Response, error) {
	//Put Request:
	//http://ip/api/{apikey}/updatenode/{id}/{longitude}/{latitude}
	//Needs to be called if the position of a node has changed.
	url := fmt.Sprintf("%s/updatenode/%s/%f/%f", r.basicURL, nodeID, longitude, latitude)
	resp, err := http.Get(url)
	return resp, err
}

func (r *restUplinkConnector) DeleteNode(nodeID string) (*http.Response, error) {
	//Delete Request:
	//http://ip/api/{apikey}/deletenode/{id}
	//Needs to be called if a node gets removed from the backend.
	client := &http.Client{}
	url := fmt.Sprintf("%s/deletenode/%s", r.basicURL, nodeID)
	req, _ := http.NewRequest("DELETE", url, nil)
	resp, err := client.Do(req)
	return resp, err
}

func (r *restUplinkConnector) NewSensor(nodeID string, sensorID string) (*http.Response, error) {
	//Post Request:
	//http://ip/api/{apikey}/newsensor/{id}/{nodeid}
	//Needs to be called if a sensor is added to a node.
	//Does he need any other sensor information?
	// if yes -> add json again
	url := fmt.Sprintf("%s/newsensor/%s/%s", r.basicURL, nodeID, sensorID)
	resp, err := http.PostForm(url, nil)
	return resp, err
}

func (r *restUplinkConnector) DeleteSensor(nodeID string, sensorID string) (*http.Response, error) {
	//Delete Request:
	//http://ip/api/{apikey}/deletesensor/{id}
	//Needs to be called if a sensor is removed from a node.
	client := &http.Client{}
	url := fmt.Sprintf("%s/deletesensor/%s/%s", r.basicURL, nodeID, sensorID)
	req, _ := http.NewRequest("DELETE", url, nil)
	resp, err := client.Do(req)
	return resp, err
	// isn't nodeID needed as well as sensorID
}

func (r *restUplinkConnector) NewData(nodeID string, message components.MessageUplinkI) (*http.Response, error) {
	//Post Request:
	//http://ip/api/{apikey}/newdata
	//Needs to be called if data from a node comes in.
	// add json to it
	url := fmt.Sprintf("%s/newData/%s", r.basicURL, nodeID)
	resp, err := http.PostForm(url, message.ToJSON())
	// isn't nodeID needed here?
	return resp, err
}
